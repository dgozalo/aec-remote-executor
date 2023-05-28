package server

import (
	"github.com/dgozalo/aec-remote-executor/pkg/compiler"
	"github.com/dgozalo/aec-remote-executor/pkg/database"
	"github.com/dgozalo/aec-remote-executor/pkg/service"
	"go.temporal.io/sdk/client"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dgozalo/aec-remote-executor/pkg/graph"
)

const defaultPort = "8080"

func RunServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	temporalURL := os.Getenv("TEMPORALITE_HOST_PORT")
	if temporalURL == "" {
		temporalURL = "localhost:7233"
	}
	// Create the client object just once per process
	c, err := client.Dial(client.Options{
		HostPort: temporalURL,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	pg, err := database.NewPostgresDBAccess()
	if err != nil {
		log.Fatalln("unable to obtain database client")
	}
	temporalCompiler := compiler.NewTemporalCompiler(c)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		ExecutionService:  service.NewExecutionService(pg),
		ManagementService: service.NewManagementService(pg),
		TemporalCompiler:  temporalCompiler,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

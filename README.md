# AEC Remote Executor TFG Project

## Description
This is a University final project for the Bachelor's Degree in Computer Engineering at  the Universidad a Distancia de Madrid (UDIMA). It consists of a web application that allows universities to define
coding problems that can be executed remotely by students to receive feedback on test cases and code quality.

It's inspired by platforms such as [HackerRank](https://www.hackerrank.com/) or [LeetCode](https://leetcode.com/) but with a focus on the academic environment.

This prototype is a work in progress and it's not ready for production.

## Technologies

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [PostgreSQL](https://www.postgresql.org/)
- [Temporal](https://temporal.io/)
- [Go](https://golang.org/)
- [GraphQL](https://graphql.org/)
- [Apollo Server](https://www.apollographql.com/docs/apollo-server/)
- [Apollo Client](https://www.apollographql.com/docs/react/)
- [React](https://reactjs.org/)
- [JavaScript](https://www.javascript.com/)

## How to run

### Prerequisites

- [Temporalite](https://github.com/temporalio/temporalite)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Node.js](https://nodejs.org/en/)
- [Yarn](https://yarnpkg.com/)
- [Go](https://golang.org/)

### Running without Docker Compose

Run the following commands in the root directory of the project:

```bash
# Start Temporalite
temporalite start --namespace default --ephemeral
# Start PostgreSQL and run the migrations
docker-compose up --force-recreate migrate
# Start the worker
go run ./cmd/worker/main.go
# Start the server
go run ./cmd/server/main.go
# Start the frontend
cd frontend && yarn start
```
### Running with Docker Compose

Run the following commands in the root directory of the project:

```bash
 docker-compose up --force-recreate --build
```

This will build all services necessary for the application and start them.

The application will be available at http://localhost:3000.

The GraphQL Playground will be available at http://localhost:8080.

Temporalite will be available at http://localhost:8233.

## License

This project is licensed under the terms of the MIT license.
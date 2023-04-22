package graph

import (
	"github.com/dgozalo/aec-remote-executor/pkg/compiler"
	"github.com/dgozalo/aec-remote-executor/pkg/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ExecutionStore   database.Database
	TemporalCompiler *compiler.TemporalCompiler
}

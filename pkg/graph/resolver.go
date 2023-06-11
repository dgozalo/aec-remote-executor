package graph

import (
	"github.com/dgozalo/aec-remote-executor/pkg/compiler"
	"github.com/dgozalo/aec-remote-executor/pkg/executions"
	"github.com/dgozalo/aec-remote-executor/pkg/management"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ExecutionService  *executions.ExecutionService  // This is the service that handles the executions
	ManagementService *management.ManagementService // This is the service that handles the management
	TemporalCompiler  *compiler.TemporalCompiler    // This is the compiler that handles the temporal compiler
}

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
	ExecutionService  *executions.ExecutionService
	ManagementService *management.ManagementService
	TemporalCompiler  *compiler.TemporalCompiler
}

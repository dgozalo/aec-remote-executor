package graph

import (
	"github.com/dgozalo/aec-remote-executor/pkg/compiler"
	"github.com/dgozalo/aec-remote-executor/pkg/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ExecutionService  *service.ExecutionService
	ManagementService *service.ManagementService
	TemporalCompiler  *compiler.TemporalCompiler
}

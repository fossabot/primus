// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package command

import (
	"github.com/raba-jp/primus/pkg/backend"
	"github.com/raba-jp/primus/pkg/operations/command/handlers"
	"github.com/raba-jp/primus/pkg/operations/command/starlarkfn"
	"go.starlark.net/starlark"
)

// Injectors from wire.go:

func Command() func(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kargs []starlark.Tuple) (starlark.Value, error) {
	execInterface := backend.NewExecInterface()
	commandHandler := handlers.New(execInterface)
	v := starlarkfn.Command(commandHandler)
	return v
}

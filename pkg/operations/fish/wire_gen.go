// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package fish

import (
	"github.com/raba-jp/primus/pkg/backend"
	"github.com/raba-jp/primus/pkg/operations/fish/handlers"
	"github.com/raba-jp/primus/pkg/operations/fish/starlarkfn"
	"go.starlark.net/starlark"
)

// Injectors from wire.go:

func SetPath() func(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kargs []starlark.Tuple) (starlark.Value, error) {
	execInterface := backend.NewExecInterface()
	setPathHandler := handlers.NewSetPath(execInterface)
	v := starlarkfn.SetPath(setPathHandler)
	return v
}

func SetVariable() func(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kargs []starlark.Tuple) (starlark.Value, error) {
	execInterface := backend.NewExecInterface()
	setVariableHandler := handlers.NewSetVariable(execInterface)
	v := starlarkfn.SetVariable(setVariableHandler)
	return v
}

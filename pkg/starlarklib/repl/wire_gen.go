// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package repl

import (
	"github.com/raba-jp/primus/pkg/backend"
	"github.com/raba-jp/primus/pkg/starlarklib/functions"
)

// Injectors from wire.go:

func Initialize() PromptFunc {
	state := NewState()
	thread := newThread()
	execInterface := backend.NewExecInterface()
	fs := backend.NewFs()
	backendBackend := backend.New(execInterface, fs)
	stringDict := functions.NewPredeclaredFunction(backendBackend, execInterface, fs)
	replREPL := NewREPL(state, thread, stringDict)
	executor := NewExecutor(replREPL)
	completer := NewCompleter()
	promptFunc := NewPrompt(replREPL, executor, completer)
	return promptFunc
}

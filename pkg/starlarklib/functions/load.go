package functions

import (
	"path/filepath"

	"github.com/raba-jp/primus/pkg/starlarklib"
	"github.com/spf13/afero"
	"go.starlark.net/starlark"
	"golang.org/x/xerrors"
)

type StarlarkLoadFn = func(thread *starlark.Thread, module string) (starlark.StringDict, error)

func Load(dryrun bool, fs afero.Fs, predeclared starlark.StringDict) StarlarkLoadFn {
	return func(thread *starlark.Thread, module string) (starlark.StringDict, error) {
		var modulePath string
		if filepath.IsAbs(module) {
			modulePath = module
		} else {
			path := starlarklib.GetCurrentFilePath(thread)
			modulePath = filepath.Join(filepath.Dir(path), module)
		}

		data, err := afero.ReadFile(fs, modulePath)
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}

		ctx := starlarklib.GetCtx(thread)
		childThread := &starlark.Thread{
			Name: module,
			Load: Load(dryrun, fs, predeclared),
		}
		starlarklib.SetCtx(ctx, childThread)
		starlarklib.SetDryRun(childThread, dryrun)

		return starlark.ExecFile(childThread, modulePath, data, predeclared)
	}
}

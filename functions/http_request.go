package functions

import (
	"context"

	"github.com/raba-jp/primus/executor"
	"go.starlark.net/starlark"
	"golang.org/x/xerrors"
)

func HttpRequest(ctx context.Context, exc executor.Executor) StarlarkFn {
	return func(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kargs []starlark.Tuple) (starlark.Value, error) {
		url, path, err := parseHttpRequestFnArgs(b, args, kargs)
		if err != nil {
			return starlark.False, xerrors.Errorf(": %w", err)
		}

		ret, err := exc.HttpRequest(ctx, &executor.HttpRequestParams{URL: url, Path: path})
		if err != nil {
			return toStarlarkBool(ret), xerrors.Errorf(": %w", err)
		}
		return toStarlarkBool(ret), nil
	}
}

func parseHttpRequestFnArgs(
	b *starlark.Builtin,
	args starlark.Tuple,
	kargs []starlark.Tuple,
) (url string, path string, err error) {
	err = starlark.UnpackArgs(b.Name(), args, kargs, "url", &url, "path", &path)
	if err != nil {
		return "", "", xerrors.Errorf("Failed to parse http_request function args: %w", err)
	}
	return
}

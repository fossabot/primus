package functions_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_executor "github.com/raba-jp/primus/executor/mock"
	"github.com/raba-jp/primus/functions"
	"go.starlark.net/starlark"
)

func TestFileMove(t *testing.T) {
	tests := []string{
		`file_move(src="/sym/src.txt", dest="/sym/dest.txt")`,
		`file_move("/sym/src.txt", "/sym/dest.txt")`,
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_executor.NewMockExecutor(ctrl)
			m.EXPECT().FileMove(gomock.Any(), gomock.Any()).Return("")

			predeclared := starlark.StringDict{
				"file_move": starlark.NewBuiltin("file_move", functions.FileMove(context.Background(), m)),
			}
			thread := &starlark.Thread{
				Name: "testing",
			}
			_, err := starlark.ExecFile(thread, "test.star", tt, predeclared)
			if err != nil {
				t.Fatalf("%v", err)
			}
		})
	}
}

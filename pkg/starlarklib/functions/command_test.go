package functions_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_executor "github.com/raba-jp/primus/pkg/executor/mock"
	"github.com/raba-jp/primus/pkg/starlarklib/functions"
	"go.starlark.net/starlark"
	"golang.org/x/xerrors"
)

func TestCommand(t *testing.T) {
	tests := []struct {
		name   string
		data   string
		mock   func(*mock_executor.MockExecutor)
		hasErr bool
	}{
		{
			name: "success",
			data: `command(name="echo", args=["hello", "world"])`,
			mock: func(m *mock_executor.MockExecutor) {
				m.EXPECT().Command(gomock.Any(), gomock.Any()).Return(true, nil)
			},
			hasErr: false,
		},
		{
			name: "success: no args",
			data: `command("echo")`,
			mock: func(m *mock_executor.MockExecutor) {
				m.EXPECT().Command(gomock.Any(), gomock.Any()).Return(true, nil)
			},
			hasErr: false,
		},
		{
			name: "success: with user and cwd",
			data: `command("echo", [], user="testuser", cwd="/home/testuser")`,
			mock: func(m *mock_executor.MockExecutor) {
				m.EXPECT().Command(gomock.Any(), gomock.Any()).Return(true, nil)
			},
			hasErr: false,
		},
		{
			name:   "error: too many arguments",
			data:   `command("echo", [], "testuser", "/home/testuser", "too many")`,
			mock:   func(m *mock_executor.MockExecutor) {},
			hasErr: true,
		},
		{
			name: "error: execute command failed",
			data: `command("echo")`,
			mock: func(m *mock_executor.MockExecutor) {
				m.EXPECT().Command(gomock.Any(), gomock.Any()).Return(false, xerrors.New("dummy"))
			},
			hasErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_executor.NewMockExecutor(ctrl)
			tt.mock(m)

			predeclared := starlark.StringDict{
				"command": starlark.NewBuiltin("command", functions.Command(m)),
			}

			thread := &starlark.Thread{
				Name: "testing",
			}
			_, err := starlark.ExecFile(thread, "test.star", tt.data, predeclared)
			if !tt.hasErr && err != nil {
				t.Error(err)
			}
		})
	}
}
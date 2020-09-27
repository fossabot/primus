package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/raba-jp/primus/pkg/cli/ui"
	"github.com/raba-jp/primus/pkg/exec"
	"github.com/spf13/afero"
	"golang.org/x/xerrors"
)

type DarwinPkgCheckInstallHandler interface {
	CheckInstall(ctx context.Context, name string) bool
}

type DarwinPkgCheckInstallHandlerFunc func(ctx context.Context, name string) bool

func (f DarwinPkgCheckInstallHandlerFunc) CheckInstall(ctx context.Context, name string) bool {
	return f(ctx, dryrun, p)
}

type DarwinPkgInstallParams struct {
	Name   string
	Option string
	Cask   bool
	Cmd    string
}

type DarwinPkgInstallHandler interface {
	Install(ctx context.Context, dryrun bool, p *DarwinPkgInstallParams) error
}

type DarwinPkgInstallHandlerFunc func(ctx context.Context, dryrun bool, p *DarwinPkgInstallParams) error

func (f DarwinPkgInstallHandlerFunc) Install(ctx context.Context, dryrun bool, p *DarwinPkgInstallParams) error {
	return f(ctx, dryrun, p)
}

type DarwinPkgUninstallParams struct {
	Name string
	Cask bool
	Cmd  string
}

type DarwinPkgUninstallHandler interface {
	Uninstall(ctx context.Context, dryrun bool, p *DarwinPkgUninstallParams) error
}

type DarwinPkgUninstallHandlerFunc func(ctx context.Context, dryrun bool, p *DarwinPkgUninstallParams) error

func (f DarwinPkgUninstallHandlerFunc) Uninstall(ctx context.Context, dryrun bool, p *DarwinPkgUninstallParams) error {
	return f(ctx, dryrun, p)
}

func NewDarwinPkgCheckInstall(execIF exec.Interface, fs afero.Fs) DarwinPkgCheckInstallHandler {
	return DarwinPkgCheckInstallHandlerFunc(func(ctx context.Context, name string) bool {
		return checkDarwinPkgInstall(ctx, execIF, fs, name)
	})
}

func NewDarwinPkgInstall(execIF exec.Interface, fs afero.Fs) DarwinPkgInstallHandler {
	return DarwinPkgInstallHandlerFunc(func(ctx context.Context, dryrun bool, p *DarwinPkgInstallParams) error {
		if dryrun {
			ui.Printf("brew install %s %s\n", p.Option, p.Name)
			return nil
		}

		if installed := checkDarwinPkgInstall(ctx, execIF, fs, p.Name); installed {
			return nil
		}

		if err := b.Exec.CommandContext(ctx, "brew", "install", p.Option, p.Name).Run(); err != nil {
			return xerrors.Errorf("Install package failed: %s: %w", p.Name, err)
		}
		return nil
	})
}

func NewDarwinPkgUninstall(execIF exec.Interface, fs afero.Fs) DarwinPkgUninstallHandler {
	return DarwinPkgUninstallHandlerFunc(func(ctx context.Context, dryrun bool, p *DarwinPkgUninstallParams) error {
		if dryrun {
			ui.Printf("brew uninstall %s\n", p.Name)
			return nil
		}

		if installed := checkDarwinPkgInstall(ctx, execIF, fs, p.Name); !installed {
			return nil
		}

		if err := b.Exec.CommandContext(ctx, "brew", "uninstall", p.Name).Run(); err != nil {
			return xerrors.Errorf("Remove package failed: %w", err)
		}
		return nil
	})
}

func checkDarwinPkgInstall(ctx context.Context, execIF exec.Interface, fs afero.Fs, name string) bool {
	installed := false
	walkFn := func(path string, info os.FileInfo, err error) error {
		installed = installed || strings.Contains(path, name)
		return nil
	}

	// brew list
	res, _ := execIF.CommandContext(ctx, "brew", "--prefix").Output()
	prefix := strings.ReplaceAll(string(res), "\n", "")
	_ = afero.Walk(fs, fmt.Sprintf("%s/Celler", prefix), walkFn)

	// brew cask list
	_ = afero.Walk(fs, "/opt/homebrew-cask/Caskroom", walkFn)
	_ = afero.Walk(fs, "/usr/local/Caskroom", walkFn)

	return installed
}

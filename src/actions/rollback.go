package actions

import (
	"errors"

	"github.com/eagerio/Stout/src/providers"
	"github.com/eagerio/Stout/src/types"
)

func Rollback(g types.GlobalFlags, r types.RollbackFlags) error {
	if g.FS == "" {
		return errors.New("The --fs flag and value are required for the `rollback` command")
	}

	if r.Version == "" {
		return errors.New("The --version flag and value are required after the `rollback` command")
	}

	err, fsProvider := providers.ValidateProviderType(g.FS, providers.FS_PROVIDER_TYPE)
	if err != nil {
		return err
	}

	fsProviderTyped, _ := fsProvider.(providers.FSProvider)
	if err := fsProviderTyped.ValidateSettings(); err != nil {
		return err
	}

	if err := fsProviderTyped.RollbackFS(g, r); err != nil {
		return err
	}

	return nil
}

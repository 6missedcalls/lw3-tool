package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/6missedcalls/lw3-tool/internal/errors"
)

var (
	Path string
)

// This method will create the path for the built contracts
// This path will be hardcoded for now but will change later in favor of a viper config
func CheckPath() (string, error) {
	// get homedir of the machine
	home, err := os.UserHomeDir()
	if err != nil {
		return errors.FailedHomedirRetrieval, err
	}
	// define the path
	Path = home + "/.scaffold/contracts/"

	// check if the path exists
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		// if it does not exist, create the path
		err = os.MkdirAll(Path, 0755)
		if err != nil {
			return errors.FailedPathCreation, err
		}
		fmt.Println("Path created at: " + Path)
	}
	return Path, nil
}

func GetVersionInfo(ctx context.Context) (string, string, string) {
	version := ctx.Value("version").(string)
	commit := ctx.Value("commit").(string)
	date := ctx.Value("date").(string)
	return version, commit, date
}

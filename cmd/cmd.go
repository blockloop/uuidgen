package cmd

import (
	"fmt"

	"github.com/satori/go.uuid"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app         = kingpin.New("uuidgen", "Create a new UUID value.")
	randomBased bool
	timeBased   bool
)

func init() {
	app.HelpFlag.Short('h')
	app.Version("2.28.2")
	app.VersionFlag.Short('V')

	app.Flag("random", "generate a random-based uuid").
		Short('r').
		BoolVar(&randomBased)

	app.Flag("time", "generate a time-based uuid").
		Short('t').
		BoolVar(&timeBased)
}

// Execute executes the root command
func Execute(cliArgs []string) error {
	_, err := app.Parse(cliArgs)
	if err != nil {
		return err
	}

	if timeBased {
		fmt.Println(uuid.NewV1())
		return nil
	}

	fmt.Println(uuid.NewV4())
	return nil
}

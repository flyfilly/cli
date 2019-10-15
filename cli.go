package cli

import (
	"fmt"
	"os"
)

// GetCommandLineArgs returns a list of a map of os.Args with keys starting with --
func GetCommandLineArgs(help []string) map[string]string {
	argsMap := map[string]string{}
	clArgs := os.Args

	for i, arg := range clArgs {
		if arg == "--help" {
			for _, item := range help {
				fmt.Println(item)
			}
			os.Exit(1)
		}

		if i%2 != 0 {
			argsMap[arg] = clArgs[i+1]
		}
	}

	return argsMap
}

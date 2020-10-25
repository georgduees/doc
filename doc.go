package doc

import (
	"fmt"

	"github.com/thatisuday/commando"
)

func main() {
	//initialize command
	commando.
		SetExecutableName("doc").
		SetVersion("0.0.1").
		SetDescription("This Tool creates Logs and Helps with Docs\n It can also create shareable links for Twitter.")
	//configure the root command
	commando.
		Register(nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `root` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})
	commando.
		Register("init").
		SetShortDescription("initialize in current repository").
		SetDescription("This command initiates the DaysOfCode Helper structure in the current directory.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `info` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// parse commandline
	commando.Parse(nil)
}

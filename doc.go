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
		AddArgument
	

		
	// parse commandline
	commando.Parse(nil)
}

package main

import (
	"os"

	"github.com/carlmjohnson/exitcode"
	"github.com/spotlightpa/at-5000/autotweeter"
)

func main() {
	exitcode.Exit(autotweeter.CLI(os.Args[1:]))
}

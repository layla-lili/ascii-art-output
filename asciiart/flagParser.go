package asciiart

import (
	"flag"
	"fmt"
	"os"
)

var Justify string
var Color string
var Output string

func Init() {
	const (
		defaultJustify = "left"
		usageJustify   = "Justify text"

		defaultColor   = "white"
		usageColor     = "Color text"

		defaultOutput  = "false"
		// usageOutput    = "Output text"
		usageOutput   = "Output text (use equal sign to specify output file):"
		// usageOutput    = "Usage: go run . [OPTION] [STRING] [BANNER]\n EX: go run . --output=<fileName.txt> something standard"
	)
	 // Define a custom usage message
	 flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [flags] input banner\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "Flags:\n")
        flag.PrintDefaults()
    }
	flag.StringVar(&Justify, "justify", defaultJustify, usageJustify)
	flag.StringVar(&Justify, "j", defaultJustify, usageJustify+" shorthand")

	flag.StringVar(&Color, "color", defaultColor, usageColor)
	flag.StringVar(&Color, "c", defaultColor, usageColor+" shorthand")

	flag.StringVar(&Output, "output", defaultOutput, usageOutput)
	flag.StringVar(&Output, "o", defaultOutput, usageOutput+" shorthand")
	
	
}

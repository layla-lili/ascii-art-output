package main

import (
	asciiart "asciiart/asciiart"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	//non english case

	for _, cha := range os.Args[1:] {
		for _, ch := range cha {
			if ch > 127 && ch < 32 {
				fmt.Println("Only Standard Ascii character PLease!")
				return
			}
		}
	}
	//Non-Flag Case
	if len(os.Args[1:]) < 3 && !strings.Contains(strings.Join(os.Args[1:], " "), "-") {
		if len(os.Args) == 0 {
			fmt.Println("At leat one Argument Needed")
			return
		}
		if len(os.Args[1:]) == 1 {
			//fmt.Println("argument is ", os.Args[1:])
			newstr := strings.Join(os.Args[1:], " ")
			// Testing banner
			readerx, err := fs.ReadFile(os.DirFS("./banners"), "standard.txt")
			if err != nil {
				fmt.Println("Error Opening File")
			}
			// Formate banner file
			asciiart.BannerFmt(readerx)
			// Handle "" && "\n"
			if newstr == "" {
				return
			} else if newstr == "\\n" {
				fmt.Println()
				return
			}

			asciiart.AsciiPrepx(newstr)

			return
		}
		if len(os.Args[1:]) == 2 {
			//newargs[1] banner name
			newargs := os.Args[1:]
			
			newstr := newargs[0]
			
			// Testing banner
			readerx, err := fs.ReadFile(os.DirFS("./banners"), newargs[1]+".txt")
			if err != nil {
				fmt.Println("Error Opening File")
			}
			// Formate banner file
			asciiart.BannerFmt(readerx)
			// Handle "" && "\n"
			if newstr == "" {
				return
			} else if newstr == "\\n" {
				fmt.Println()
				return
			}

			asciiart.AsciiPrepx(newstr)

			return
		}
	}

	asciiart.Init()

	//check Equal Sign
	equalSign := false
	cmdArgs := os.Args[1:]

	re := regexp.MustCompile(`(?:-output|-o|--output|--o)\s*([^\s]*)`)

	match := re.FindStringSubmatch(strings.Join(cmdArgs, " "))

	if match != nil {

		if strings.Contains(match[1], "=") {

			equalSign = true
		}
	}

	if !equalSign && len(os.Args[1:]) > 0 {
		fmt.Println("Error: Must have equal '=' after the flag")
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}

	//check for non-flag missing argument
	if err := run(); err != nil {
		//fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}

	// Access the flag values
	// fmt.Printf("Justify: %s\n", asciiart.Justify)
	// fmt.Printf("Color: %s\n", asciiart.Color)
	//fmt.Printf("Output: %s\n", asciiart.Output)

	// Access the non-flag arguments
	var bannerFont string
	var TheString string
	args := flag.Args()

	bannerFont = args[len(args)-1]
	TheString = strings.Join(args[:len(args)-1], " ")

	/////////////////
	// Testing banner
	reader, err := fs.ReadFile(os.DirFS("./banners"), bannerFont+".txt")
	if err != nil {
		fmt.Println("Error Opening File")
		return
	}

	// Formate banner file
	asciiart.BannerFmt(reader)

	// Handle "" && "\n"
	if TheString == "" {
		return
	} else if TheString == "\\n" {
		//fmt.Println()
		return
	}

	outputString := asciiart.AsciiPrep(TheString)
	var str string
	for _, b := range outputString {
		str += string(b)
	}

	// fmt.Printmln(str)
	/////////////////
	// If output filename is specified, write ASCII art to file
	if strings.TrimPrefix(filepath.Ext(asciiart.Output), ".") != "txt" {
		fmt.Println("File name should Ends with .txt!")
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	} else {
		file, err := os.Create(asciiart.Output)
		if err != nil {
			// Handle error
			fmt.Println(err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(str)
		if err != nil {
			//Handle error
			fmt.Println(err)
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("Example: go run . --output=<fileName.txt> something standard")
			return

		}
	}
}

// ////////////////////////////////
func run() error {
	// Parse command-line arguments
	flag.Parse()

	//check if file name provided
	if len(asciiart.Output) == 0 || asciiart.Output == "false" {
		return errors.New("please provide file name")
	}

	// Check if there are any non-flag arguments

	if flag.NArg() == 0 {
		return errors.New("no input string specified")
	}

	if flag.NArg() == 1 {
		return errors.New("no banner specified")
	}

	if flag.NArg() > 2 {
		return errors.New("too many arguments specified")
	}

	return nil
}

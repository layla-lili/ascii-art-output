package main

import (
	asciiart "asciiart/asciiart"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
)


func main() {
	asciiart.Init()
    // Parse the command line arguments
	flag.Parse()

	// Access the flag values
	fmt.Printf("Justify: %s\n", asciiart.Justify)
	fmt.Printf("Color: %s\n", asciiart.Color)
	fmt.Printf("Output: %s\n", asciiart.Output)

	// Access the non-flag arguments
	args := flag.Args()
	bannerFont := args[len(args)-1]
	TheString := strings.Join(args[:len(args)-1]," ")
	fmt.Printf("Non-flag arguments: %v\n", bannerFont)
	fmt.Printf("Non-flag arguments: %v\n", TheString)


	
    /////////////////
	// Testing banner
	reader, err := fs.ReadFile(os.DirFS("./banners"), bannerFont+ ".txt")
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

    fmt.Println(str)

	/////////////////
	file, err := os.Create(asciiart.Output)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	defer file.Close()


	_, err = file.WriteString(str)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}
	fmt.Println("String written to file successfully!")
}
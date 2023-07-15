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

	asciiart.Init()

	//check number of flag variable
	// numberOfFlags := flag.NFlag()
	// if numberOfFlags <=0 {
	// 	//fmt.Println("At Least one flag should be used \n Usage: go run . [OPTION] [STRING] [BANNER]")
	// 	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	// 	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
	// 	//os.Exit(1)
	// }

	//check Equal Sign
	equalSign := false
	cmdArgs := os.Args[1:]
	re := regexp.MustCompile(`(?:-output|-o|--output|--o)\s*([^\s]*)\b`)
   
    match := re.FindStringSubmatch(strings.Join(cmdArgs, " "))
    if match != nil {
      if strings.Contains(match[1],"="){
		equalSign = true
	  }
    }

	
	if !equalSign{
		fmt.Println("Error: Must have equal '=' after the flag")
		os.Exit(1)
	}


	if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
	  

    // Check for required input arguments
    // if flag.NArg() < 2 {
    //     fmt.Fprintf(os.Stderr, "Error: not enough input arguments\n")
    //     flag.Usage()
    //     os.Exit(1)
    // }
   
	//check for non-flag missing argument 
	if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
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
	TheString = strings.Join(args[:len(args)-1]," ")
	
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

   // fmt.Printmln(str)
	/////////////////
 // If output filename is specified, write ASCII art to file
 
//   if asciiart.Output != "" {
	if strings.TrimPrefix(filepath.Ext(asciiart.Output), ".") != "txt"{
			fmt.Println("File name should Ends with .txt!")
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
			os.Exit(1)
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
//////////////////////////////////
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
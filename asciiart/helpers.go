package asciiart

import (
	"bytes"
	"fmt"
)

var Chars [][]byte
var Text [][]byte
var Indx []int

// Create a new [][]byte from banner into Chars
// // case standard & shadow banners
func BannerReFmt(r []byte) {
	split := bytes.Split(r, []byte("\n")) // create banner split on new lines

	for _, line := range split {
		if len(line) == 0 {
			continue
		}
		Chars = append(Chars, line)
	}
}

// special BannerFormat for thinkertoy banner
func BannerFmt(r []byte) {
	split := bytes.Split(r, []byte("\r")) // create banner split on carriage returns

	for _, line := range split {
		if len(line) == 0 {
			continue
		}
		Chars = append(Chars, line)
	}
	newBannerFMT := ReformatTT(Chars) // copy banner split on carriage return from Chars (type [][]byte) into []byte
	Chars = nil                       // empty Chars (prepare to repopulate)
	BannerReFmt(newBannerFMT)           // repopulate Chars with banner split on new lines
}

// Prepare thinkertoy banner
func ReformatTT(chars [][]byte) []byte {
	var ssB []byte
	for _, line := range chars {
		ssB = append(ssB, line...)
	}
	return ssB
}

// Main text preparation 
func AsciiPrep(s string) [][]byte {
	words := Stob(s)
	for i, slice := range words {
		if slice[0] == 10 && i == 0 {
			Text = append(Text, []byte("\n"))
			Indx = nil
			continue
		}
		if slice[0] == 10 {
			if words[i-1][0] == 10 {
				Text = append(Text, []byte("\n"))
				Indx = nil
				continue
			} else {
				Indx = nil
				continue
			}

		}
		for _, char := range slice {
			Indexer(char)
		}
		LineFmt(Indx)

	}
	//Printer(Text)
	return Text
}
//Printer ascii printer ascii preperation 
// Main text preparation and printing function
func AsciiPrepx(s string) {
	words := Stob(s)
	for i, slice := range words {
		if slice[0] == 10 && i == 0 {
			Text = append(Text, []byte("\n"))
			Indx = nil
			continue
		}
		if slice[0] == 10 {
			if words[i-1][0] == 10 {
				Text = append(Text, []byte("\n"))
				Indx = nil
				continue
			} else {
				Indx = nil
				continue
			}
		}
		for _, char := range slice {
			Indexer(char)
		}
		LineFmt(Indx)
	}
	Printer(Text)
}

// locate first line index for each char in string (word)
func Indexer(b byte) {
	s, _ := Location(b)
	Indx = append(Indx, s)
}

// Locate StartingIndex:EndingIndex for each char
func Location(b byte) (s int, e int) {
	locS := (int(b) - 32) * 8
	locE := locS + 8

	return locS, locE
}

// create a new [][]byte of string (word) for printing
// // by appending each line of char byte slice into one []byte
func LineFmt(indx []int) {
	for i := 0; i < 8; i++ {
		var line []byte
		for _, item := range indx {
			
			slice := Chars[item+i]
			line = append(line, slice...)
		}
		line = append(line, byte('\n'))
		Text = append(Text, line)
	}
}

func Printer(Text [][]byte) {
	for _, line := range Text {
		for _, char := range line {
			fmt.Printf("%s", string(char))
		}
	}
}

// Convert string input to [][]byte
func Stob(s string) [][]byte {
	var bslice [][]byte
	var bline []byte
	for i, item := range s {
		if item == '\\' {
			if s[i+1] == 'n' {
				if bline != nil {

					bslice = append(bslice, bline)
				}
				bline = nil
				bline = append(bline, byte(10))
				bslice = append(bslice, bline)
				bline = nil
				continue
			}
		}
		if item == 'n' {
			if s[i-1] == '\\' {
				continue
			}
		}
		bline = append(bline, byte(item))
	}
	if bline != nil {

		bslice = append(bslice, bline)
	}
	return bslice
}

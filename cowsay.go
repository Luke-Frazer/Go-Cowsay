package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func print_animal(name string) {

	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	var stegosaurus = `         \                      .       .
          \                    / ` + "`" + `.   .' "
           \           .---.  <    > <    >  .---.
            \          |    \  \ - ~ ~ - /  /    |
          _____           ..-~             ~-..-~
         |     |   \~~~\\.'                    ` + "`" + `./~~~/
        ---------   \__/                         \__/
       .'  O    \     /               /       \  "
      (_____,    ` + "`" + `._.'               |         }  \/~~~/
       ` + "`" + `----.          /       }     |        /    \__/
             ` + "`" + `-.      |       /      |       /      ` + "`" + `. ,~~|
                 ~-.__|      /_ - ~ ^|      /- _      ` + "`" + `..-'
                      |     /        |     /     ~-.     ` + "`" + `-. _  _  _
                      |_____|        |_____|         ~ - . _ _ _ _ _>

	`

	var tRex = `       \     
	\
	 \
	  \
             .-=-==--==--.
       ..-=="  ,'o` + "`" + `)      ` + "`" + `.
     ,'         ` + "`" + `"'         \
    :  (                     ` + "`" + `.__...._
    |                  )    /         ` + "`" + `-=-.
    :       ,vv.-._   /    /               ` + "`" + `---==-._
     \/\/\/VV ^ d88` + "`" + `;'    /                         ` + "`" + `.
         ` + "`" + `` + "`" + `  ^/d88P!'    /             ,              ` + "`" + `._
            ^/    !'   ,.      ,      /                  "-,,__,,--'""""-.
           ^/    !'  ,'  \ . .(      (         _           )  ) ) ) ))_,-.\
          ^(__ ,!',"'   ;:+.:%:a.     \:.. . ,'          )  )  ) ) ,"'    '
          ',,,'','     /o:::":%:%a.    \:.:.:         .    )  ) _,'
           """'       ;':::'' ` + "`" + `+%%%a._  \%:%|         ;.). _,-""
                  ,-='_.-'      ` + "`" + `` + "`" + `:%::)  )%:|        /:._,"
                 (/(/"           ," ,'_,'%%%:       (_,'
                                (  (//(` + "`" + `.___;        \
                                 \     \    ` + "`" + `         ` + "`" + `
                                  ` + "`" + `.    ` + "`" + `.   ` + "`" + `.        :
                                    \. . .\    : . . . :
                                     \. . .:    ` + "`" + `.. . .:
                                      ` + "`" + `..:.:\     \:...\
                                       ;:.:.;      ::...:
                                       ):%::       :::::;
                                   __,::%:(        :::::
                                ,;:%%%%%%%:        ;:%::
                                  ;,--""-.` + "`" + `\  ,=--':%:%:\
                                 /"       "| /-".:%%%%%%%\
                                                 ;,-"'` + "`" + `)%%)   
                                                /"      "|`

	switch name {
	case "cow":
		fmt.Println(cow)
	case "stegosaurus":
		fmt.Println(stegosaurus)
	case "trex":
		fmt.Println(tRex)
	default:
		fmt.Println("Unknown Animal")
	}
}

// converts all tabs in the list of strings to 4 spaces
func tabs_to_spaces(lines []string) []string {
	var returnString []string
	for _, line := range lines {
		line = strings.Replace(line, "\t", "    ", -1)
		returnString = append(returnString, line)
	}
	return returnString
}

// gets the max width line in the passed in lines
func calculate_max_width(lines []string) int {
	widthWinner := 0
	for _, line := range lines {
		if utf8.RuneCountInString(line) > widthWinner {
			widthWinner = len(line)
		}
	}
	return widthWinner
}

// subtracts from maxWidth the number of characters in each line and fills in the rest with spaces
func normalize_string_length(lines []string, maxWidth int) []string {
	var returnString []string
	var currentString string
	for _, line := range lines {
		currentString = line + strings.Repeat(" ", maxWidth-utf8.RuneCountInString(line))
		returnString = append(returnString, currentString)
	}
	return returnString
}

// creates the bubble around the text and returns it as a string
func build_bubble(lines []string, maxWidth int) string {
	var numLines int = len(lines)
	var returnString []string
	var currentString string

	topBorder := " " + strings.Repeat("_", maxWidth+2)
	bottomBorder := " " + strings.Repeat("-", maxWidth+2)

	returnString = append(returnString, topBorder)
	if numLines == 1 {
		currentString = fmt.Sprintf("< %s >", lines[0])
		returnString = append(returnString, currentString)
	} else {
		currentString = fmt.Sprintf("/ %s \\", lines[0])
		returnString = append(returnString, currentString)

		for i := 1; i < numLines-1; i++ {
			currentString = fmt.Sprintf("| %s |", lines[i])
			returnString = append(returnString, currentString)
		}

		currentString = fmt.Sprintf("\\ %s /", lines[numLines-1])
		returnString = append(returnString, currentString)
	}

	returnString = append(returnString, bottomBorder)
	return strings.Join(returnString, "\n")
}

func main() {
	var lines []string
	var animal string

	// allow the user to pass in a `-f <animal>` to print with, defaults to cow
	flag.StringVar(&animal, "f", "cow", "Animal name you wish to print. Valid names are:\n'cow' | 'stegosaurus' | 'trex'\n")
	flag.Parse()

	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Printf("There was an error with your command, printing program\nERROR:\n%v", err)
		os.Exit(-1)
	}

	// gets the file mode (Ex: Dcrw-rw-rw-) and filters out specifically the ModeChar Device (c in the prev example)
	// if this is ran as a piped function, it will be empty (0), but if we dont, it will be a value.
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo <String> | ./cowsay")
		os.Exit(0)
	}

	// take in the input from the stdin
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line string = scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("There was an issue reading your string, make sure to echo it out like so:\necho <string> | cowsay")
		log.Fatal(err)
		os.Exit(-1)
	} else if len(lines) == 0 {
		fmt.Println("You didn't pass in a printed string, try echoing it out like so:\necho <string> | cowsay")
		os.Exit(0)
	}

	lines = tabs_to_spaces(lines)
	maxwidth := calculate_max_width(lines)
	messages := normalize_string_length(lines, maxwidth)
	balloon := build_bubble(messages, maxwidth)
	fmt.Println(balloon)
	print_animal(animal)
	fmt.Println()
}

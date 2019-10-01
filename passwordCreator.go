package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	passwordLength = flag.Int("length", 8, "Length of the generated password")
	passwordCount  = flag.Int("count", 1, "Determine how many passwords to generate")
	file           = flag.String("file", "", "Write passwords to the named file, instead of standard output")

	// Variables that define what characters to use in the password

	lowerCase         bool
	upperCase         bool
	numerals          bool
	specialCharacters bool
	usersCharacters   string
)

func usage() {
	command := os.Args[0]

	fmt.Fprintf(os.Stderr,
		`Usage: %s [all] [alphanum] [lower] [upper] [numbers] [special] [own=CHARACTERS]
%s requires at least one of the following subcommands to specify what characters
may be used in the password:
  all: 		Equivalent to 'alphanum special'
  alphanum:	Equivalent to 'lower upper numbers'
  lower: 	Use lower-case letters
  upper: 	Use upper-case letters
  numbers: 	Use digits
  special: 	Use special characters
  own: 		Specifies a custom set of characters to use

'all', 'alphanum', 'lower', 'upper', 'numbers', and 'special' may be followed by
'=f' to nullify that character set.
Options:
`,
		command, command)
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	flag.Parse()

	fmt.Println("Length:", *passwordLength)
	fmt.Println("Count: ", *passwordCount)
	fmt.Println("File: ", *file)
	fmt.Println(flag.Args())

	for _, value := range flag.Args() {
		splitFlag := strings.SplitN(value, "=", 2)

		switch splitFlag[0].lowerCase() {
		case "all":
			fmt.Printf("you said all :)")

		case "alphanum":

		case "lower":

		case "numbers":

		case "special":

		case "own":

		default:
			printError(fmt.Errorf("Invalid argument: %s", splitFlag[0]))

		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "Error: "+err.Error())
}

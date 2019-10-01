package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unsafe"
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

func printBools() {
	fmt.Println("lowerCase:", lowerCase)
	fmt.Println("upperCase:", upperCase)
	fmt.Println("numerals:", numerals)
	fmt.Println("specialCharacters:", specialCharacters)
	fmt.Println("usersCharacters:", usersCharacters)
}

func setBool(args []string, boolToSet ...*bool) {
	setFlag := true
	fmt.Println(args)
	// Parameters can be passed as lower or lower=true or lower=false

	if len(args) > 1 {
		var err error
		setFlag, err = strconv.ParseBool(args[1])
		if err != nil {
			printError(err)
			os.Exit(1)
		}
	}

	for _, value := range boolToSet {
		*value = setFlag
	}
	printBools()
}

func main() {
	flag.Usage = usage

	flag.Parse()

	fmt.Println("Length:", *passwordLength)
	fmt.Println("Count: ", *passwordCount)
	fmt.Println("File: ", *file)

	// Parameters can be passsed as lower or lower=true or lower=flase
	// Separate into two strings "lower" and "true"
	for _, value := range flag.Args() {
		optionalFlags := strings.SplitN(value, "=", 2)

		switch strings.ToLower(optionalFlags[0]) {
		case "all":
			setBool(optionalFlags, &lowerCase, &upperCase, &numerals, &specialCharacters)

		case "alphanum":
			setBool(optionalFlags, &lowerCase, &upperCase, &numerals)

		case "lower":
			setBool(optionalFlags, &lowerCase)

		case "upper":
			setBool(optionalFlags, &upperCase)

		case "numbers":
			setBool(optionalFlags, &numerals)

		case "special":
			setBool(optionalFlags, &specialCharacters)

		case "own":

		default:
			printError(fmt.Errorf("Invalid argument: %s", optionalFlags[0]))
		}

		var output *os.File
		// var fileError error

		if *file != "" {
			fmt.Println("file is", *file)
			output = os.Stdout
		} else {
			output = os.Stdout
		}
		fmt.Println("Password to output to", output)

		creator, err := NewCreator(output, lowerCase, upperCase, numerals, specialCharacters, "")

		fmt.Println("Creator and error is ", creator, err)
		fmt.Println("size of output is ", unsafe.Sizeof(output))
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "Error: "+err.Error())
}

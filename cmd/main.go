package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Subcommands
	hashCommand := flag.NewFlagSet("hash", flag.ExitOnError)
	verifyCommand := flag.NewFlagSet("verify", flag.ExitOnError)

	// Adding a new choice for --metric of 'substring' and a new --substring flag
	hashPasswordPtr := hashCommand.String("password", "", "Text to parse. (Required)")

	verifyHashPtr := verifyCommand.String("hash", "", "Hash {chars|words|lines|substring}. (Required)")
	verifyPasswordPtr := verifyCommand.String("password", "", "Hash {chars|words|lines|substring}. (Required)")

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("list or count subcommand is required")
		os.Exit(1)
	}

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	case "hash":
		hashCommand.Parse(os.Args[2:])
	case "verify":
		verifyCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
	if hashCommand.Parsed() {
		// Required Flags
		if *hashPasswordPtr == "" {
			hashCommand.PrintDefaults()
			os.Exit(1)
		}
		fmt.Printf(generateHash([]byte(*hashPasswordPtr)))
	}

	if verifyCommand.Parsed() {
		// Required Flags
		if *verifyHashPtr == "" {
			hashCommand.PrintDefaults()
			os.Exit(1)
		}
		if *verifyPasswordPtr == "" {
			hashCommand.PrintDefaults()
			os.Exit(1)
		}
		fmt.Printf(verifyHash([]byte(*verifyHashPtr), []byte(*verifyPasswordPtr)))
	}
}

func generateHash(password []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(password, 10)
	return string(hash)
}

func verifyHash(hash, password []byte) string {
	if err := bcrypt.CompareHashAndPassword(hash, password); err != nil {
		return "false"
	}
	return "true"
}

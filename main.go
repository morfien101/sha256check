package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	stdinHint = "-"
)

var (
	version = "0.0.2"

	helpFlag    = flag.Bool("h", false, "Shows the help menu")
	versionFlag = flag.Bool("v", false, "Show the version of the application")
	verboseFlag = flag.Bool("verbose", false, "Enables verbose logging")

	filePathFlag = flag.String("f", "", "The file that you want to check. Use '-' for STDIN")
	sumFlag      = flag.String("s", "", "SHA256 value that you want to assert against")
	displayFlag  = flag.Bool("d", false, "Display the SHA256 sum value of the passed in file path")
)

func main() {
	digestFlags()
	err := validateFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileBytes, err := readFile(*filePathFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sum := getSHA256(fileBytes)

	if *displayFlag {
		displaySHA(sum)
	}
	exitcode := 0
	if *sumFlag != "" {
		if compairSHAs(convertSumToString(sum), *sumFlag) {
			if *verboseFlag {
				fmt.Println("OK")
			}
		} else {
			if *verboseFlag {
				fmt.Println("FAIL")
			}
			exitcode = 1
		}
	}
	os.Exit(exitcode)
}

func digestFlags() {
	// Parse the flags to get the state we need to run in.
	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	if *helpFlag {
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func validateFlags() error {
	errors := []string{}
	if *filePathFlag == "" {
		errors = append(errors, fmt.Sprint("-f can not be blank"))
	}
	if !*displayFlag && *sumFlag == "" {
		errors = append(errors, fmt.Sprint("-s can not be blank if -d is not also set."))
	}
	if len(*sumFlag) > 64 {
		errors = append(errors, fmt.Sprint("-s can not be longer than 64 characters"))
	}

	if len(errors) != 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}
	return nil
}

func readFile(path string) ([]byte, error) {
	if path == stdinHint {
		return ioutil.ReadAll(os.Stdin)
	}
	return ioutil.ReadFile(path)
}

func getSHA256(b []byte) [32]byte {
	return sha256.Sum256(b)
}

func displaySHA(sum [32]byte) {
	fmt.Printf("%x\n", sum)
}

func compairSHAs(input, expected string) bool {
	return input == expected
}

func convertSumToString(sum [32]byte) string {
	return fmt.Sprintf("%x", sum)
}

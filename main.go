package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

var attached bool

func main() {
	// since this only reproduces in some terminals, debugging requires the
	// ability to attach to a separately launched executable
	if os.Getenv("DEBUGME") != "" {
		fmt.Println("Waiting for debugger ...")
		// change this var in the debugger once attached
		for !attached {
			time.Sleep(time.Second / 2)
		}
	} else {
		// pause to make it easy to separate startup from prompting when run under strace
		time.Sleep(time.Second / 2)
	}

	var result bool
	if err := survey.AskOne(&survey.Confirm{Message: "Does it work?"}, &result); err != nil {
		fmt.Println("failed:", err)
	} else {
		fmt.Println("success:", result)
	}
}

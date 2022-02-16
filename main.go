package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	var result bool
	if err := survey.AskOne(&survey.Confirm{Message: "Does it work?"}, &result); err != nil {
		fmt.Println("failed:", err)
	} else {
		fmt.Println("success:", result)
	}
}

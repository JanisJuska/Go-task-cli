package utils

import (
	"fmt"
	"strings"
)

func ShowHelp() error {
	readmeData, err := OpenAndReadFile("./README.md")
	if err != nil {
		return err
	}

	readme := string(readmeData)

	readme = strings.ReplaceAll(readme, "task-cli", "")
	readme = strings.ReplaceAll(readme, "```", "")
	readme = strings.ReplaceAll(readme, "bash", "")

	// usageIndex := strings.Index(readme, "Usage")
	// fmt.Println(usageIndex)
	// usageIndex = strings.Index(readme, "Notes")
	// fmt.Println(usageIndex)

	manualString := readme[423:845]

	fmt.Println(manualString)

	return nil

}

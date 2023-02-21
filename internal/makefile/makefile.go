package makefile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type target struct {
	Name        string
	Description string
}

// GetMakeTargets reads Makefile and returns targets with description
func GetMakeTargets() []target {
	targets := []target{}

	readFile, err := os.Open("Makefile")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		before, after, found := strings.Cut(line, ":")
		if found {
			name := strings.TrimSpace(before)
			if strings.HasPrefix(name, ".") {
				continue
			}

			_, description, _ := strings.Cut(after, "##")
			targets = append(targets, target{
				Name:        name,
				Description: strings.TrimSpace(description),
			})

		}
	}

	return targets
}

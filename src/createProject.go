package src

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func NewProject() {
	s := "oc new-project jenkins-test"
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to create new project: %v", err)
	}

	fmt.Printf("%s\n", b)
}

package src

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func JenkinsEphemeral() {

	//s := "oc create -f " + template
	p := "oc process jenkins-ephemeral-monitored -n jenkins-test1 | oc create -f -"
	//args := strings.Split(s, " ")
	pargs := strings.Split(p, " ")
	//cmd := exec.Command(args[0], args[1:]...)

	//b, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Printf("Failed to create new app from ephemeral template : %v", err)
	//}
	//
	//fmt.Printf("%s\n", b)

	cmd := exec.Command(pargs[0], pargs[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed : %v", err)
	}

	fmt.Printf("%s\n", b)

}

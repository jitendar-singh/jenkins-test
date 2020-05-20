package main

import (
	"github.com/jitendar-singh/jenkins-test/src"
)

//var testn = "1"

func main() {
	//ephemeralTemplate := "/home/jsingh/go/src/github.com/jitendar-singh/jenkins-test/templates/jenkins-ephemeral-monitored.json"
	src.MavenAgent()
	src.NewProject()
	src.JenkinsEphemeral()
	//src.EphemerialTemplate(template)

}

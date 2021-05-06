package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var label string
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	generateDashboardHCL(label, input)
}

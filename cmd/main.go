package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var snakeCaseRE = regexp.MustCompile("^[a-z]+(_[a-z]+)*$")

func main() {
	labelPtr := flag.String("l", "generated_dashboard", "the resource label")
	inFilePtr := flag.String("i", "", "the input file")

	flag.Parse()

	if ok := snakeCaseRE.MatchString(*labelPtr); !ok {
		log.Fatalf("label must be formatted with snake case: %s", *labelPtr)
	}

	var input []byte
	var err error
	if *inFilePtr != "" {
		input, err = ioutil.ReadFile(*inFilePtr)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	}

	generateDashboardHCL(*labelPtr, input)
}

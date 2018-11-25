package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Usage = toolUsage
	action := "generateNetwork"
	genExample := flag.Bool("ex", false, "Generate example chaincode")
	flag.Parse()
	if *genExample {
		action = "generateExample"
	}

	args := flag.Args()
	fmt.Printf("Starting the application.... \n")
	switch action {
	case "generateNetwork":
		if len(args) == 0 {
			flag.Usage()
			os.Exit(1)
		}
		fmt.Printf("Reading the input .... %v\n", args[0])
		configBytes, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println("Error in reading input json")
			os.Exit(2)
		}

		GenerateNetworkItems(configBytes, ".")
	case "generateExample":
		GenerateExampleCC("v1", "./")
	default:
		flag.Usage()
	}

}

var toolUsage = func() {
	fmt.Printf("Usage : fabricnetgen [flags] <network json file >\n")
	fmt.Printf("Flags : -ex Generates an example chaincode in the current directory\n")

}

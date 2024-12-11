package toolkit

import "fmt"

var CLI_ARGS []string = []string{}

func GetCliArgs()[]string{// Returns the CLI_ARGS 
	return CLI_ARGS
}

func Update(){
	fmt.Println("updating vint...")
}
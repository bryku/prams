package main

import(
	"fmt"
	"github.com/bryku/prams"
)

func main(){

	help, isHelp := prams.Get("-help")
	// help = the follow str after the arg
	// isHelp = if the arg is found
		if(isHelp == true){
			// do something
		}

	fmt.Print("Help = \"",help,"\"\n")
	fmt.Print("isHelp = \"",isHelp,"\"\n")

}

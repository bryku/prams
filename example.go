package main

import(
	"fmt"
	"github.com/bryku/terminal/terminal"
)

func main(){
	help, isHelp := terminal.Arg("-help")
	// help = the follow str after the arg
	// isHelp = if the arg is found
		if(isHelp == true){
			// do something
		}

	fmt.Print("Help = \"",help,"\"\n")
	fmt.Print("isHelp = \"",isHelp,"\"\n")

	terminal.Size("100","40")

}
/* Examples

    ./example
    Help = ""
    isHelp = false

    ./example -help
    Help = ""
    isHelp = true

    ./example -help helloworld
    Help = "helloworld"
    isHelp = true
    
*/

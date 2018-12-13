# prams
handles parameters in terminal, similar to golangs flags, but very small and compact

### Example

	package main
	import(
		"fmt"
		"github.com/bryku/prams"
	)

	func main(){
		help, isHelp := prams.Get("-help")

		fmt.Print("Help = \"",help,"\"\n")
		fmt.Print("isHelp = \"",isHelp,"\"\n")
	}


### Usages  './main'

**Output:** 

	Help = ""
	isHelp = false

### Usages  './main -help'

**Output:** 

	Help = ""
	isHelp = true

### Usages  './main -help me'

**Output:** 

	Help = "me"
	isHelp = true



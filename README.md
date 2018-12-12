# terminal
handles prams in terminal, similar to golangs flags, but 20 times smaller,

### terminal.Arg()

---
func main(){
	help, isHelp := terminal.Arg("-help")
		if(isHelp == true){
			// do something
		}

	fmt.Print("Help = \"",help,"\"\n")
	fmt.Print("isHelp = \"",isHelp,"\"\n")
}
---

    ./example
    Help = ""
    isHelp = false

    ./example -help
    Help = ""
    isHelp = true

    ./example -help helloworld
    Help = "helloworld"
    isHelp = true
    
### **CODE** terminal.Size()

---
func main(){
	terminal.Size("100","40")
}
---

Sets terminal gui size (cols, rows)

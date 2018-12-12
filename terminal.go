package terminal

import(
	"os"
	"fmt"
	"runtime"
)

func Arg(str string) (string, bool){
	a := os.Args[1:];
	argsString := ""
	argsBool := false
	for key, _ := range a {
		if(a[key] == str){
			argsBool = true
			if(len(a) > key + 1){
				argsString = a[key + 1]
			}else{
				argsString = ""
			}
		}
	}
	return argsString, argsBool
}

func Size(w string, h string){
	switch(runtime.GOOS){
		case "windows":
			fmt.Println("terminal.Size: Windows not supported")
		case "darwin":
			fmt.Println("terminal.Size: Unix not supported")
		case "linux":
			//args := "'\033[8;"+ h +";"+ w +"t'"
			args:= "\033[8;"+h+";"+w+"t"
			fmt.Print(args)
	}
}

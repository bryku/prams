package prams
import("os")
func Get(str string) (string, bool){
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

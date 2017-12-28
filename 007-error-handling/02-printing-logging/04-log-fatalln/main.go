package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened err",err)
		//log.Println("err happened err",err)
		log.Fatalln(err)
		//panic(err)
	}
}
func foo(){
	fmt.Println("when the os.Exit() file is called, deffered functions don't run")
}

/*
... the Fatal functions call os.Exit(1) after writing the log message...
*/
// fatalln is equivalent to Println() followed by a call to os.Exit(1).

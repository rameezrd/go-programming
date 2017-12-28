package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened err",err)
		//log.Println("err happened err",err)
		//log.Fatalln(err)
		panic(err)
	}
}

/*
... the Fatal functions call os.Exit(1) after writing the log message...
*/
// fatalln is equivalent to Println() followed by a call to os.Exit(1).

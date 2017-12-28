package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened err",err)
		log.Println("err happened err", err)
		//log.Fatalln(err)
		//panic(err)
	}
	defer f2.Close()
	fmt.Println("check the log.txt file in the directory")
}

/*
... the Fatal functions call os.Exit(1) after writing the log message...
*/
// fatalln is equivalent to Println() followed by a call to os.Exit(1).

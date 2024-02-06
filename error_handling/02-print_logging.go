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
		// By default, these two are the same
		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err) // Like Println, but also calls os.Exit(1) after writing
        //      log.Panicln(err) // Like Println, but also calls panic() after writing
		//		panic(err)
        /* deferred functions are still called after panicking,
        and you can recover */
	}
	
	//However, you can set log to use different outputs. i.e:
	//Create file and check for an error
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close() // make sure you close file at end of main
	log.SetOutput(f) // Send output to f
	f2, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}

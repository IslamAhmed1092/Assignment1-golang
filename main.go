package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("ExampleIn.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	inputString := string(file)

	inputArray := strings.Split(inputString, " ")

	for i := 0; i < len(inputArray); i++ {
		fmt.Println(inputArray[i], " ")
	}


}

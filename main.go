package main

import (
	"bufio"
	// "fmt"
	"strconv"

	// "io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

var resultMap map[string]int
var mu sync.Mutex
var wg sync.WaitGroup

func computerWordCounts(textArray []string) {
	defer wg.Done()

	routineResult := make(map[string]int)
	for _, v := range textArray {
		routineResult[v] += 1
	}

	reducer(routineResult)
}

func reducer(routineResult map[string]int) {
	mu.Lock()

	for key, value := range routineResult {
		resultMap[key] += value
	}

	mu.Unlock()
}

func sortAndWrite() {

	type pair struct {
		Key   string
		Value int
	}

	var pairsArray []pair
	for key, value := range resultMap {
		pairsArray = append(pairsArray, pair{key, value})
	}

	sort.Slice(pairsArray, func(i, j int) bool {
		if pairsArray[i].Value == pairsArray[j].Value {
			return pairsArray[i].Key < pairsArray[j].Key
		}
		return pairsArray[i].Value > pairsArray[j].Value
	})

	file, err := os.Create("output.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for _, pair := range pairsArray {
		file.WriteString(pair.Key + " : " + strconv.Itoa(pair.Value) + " \n")
	}
}

func main() {

	resultMap = make(map[string]int)

	file, err := os.Open("ExampleIn.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputString := ""

	for scanner.Scan() {
		inputString += scanner.Text() + " "
	}

	inputString = strings.ToLower(string(inputString[:len(inputString)-1]))

	inputArray := strings.Split(inputString, " ")

	numberOfWords := len(inputArray)

	for i := 0; i < 5; i++ {
		var startIndex int = i / 5 * numberOfWords
		var endIndex int = (i + 1) / 5 * numberOfWords
		if i == 4 {
			endIndex = numberOfWords
		}
		wg.Add(1)
		go computerWordCounts(inputArray[startIndex:endIndex])
	}

	wg.Wait()

	sortAndWrite()
}

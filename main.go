package main

import (
	"bufio"
	"strconv"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

type WordsCount struct {
	resultMap map[string]int
	mu sync.Mutex
	wg sync.WaitGroup

}

func (wc* WordsCount) computerWordCounts(textArray []string) {

	routineResult := make(map[string]int)
	for _, v := range textArray {
		routineResult[v] += 1
	}

	wc.reducer(routineResult)
	wc.wg.Done()
}

func (wc* WordsCount) reducer(routineResult map[string]int) {
	wc.mu.Lock()

	for key, value := range routineResult {
		wc.resultMap[key] += value
	}

	wc.mu.Unlock()
}

func sortAndWrite(resultMap map[string]int) {

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

	wc := WordsCount{resultMap: make(map[string]int)}

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

	inputString = strings.ToLower(inputString[:len(inputString)-1])

	inputArray := strings.Split(inputString, " ")

	numberOfWords := len(inputArray)

	for i := 0; i < 5; i++ {
		var startIndex int = i / 5 * numberOfWords
		var endIndex int = (i + 1) / 5 * numberOfWords
		if i == 4 {
			endIndex = numberOfWords
		}
		wc.wg.Add(1)
		go wc.computerWordCounts(inputArray[startIndex:endIndex])
	}

	wc.wg.Wait()

	sortAndWrite(wc.resultMap)
}

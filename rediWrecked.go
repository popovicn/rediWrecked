package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var outputMutex sync.Mutex

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func check(err error) {
	if err != nil {
		exit(err.Error())
	}
}

func formatOutput(url, status, redirect string) string {
	return url + "\t" + status + "\t" + redirect
}

func writeResultLine(line, outputFile string) {
	outputMutex.Lock()
	defer outputMutex.Unlock()
	file, _ := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, _ = fmt.Fprintln(writer, line)
	_ = writer.Flush()
}

func processUrl(url string, outputFile string, outputChannel chan string, wg *sync.WaitGroup){
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	response, err := client.Do(req)
	if err == nil {
		output := formatOutput(url, strconv.Itoa(response.StatusCode), response.Header.Get("Location"))
		fmt.Println(output)
		writeResultLine(output, outputFile)
	}
	wg.Done()
}

func main() {
	var inputFilePath = flag.String("i", "","Input file path (required)")
	var parallelism = flag.Int("p", 40, "Parallelism")
	var outputFilePath = flag.String("o", "output.txt","Output file path")
	flag.Parse()

	if *inputFilePath == "" {
		exit("Error: input file is required.")
	}
	_, err := os.Create(*outputFilePath)
	check(err)

	inputFile, err := os.Open(*inputFilePath)
	check(err)
	defer inputFile.Close()

	var wg sync.WaitGroup
	parallelismBlocker := make(chan struct{}, *parallelism)
	outputChannel := make(chan string, *parallelism + 1)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		parallelismBlocker <- struct{}{}
		wg.Add(1)
		go processUrl(scanner.Text(), *outputFilePath, outputChannel, &wg)
		<-parallelismBlocker
	}
	wg.Wait()
}
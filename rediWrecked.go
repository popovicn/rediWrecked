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

var simpleOutput = false

var banner = "\033[31;1m"  + `
    ____ ____ ___  _ _ _ _ ____ ____ ____ _  _ ____ ___  
    |__/ |___ |  \ | | | | |__/ |___ |    |_/  |___ |  \ 
    |  \ |___ |__/ | |_|_| |  \ |___ |___ | \_ |___ |__/ 

` + "\033[0m"

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

func formatFileOutput(url, status, redirect string) string {
    return url + "\t" + status + "\t" + redirect
}

func formatCliOutput(url, status, redirect string) string {
    color := ""
    statusCode, err := strconv.Atoi(status)
    if err == nil {
        switch statusCode / 100 {
        case 2: color = "\033[32;1m"
        case 3: color = "\033[34;1m"
        case 4: color = "\033[31;1m"
        case 5: color = "\033[31;2m"
        }
    }
    output := color + status + "\033[0m " + url
    if redirect != "" {
        output += "\033[34;1m → " + redirect + "\033[0m"
    }
    return output
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

func processUrl(url string, outputFile string, outputChannel chan string, wg *sync.WaitGroup, queue chan struct{}){
    client := &http.Client{
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }
    req, err := http.NewRequest("GET", url, nil)
    response, err := client.Do(req)
    if err == nil {
        output := formatFileOutput(url, strconv.Itoa(response.StatusCode), response.Header.Get("Location"))
        if simpleOutput {
            fmt.Println(output)
        } else {
            fmt.Println(formatCliOutput(url, strconv.Itoa(response.StatusCode), response.Header.Get("Location")))
        }
        writeResultLine(output, outputFile)
    }
    <- queue
    wg.Done()

}

func main() {
    var inputFilePath = flag.String("i", "","Input file path (required)")
    var parallelism = flag.Int("p", 50, "Parallelism")
    var outputFilePath = flag.String("o", "output.txt","Output file path")
    var simple = flag.Bool("s", false, "Simple CLI (tab separated)")
    flag.Parse()

    if *simple {
        simpleOutput = true
    } else {
        fmt.Println(banner)
    }
    if *inputFilePath == "" {
        exit("Error: input file is required.")
    }
    _, err := os.Create(*outputFilePath)
    check(err)

    inputFile, err := os.Open(*inputFilePath)
    check(err)
    defer inputFile.Close()

    var wg sync.WaitGroup
    queue := make(chan struct{}, *parallelism)
    outputChannel := make(chan string, *parallelism + 1)
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        queue <- struct{}{}
        wg.Add(1)
        go processUrl(scanner.Text(), *outputFilePath, outputChannel, &wg, queue)
    }
    wg.Wait()
}
package main

import (
    "fmt"
    "os"
    "strconv"
    "regexp"
)

func getInputData(path string) []byte {
    data, err := os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
        return []byte{}
    }

    return data
}

func getOperations(data string) []string{
    re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
    return re.FindAllString(data,-1)
}

func getComplexOperations(data string) []string{
    re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don\'t\(\)`)
    return re.FindAllString(data,-1)
}

func mul(op string) int {

    re := regexp.MustCompile(`\d+`)
    numbers := re.FindAllString(op,-1)

    if len(numbers) != 2 {
        fmt.Println("Error in mul function")
    }

    a, _:= strconv.Atoi(numbers[0])
    b, _:= strconv.Atoi(numbers[1])

    
    return a * b
}

func runOperations(data []string) int {
    result := 0

    for _, value := range data {
        result += mul(value)
    }

    return result
}

func runComplexOperations(data []string) int {
    result := 0
    flag := true

    for _, value := range data {
        if(value == "do()") { flag = true; continue}
        if(value == "don't()") { flag = false; continue }
        if(flag) {
            result += mul(value)
        }
    }
    return result
}

func main() {
    data := getInputData(os.Args[1])
    parsedData := string(data[:])

    operations := getOperations(parsedData)
    value := runOperations(operations)
    fmt.Printf("Result: %d\n",value)

    complexOperations := getComplexOperations(parsedData)
    complexValue := runComplexOperations(complexOperations)
    fmt.Printf("Complex result: %d\n",complexValue)

}

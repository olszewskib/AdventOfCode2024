package main

import (
    "fmt"
    "strings"
    "strconv"
//    "sort"
    "math"
    "os"
)

type reportValidator func([]int) bool

func getInputData(path string) []byte {
    data, err := os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
        return []byte{}
    }

    return data
}

func parseData(data []byte) [][]int {
    var reports [][]int
    var report string

    for _, value := range data {
        if string(value) == "\n" {

            numbers := strings.Split(report, " ")
            var levels  = make([]int, len(numbers))

            for index, value  := range numbers {
                number, err  := strconv.Atoi(value)
                if err != nil {
                    fmt.Println(err)
                    return [][]int{}
                }
                levels[index] = number
            }

            reports = append(reports, levels)
            report = ""

        } else {
            report += string(value)
        }
    }

    return reports
}


func runChecks(x int, y int, increasing bool) bool {

    if math.Abs(float64(x - y)) > 3 {
        return false
    }

    if x == y {
        return false
    }

    check := x < y
    if check != increasing {
        return false
    }

    return true
}

func validateReport(report []int) bool {
    var increasing bool = report[0] < report[1]

    for i:=0; i < len(report)-1; i++ {

        if !runChecks(report[i], report[i+1], increasing) {
            return false
        }
    }

    return true
}

func validateReportWithOneError(report []int) bool {
    var increasing bool = report[0] < report[1]
    var errors int = 0

    for i:=0; i < len(report)-1; i++ {

        if !runChecks(report[i], report[i+1], increasing) {
            errors++

            // To be fixed
            // Check if we can delete the error, if not return false
            if errors == 1 && i < len(report) - 2 {
                if !runChecks(report[i], report[i+2], increasing) && !runChecks(report[i+1], report[i+2], increasing) {
                    return false
                }
            }

            if errors == 1 && i == len(report) - 2 {
                if !runChecks(report[i-1], report[i], increasing) && !runChecks(report[i-1], report[i+1], increasing) {
                    return false
                }
            }

            i++
            continue
        }

        if errors > 1 {
            return false
        }
    }

    return true
}

func printReport(report []int) {
    for _, value := range report {
        print(value)
        print(" ")
    }
}


func validateReports(reports [][]int, validate reportValidator) int {
    validReports := 0

    for _, report := range reports {
        if validate(report) {
            validReports++
        }
    }

    return validReports
}

func main() {

    inputFile := os.Args[1]
    data := getInputData(inputFile)
    parsedData := parseData(data)
    validReports := validateReports(parsedData, validateReportWithOneError)
    fmt.Printf("Valid reports: %d",validReports)

}

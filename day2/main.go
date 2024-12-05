package main

import (
    "fmt"
    "strings"
    "strconv"
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

func isGradual(x int, y int) bool {

    if math.Abs(float64(x - y)) > 3 {
        return false
    }

    if x == y {
        return false
    }

    return true
}

func getDirection(first int, second int) int {
    if first == second {
        return 0
    }

    if first < second {
        return 1
    }

    return -1
}

func isMonothonic(first int, second int, direction int) bool {
    return getDirection(first, second) == direction
}

func validateReport(report []int) bool {
    var direction int = getDirection(report[0], report[1])

    if direction == 0 {
        return false
    }

    for i:=0; i < len(report)-1; i++ {

        if !isGradual(report[i], report[i+1]) || !isMonothonic(report[i], report[i+1], direction) {
            return false
        }
    }

    return true
}

func fixReport(report []int, index int) []int{
    if index == len(report) - 1 {
        return report[:index]
    }

    fixedReport := []int{}
    for it, value := range report {
        if it == index {
            continue
        }
        fixedReport = append(fixedReport, value)
    }
    return fixedReport
}

func validateReportWithOneError(report []int) bool {

    if validateReport(report) {
                return true
    }

    for i:=0; i < len(report)-1; i++ {


        fixedReportLeft := fixReport(report, i)
        left := validateReport(fixedReportLeft)

        fixedReportRight := fixReport(report, i+1)
        right := validateReport(fixedReportRight)


        if left || right {
            return true
        }

    }

    return false 
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
    validReports := validateReports(parsedData, validateReport)
    fmt.Printf("Valid reports: %d\n",validReports)
    validReportsWithOneError := validateReports(parsedData, validateReportWithOneError)
    fmt.Printf("Valid reports with error: %d\n",validReportsWithOneError)

}

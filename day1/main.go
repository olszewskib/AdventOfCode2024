package main

import (
    "fmt"
    "strings"
    "strconv"
    "sort"
    "math"
    "os"
)

func getInputData(path string) []byte {

    data, err := os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
        return []byte{}
    }

    return data
}

func parseData(data []byte) [2][]int {

    var first []int
    var second []int

    var i int = 0

    for i < len(data) {
        numbers := ""
        for string(data[i]) != "\n" {
            numbers += string(data[i])
            i++
        }

        data := strings.Split(numbers, "   ")
        left, err := strconv.Atoi(data[0])
        if err != nil {
            fmt.Println(err)
            return [2][]int{}
        }
        right, err := strconv.Atoi(data[1])
        if err != nil {
            fmt.Println(err)
            return [2][]int{}
        }

        first = append(first, left)
        second = append(second, right)

        i++
    }

    return [2][]int{first, second}
}

func sortLists(parsedData [2][]int) {

    sort.Ints(parsedData[0])
    sort.Ints(parsedData[1])

// Alternative sorting when when dealing with slices of other types
//
//    sort.Slice(parsedData[0], func(i, j int) bool {
//        left := parsedData[0][i]
//        right := parsedData[0][j]
//        return left < right
//    })
//
//    sort.Slice(parsedData[1], func(i, j int) bool {
//        left := parsedData[1][i]
//        right := parsedData[1][j]
//        return left < right
//    })
}

func sumLists(data [2][]int) int {

    sum := 0
    for i := 0; i < len(data[0]); i++ {
        left := data[0][i]
        right := data[1][i]
        distance := math.Abs(float64(left - right))
        sum += int(distance)
    }

    return sum
}

func score(data []int, number int) int {
    score := 0
    for _, value := range data {
        if value == number {
            score++
        }
    }
    return score
}

func similarityScore(data [2][]int) int {

    result := 0

    for _, value := range data[0] {
        digitScore := score(data[1], value)
        digitScore *= value
        result += digitScore
    }

    return result
}

func main() {

    input := "input.txt"

    data := getInputData(input)
    if len(data) == 0 {
        fmt.Println("File read error")
        return
    }

    parsedData := parseData(data)

    // Part 1
    sortLists(parsedData)
    sum := sumLists(parsedData)

    // Part 2
    score := similarityScore(parsedData)

    fmt.Println("Part 1: ", sum)
    fmt.Println("Part 2: ", score)

}

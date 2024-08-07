package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
    "io"
)

func matchStrings(stringList []string, queries []string) []int32 {
    var results []int32
    for q := 0; q < len(queries); q++ {
        results = append(results, 0)

        for s := 0; s < len(stringList); s++ {
            if queries[q] == stringList[s] {
                results[q]++
            }
        }
    }

    return results
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(os.Stdout, 16 * 1024 * 1024)

    stringListCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var stringList []string

    for i := 0; i < int(stringListCount); i++ {
        stringListItem := readLine(reader)
        stringList = append(stringList, stringListItem)
    }

    queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var queries []string

    for i := 0; i < int(queriesCount); i++ {
        queriesItem := readLine(reader)
        queries = append(queries, queriesItem)
    }

    res := matchStrings(stringList, queries)

    for i, resItem := range res {
        fmt.Fprintf(writer, "%d", resItem)

        if i != len(res) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumLoss' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts LONG_INTEGER_ARRAY price as parameter.
 */

func minimumLoss(price []int64) int32 {
    var priceInts [][]int
    l := len(price)
    for i := 0; i < l; i++ {
        var pair []int
        pair = append(pair, int(price[i]))
        pair = append(pair, i)
        priceInts = append(priceInts, pair)
    }
    
    sort.Slice(priceInts, func(i, j int) bool {
        for x := range priceInts[i] {
            if priceInts[i][x] == priceInts[j][x] {
                continue
            }
            return priceInts[i][x] > priceInts[j][x]
        }
        return false
    })
        
    miniLoss := priceInts[0][0]
    for i := 0; i < l - 1; i++ {
        if priceInts[i][1] < priceInts[i + 1][1] {
            loss := priceInts[i][0] - priceInts[i + 1][0]
            if miniLoss > loss {
                miniLoss = loss
            }
        }
    }
    
    return int32(miniLoss)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)
    defer stdout.Close()
    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)
    priceTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
    var price []int64

    for i := 0; i < int(n); i++ {
        priceItem, err := strconv.ParseInt(priceTemp[i], 10, 64)
        checkError(err)
        price = append(price, priceItem)
    }

    result := minimumLoss(price)
    fmt.Fprintf(writer, "%d\n", result)
    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Reduces string to either an empty string or a non-reduceable string.
 * Accepts STRING s as parameter.
 * Returns STRING.
 */

func superReducedString(s string) string {
    for i := 0; i < len(s) - 1; i++ {
        if s[i] == s[i + 1] {
            if i == 0 {
                s = s[2:]
            } else if i + 2 == len(s) {
                s = s[:i]
            } else {
                s = s[:i] + s[i+2:]
            }
            i = -1
        }
    }
    
    if s == "" {
        return "Empty String"
    }
    return s
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)
    defer stdout.Close()
    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
    s := readLine(reader)
    result := superReducedString(s)
    fmt.Fprintf(writer, "%s\n", result)
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

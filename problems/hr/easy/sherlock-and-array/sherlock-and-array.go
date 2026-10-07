package sherlockandarray

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'balancedSums' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func balancedSums(arr []int32) string {
	n := len(arr)

	if n == 1 {
		return "YES"
	}
	prefix := make([]int32, n)
	suffix := make([]int32, n)

	prefix[0] = arr[0]
	suffix[n-1] = arr[n-1]

	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + arr[i]
	}

	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] + arr[i]
	}

	for i := 0; i < n-1; i++ {
		if i == 0 && suffix[i+1] == 0 {
			return "YES"
		} else if i > 0 && prefix[i-1] == suffix[i+1] {
			return "YES"
		}
	}
	return "NO"
}

/* Codes written below are written by HackerRank */

func Implementation() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var arr []int32

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := balancedSums(arr)

		fmt.Fprintf(writer, "%s\n", result)
	}

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

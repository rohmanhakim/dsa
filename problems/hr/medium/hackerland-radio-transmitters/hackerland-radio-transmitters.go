package hackerlandradiotransmitters

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'hackerlandRadioTransmitters' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY x
 *  2. INTEGER k
 */

func hackerlandRadioTransmitters(x []int32, k int32) int32 {
	// Write your code here

	n := int32(len(x))

	sorted := slices.Clone(x)

	slices.Sort(sorted)

	i := int32(0)
	total := 0
	for i < n {
		curr := sorted[i]

		idx := int32(sort.Search(len(sorted[i:]), func(j int) bool {
			return sorted[i:][j]-curr > k
		}))
		idx--

		total++

		spot := sorted[i+idx]

		for i < n && sorted[i] <= spot+k {
			i++
		}

	}

	return int32(total)
}

func Implementation() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	xTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var x []int32

	for i := 0; i < int(n); i++ {
		xItemTemp, err := strconv.ParseInt(xTemp[i], 10, 64)
		checkError(err)
		xItem := int32(xItemTemp)
		x = append(x, xItem)
	}

	result := hackerlandRadioTransmitters(x, k)

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

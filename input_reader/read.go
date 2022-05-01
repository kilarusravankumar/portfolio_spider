package input_reader

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var reader *bufio.Reader

func ConsoleInput(label string) string {
	fmt.Println(label)
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}

	inputString, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}
	var replaceString string
	if runtime.GOOS == "windows" {
		replaceString = "\r\n"
	} else {
		replaceString = "\n"
	}
	inputString = strings.Replace(inputString, replaceString, "", -1)
	return strings.TrimSpace(inputString)

}

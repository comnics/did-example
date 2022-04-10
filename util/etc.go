package util

import (
	"bufio"
	"fmt"
	"os"
)

func PressKey(msg string) {
	kbReader := bufio.NewReader(os.Stdin)

	fmt.Println(msg)
	kbReader.ReadString('\n')
}

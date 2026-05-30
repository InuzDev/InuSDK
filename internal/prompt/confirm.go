package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func Confirm(message string) bool {
	fmt.Printf("%s [y/N]: ", message)

	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))

	return answer == "y"
}

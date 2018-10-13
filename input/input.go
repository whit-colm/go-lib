package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Prompt(p string) (t string, err error) {
	fmt.Print(p)
	reader := bufio.NewReader(os.Stdin)
	t, err = reader.ReadString('\n')
	t = strings.TrimSpace(t)
	return t, nil
}

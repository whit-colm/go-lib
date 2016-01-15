package input

import (
	"bufio"
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
	"strings"
)

func Prompt(p string) (t string, err error) {
	//P in Prompt is capitalised because it's a public function,
	//meaning it can be used in any go file if you know where to look.
	fmt.Print(c.G + p)
	reader := bufio.NewReader(os.Stdin)
	t, err = reader.ReadString('\n')
	if err != nil {
		return t, err
	}
	t = strings.TrimSpace(t)
	return t, nil
}

/*
func Ask(q string) string {
	t, err := Prompt(q)
	if err != nil {
		return ""
	}
	return t
}
*/

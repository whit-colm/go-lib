package nil

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "input"
	s "other"
	"strings"
)

//////////////////////////////////////////////////
//This is just a hub for each of the modules//////
//The real programming code is in another castle//
//////////////////////////////////////////////////

func Startup() {
	//the only function in this
	isDone := false
	for isDone == false {
		//So that you don't have to restart the program each time
		fmt.Println(c.CL + c.B00 + "Mathmatic solver program thing.")
		s.Spacer(2)
		//Adds two lines
		fmt.Println(c.B1 + "What are you too lazy to do today?")
		fmt.Println(c.Y + "Select one of the letters in red")
		s.Spacer(1)
	}
}

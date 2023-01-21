package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	prompOptions(b)

	return b
}

func prompOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip)", reader)

	switch opt {
	case "a":
		b.handlePromptAdd(reader)
	case "t":
		b.handlePromptAddTip(reader)
	case "s":
		b.handlePromptSave()
	default:
		fmt.Println("invalid input...")
		prompOptions(b)
	}
}

func main() {
	createBill()
}

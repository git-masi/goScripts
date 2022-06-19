package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const template = `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's write some bugs!")
}
`

func main() {
	// The timestamp is easier to work with in other programs for sorting
	// but the formatted string is human readable
	t := time.Now().Unix()
	tf := time.Now().Format("2006-01-02-03_04_05")
	args := os.Args[1:]
	l := len(args)
	ext := ".go"
	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	var f string

	switch {
	case l == 0:
		f = "scratch_" + fmt.Sprint(t) + ext

	case l == 1 && args[0] == "-tf":
		f = "scratch_" + tf + ext

	case l == 1:
		f = args[0] + ext

	case l == 2 && args[1] == "-t":
		f = args[0] + "_" + fmt.Sprint(t) + ext

	case l == 2 && args[1] == "-tf":
		f = args[0] + "_" + tf + ext

	default:
		log.Fatal("Invalid args")
	}

	n := path + "/" + f

	err = os.WriteFile(n, []byte(template), 0666)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\nNew file: %v\nIn dir: %v", f, n)

	cmd := exec.Command("code", f)

	cmd.Dir = path

	_, err = cmd.Output()

	if err != nil {
		log.Fatal("The file was created but it couldn't be opened in VS Code. Perhaps your path doesn't have \"code\"?")
	}
}

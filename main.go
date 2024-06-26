package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Mehloul-Mohamed/butler/app"
)

func help(f string) {
	switch f {
	case "start":
		fmt.Print("usage: butler start name url token\n\n")
		fmt.Println("positional aguments:")
		fmt.Println("  name\t\t\tCTF name")
		fmt.Println("  url\t\t\tCTF url")
		fmt.Println("  token\t\t\tYour API token")
	case "list":
		fmt.Println("usage: butler list")
	case "attempt":
		fmt.Print("usage: butler attempt id\n\n")
		fmt.Println("positional aguments:")
		fmt.Println("  id\t\t\tChallenge id")
	case "main":
		fmt.Println("usage: butler {start,list,attempt}")
		fmt.Println("positional aguments:")
		fmt.Println("  {start,list,attempt}")
		fmt.Println("\tstart\t\t\tStart a CTF")
		fmt.Println("\tlist\t\t\tShow challenge list")
		fmt.Println("\tattempt\t\t\tAttempt a challenge")
	}
}

func main() {
	if len(os.Args) < 2 {
		help("main")
		return
	}
	switch os.Args[1] {
	case "start":
		if len(os.Args) != 5 {
			help("start")
			return
		}
		name := os.Args[2]
		url := strings.TrimSuffix(os.Args[3], "/")
		token := "token " + os.Args[4]
		if name == "" || url == "" || token == "" {
			help("main")
		}
		err := app.StartCtf(name, url, token)
		if err != nil {
			panic(err)
		}
	case "list":
		if len(os.Args) != 2 {
			help("list")
			return
		}
		err := app.DisplayChallList()
		if err != nil {
			panic(err)
		}
	case "attempt":
		if len(os.Args) != 3 || os.Args[2] == "-h" {
			help("attempt")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic("id must be an integer")
		}
		app.Attempt(id)
	default:
		help("main")
	}
}

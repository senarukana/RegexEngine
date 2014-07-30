package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/senarukana/regex"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	regex.Debug = true
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			fmt.Printf("Good byte %s\n", err.Error())
			os.Exit(0)
		}
		re, err := regex.NewRegex(str[:len(str)-1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		match, _ := rd.ReadString('\n')
		fmt.Println(re.Match(match[:len(match)-1]))
	}
}

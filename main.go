package main

import (
	"fmt"
	"gcom"
	"strings"
)

func main() {
	g := gcom.New()
	outpue := g.GCmd.ExecCommand("git", "remote", "-v")
	fmt.Println(outpue)
	x := strings.Split(outpue, "\n")
	fmt.Println(len(x))
	fmt.Println("1:", x[0])
	fmt.Println("2:", x[1])
	fmt.Println("3:", x[2])
}

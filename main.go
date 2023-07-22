package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func startBrowser(port int) {
	exec.Command("cmd.exe", "/c", fmt.Sprintf("start http://localhost:%d", port)).Start()
}

var (
	mode string
	port int
)

func main() {
	port = 23456
	flag.StringVar(&mode, "m", "", "-m dev")
	flag.Parse()
	fmt.Printf("welcome to use go Searcher!\nlisten at %d\n", port)
	InitDb()
	go RunServer(port, mode)
	if IS_PRO {
		startBrowser(port)
	}
	for {
		fmt.Println("input 1 to open browser")
		var a string
		fmt.Scanln(&a)
		startBrowser(port)
	}
}

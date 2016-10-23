package main

import (
	"chess"
	"runtime/pprof"
	"os"
)

func main() {
	f, _ := os.Create("profiles/profile.data")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	chess.StartServer()
}
package main

import (
	"chess"
	"runtime/pprof"
	"os"
	"flag"
)

var depth = flag.Int("d", 4, "-d=num")

func main() {
	f, _ := os.Create("profiles/profile.data")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	chess.StartServer(int8(*depth))
}
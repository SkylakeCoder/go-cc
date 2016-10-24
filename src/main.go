package main

import (
	"chess"
	"runtime/pprof"
	"os"
	"flag"
	"log"
)

var depth = flag.Int("d", 4, "-d=num")

func main() {
	flag.Parse()
	f, _ := os.Create("profiles/profile.data")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	log.Println("depth=", *depth)
	chess.StartServer(int8(*depth))
}
package utils

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

// https://github.com/DataDog/go-profiler-notes/tree/main
// https://github.com/google/pprof/blob/main/doc/README.md
// https://blog.stackademic.com/profiling-go-applications-in-the-right-way-with-examples-e784526e9481

func CpuProfiling() {
	f, err := os.Create("./profs/cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
}

func MemoryProfiling() {
	f, err := os.Create("./profs/mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()
	runtime.GC()    // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}

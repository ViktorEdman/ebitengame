package main

import (
	"fmt"
	"runtime"
)

func getMemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("Alloc = %.3f MiB", float32(m.Alloc) / 1024 / 1024)
}

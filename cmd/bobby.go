package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/timchurchard/bitcoin-puzzles/internal"
)

func Bobby(out io.Writer) int {
	const (
		usageWorkers = "Number of workers (default runtime.NumCPU())"
		usageStats   = "Number of tries between stats print (default 10,000)"

		defaultStats = 10000
	)
	var (
		workers int
		stats   int
	)

	flag.IntVar(&workers, "workers", runtime.NumCPU(), usageWorkers)
	flag.IntVar(&workers, "w", runtime.NumCPU(), usageWorkers)

	flag.IntVar(&stats, "stats", defaultStats, usageStats)
	flag.IntVar(&stats, "s", defaultStats, usageStats)

	flag.Usage = func() {
		fmt.Fprintf(out, "Usage of %s %s:\n", os.Args[0], os.Args[1])

		flag.PrintDefaults()
	}

	flag.Parse()

	ch := make(chan bool)

	for i := 0; i < workers; i++ {
		go internal.Bobby(ch, i, stats)
	}

	result := <-ch
	if result {
		return 0
	}

	return 1
}

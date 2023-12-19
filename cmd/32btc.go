package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/timchurchard/bitcoin-puzzles/internal"
)

func BTC32(out io.Writer) int {
	const (
		usageWorkers  = "Number of workers (default runtime.NumCPU())"
		usageCount    = "Number of tries on a given random (default 16777216)"
		usageStats    = "Number of tries between stats print (default 1000000)"
		usageMiniMode = "Check 66, 67 & 68-bit only (default true) If false then up to 99 bits is checked."

		defaultCount    = 16777216
		defaultStats    = 1000000
		defaultMiniMode = true
	)
	var (
		workers  int
		count    int
		stats    int
		miniMode bool
	)

	flag.IntVar(&workers, "workers", runtime.NumCPU(), usageWorkers)
	flag.IntVar(&workers, "w", runtime.NumCPU(), usageWorkers)

	flag.IntVar(&count, "count", defaultCount, usageCount)
	flag.IntVar(&count, "c", defaultCount, usageCount)

	flag.IntVar(&stats, "stats", defaultStats, usageStats)
	flag.IntVar(&stats, "s", defaultStats, usageStats)

	flag.BoolVar(&miniMode, "mini", defaultMiniMode, usageMiniMode)
	flag.BoolVar(&miniMode, "m", defaultMiniMode, usageMiniMode)

	flag.Usage = func() {
		fmt.Fprintf(out, "Usage of %s %s:\n", os.Args[0], os.Args[1])

		flag.PrintDefaults()
	}

	flag.Parse()

	ch := make(chan bool)

	for i := 0; i < workers; i++ {
		go internal.BTC32Worker(ch, i, count, stats, miniMode)
	}

	result := <-ch
	if result {
		return 0
	}

	return 1
}

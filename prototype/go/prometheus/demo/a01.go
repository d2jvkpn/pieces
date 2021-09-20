package main

import (
	// "encoding/json"
	"fmt"
	"runtime/metrics"
	"sync"
	"time"
)

var metricsNames = []string{
	"/gc/cycles/automatic:gc-cycles",
	"/gc/cycles/forced:gc-cycles",
	"/gc/cycles/total:gc-cycles",
	"/gc/heap/allocs-by-size:bytes",
	"/gc/heap/allocs:bytes",
	"/gc/heap/allocs:objects",
	"/gc/heap/frees-by-size:bytes",
	"/gc/heap/frees:bytes",
	"/gc/heap/frees:objects",
	"/gc/heap/goal:bytes",
	"/gc/heap/objects:objects",
	"/gc/heap/tiny/allocs:objects",
	"/gc/pauses:seconds",
	"/memory/classes/heap/free:bytes",
	"/memory/classes/heap/objects:bytes",
	"/memory/classes/heap/released:bytes",
	"/memory/classes/heap/stacks:bytes",
	"/memory/classes/heap/unused:bytes",
	"/memory/classes/metadata/mcache/free:bytes",
	"/memory/classes/metadata/mcache/inuse:bytes",
	"/memory/classes/metadata/mspan/free:bytes",
	"/memory/classes/metadata/mspan/inuse:bytes",
	"/memory/classes/metadata/other:bytes",
	"/memory/classes/os-stacks:bytes",
	"/memory/classes/other:bytes",
	"/memory/classes/total:bytes",
	"/sched/goroutines:goroutines",
	"/sched/latencies:seconds",
}

func main() {
	const nGo = "/sched/goroutines:goroutines"
	// A slice for getting metric samples
	getMetric := make([]metrics.Sample, 1)
	getMetric[0].Name = nGo

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(4 * time.Second)
		}()

		// Get actual data
		metrics.Read(getMetric)
		if getMetric[0].Value.Kind() == metrics.KindBad {
			fmt.Printf("metric %q no longer supported\n", nGo)
		}

		mVal := getMetric[0].Value.Uint64()
		fmt.Printf("Number of goroutines: %d\n", mVal)

		// descs := metrics.All()
		// bts, _ := json.MarshalIndent(descs, "", "  ")
		// fmt.Printf("%s\n", bts)
	}

	wg.Wait()
	metrics.Read(getMetric)
	mVal := getMetric[0].Value.Uint64()
	fmt.Printf("Before exiting: %d\n", mVal)
}

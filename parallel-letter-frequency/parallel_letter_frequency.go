package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := make(FreqMap)
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	freqChan := make(chan FreqMap, len(l))
	var wg sync.WaitGroup
	wg.Add(len(l))

	for _, s := range l {
		go func(s string, wg *sync.WaitGroup) {
			freqChan <- Frequency(s)
			wg.Done()
		}(s, &wg)
	}

	wg.Wait()
	close(freqChan)

	res := make(FreqMap)
	for m := range freqChan {
		for k, v := range m {
			res[k] += v
		}
	}
	return res
}

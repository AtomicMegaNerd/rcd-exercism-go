package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// LockedFreqMap adds a mutex to the map.
type LockedFreqMap struct {
	sync.RWMutex
	internal FreqMap
}

func NewLockedFreqMap() *LockedFreqMap {
	return &LockedFreqMap{internal: make(map[rune]int)}
}

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
	var wg sync.WaitGroup
	lm := NewLockedFreqMap()
	for _, segment := range l {
		wg.Add(1)
		go doCalc(segment, &wg, lm)
	}
	wg.Wait()
	return *&lm.internal
}

func doCalc(segment string, wg *sync.WaitGroup, m *LockedFreqMap) {
	m.Lock()
	defer m.Unlock()
	defer wg.Done()
	for _, r := range segment {
		m.internal[r]++
	}
}

package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	ch := make(chan FreqMap)
	result := FreqMap{}
	for _, text := range texts {
		go func(text string) {
			ch <- Frequency(text)
		}(text)
	}
	for range texts {
		for letter, count := range <-ch {
			result[letter] += count
		}
	}
	return result
}
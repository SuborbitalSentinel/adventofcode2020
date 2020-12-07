package types

// StringSet is a valueless map
type StringSet map[string]int

// Add appends a new value to the StringSet
func (s StringSet) Add(input string) {
	s[input]++
}

// Has returns true if the StringSet contains a value
func (s StringSet) Has(input string) bool {
	_, ok := s[input]
	return ok
}

// SeenExactly returns the keys that have been inserted the exact amount of times
func (s StringSet) SeenExactly(input int) (keys []string) {
	for key, value := range s {
		if value == input {
			keys = append(keys, key)
		}
	}
	return
}

package types

// StringSet is a valueless map
type StringSet map[string]struct{}

// Add appends a new value to the StringSet
func (s StringSet) Add(input string) {
	var a struct{}
	s[input] = a
}

// Has returns true if the StringSet contains a value
func (s StringSet) Has(input string) bool {
	_, ok := s[input]
	return ok
}

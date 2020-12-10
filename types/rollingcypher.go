package types

// RollingCypher is a type to aid in attacking xmas encryption
type RollingCypher struct {
	Cypher []int
}

// NewRollingCypher initilizes a new RollingCypher with an initial preamble
func NewRollingCypher(preamble []int) *RollingCypher {
	return &RollingCypher{Cypher: preamble}
}

// Increment pops the oldest value from the cypher and appends the newest to the end
func (c *RollingCypher) Increment(value int) {
	c.Cypher = append(c.Cypher[1:], value)
}

// Validate returns true if it passes the cypher
func (c *RollingCypher) Validate(value int) bool {
	for i := 0; i < len(c.Cypher); i++ {
		for j := i + 1; j < len(c.Cypher); j++ {
			if value == (c.Cypher[i] + c.Cypher[j]) {
				return true
			}
		}
	}
	return false
}

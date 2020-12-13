package types

// RolloverCounter will stay between 0 and 360
type RolloverCounter struct {
	value int
}

// Value returns the counter's value
func (rc *RolloverCounter) Value() int {
	return rc.value
}

// Change modifies the counter's value
func (rc *RolloverCounter) Change(val int) {
	rc.value += val
	if rc.value > 0 {
		rc.value = rc.value % 360
	} else {
		rc.value = (360 + (rc.value % 360)) % 360
	}
}

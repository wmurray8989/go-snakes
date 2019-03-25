package tournament

// Update runs the tournament
func (t *Tournament) Update() {

	// Update active match
	t.activeMatch.Update()
}

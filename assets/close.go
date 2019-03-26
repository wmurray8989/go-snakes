package assets

// Close closes all the assets
func (a *Assets) Close() {
	a.Font30.Close()
	a.Font50.Close()
}

package apiInteraction

type Location struct {
	count    int
	next     string
	previous *string
	results  []struct {
		name string
		url  string
	}
}

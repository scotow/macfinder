package macfinder

type Specs struct {
	Capacity string `json:"dimensionCapacity"`
	Name string `json:"refurbClearModel"`
	Year string `json:"dimensionRelYear"`
	Color string `json:"dimensionColor"`
	Ram string `json:"tsMemorySize"`
	Screen string `json:"dimensionScreensize"`
}

func (s Specs) equal(to Specs) bool {
	return s.Capacity == to.Capacity &&
		s.Name == to.Name &&
		s.Year == to.Year &&
		s.Color == to.Color &&
		s.Ram == to.Ram &&
		s.Screen == to.Screen
}
package macfinder

type Specs struct {
	Capacity string `json:"dimensionCapacity"`
	Name     string `json:"refurbClearModel"`
	Year     string `json:"dimensionRelYear"`
	Color    string `json:"dimensionColor"`
	Ram      string `json:"tsMemorySize"`
	Screen   string `json:"dimensionScreensize"`
}

func (s1 Specs) match(s2 Specs) bool {
	if s2.Capacity != "" && s1.Capacity != s2.Capacity {
		return false
	}
	if s2.Name != "" && s1.Name != s2.Name {
		return false
	}
	if s2.Year != "" && s1.Year != s2.Year {
		return false
	}
	if s2.Color != "" && s1.Color != s2.Color {
		return false
	}
	if s2.Ram != "" && s1.Ram != s2.Ram {
		return false
	}
	if s2.Screen != "" && s1.Screen != s2.Screen {
		return false
	}
	return true
}

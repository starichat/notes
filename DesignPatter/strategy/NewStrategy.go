package strategy


func NewStrategy(t string) res Stragtegier {
	switch t {
	case "m":
		res = Multiplication{}
	case "d":
		res = Division{}
	case "a":
		fallthrough
	default:
		res = Addition{}
	}
	return
}

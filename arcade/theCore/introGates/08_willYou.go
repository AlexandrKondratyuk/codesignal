package introGates

func main() {

}

func willYou(young bool, beautiful bool, loved bool) bool {
	switch {
	case young && beautiful && !loved:
		return true
	case young && !beautiful && loved:
		return true
	case !young && beautiful && loved:
		return true
	case !young && !beautiful && loved:
		return true
	default:
		return false
	}
	return false
}

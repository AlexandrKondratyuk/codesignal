package atTheCrossroads

func reachNextLevel(experience int, threshold int, reward int) bool {
	return threshold <= experience + reward
}

package piscine

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)

	for i := 0; i < n/2; i++ {
		podium[i], podium[n-1-i] = podium[n-1-i], podium[i]
	}

	return podium
}

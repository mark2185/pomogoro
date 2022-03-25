package util

func SecondsToMinutes(seconds int) (int, int) {
	return seconds / 60, seconds % 60
}

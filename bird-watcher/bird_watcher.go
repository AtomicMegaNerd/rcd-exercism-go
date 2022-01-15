package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
	for _, birds := range birdsPerDay {
		count += birds
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startDay := week*7 - 7
	birdsInGivenWeek := birdsPerDay[startDay : startDay+7]
	return TotalBirdCount(birdsInGivenWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for day := range birdsPerDay {
		if day%2 == 0 {
			birdsPerDay[day] += 1
		}
	}
	return birdsPerDay
}

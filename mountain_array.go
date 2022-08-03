package roman_number

func peakIndexInMountainArray(arr []int) int {
	peakIndex := 0
	peakHeight := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > peakHeight {
			peakHeight = arr[i]
			peakIndex = i
		}
	}
	return peakIndex
}

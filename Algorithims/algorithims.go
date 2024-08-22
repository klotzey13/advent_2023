package main

func main() {

}

func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i // Return the index where the target is found
		}
	}
	return -1 // Return -1 if the target is not found
}

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2 // Calculate the middle index

		if arr[mid] == target {
			return mid // Target found at the middle
		} else if arr[mid] < target {
			low = mid + 1 // Search in the right half
		} else {
			high = mid - 1 // Search in the left half
		}
	}

	return -1 // Target not found
}

func BubbleSort(arr []int) {
	n := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i] // Swap elements
				swapped = true
			}
		}
		n-- // Last n elements are already in place
	}
}

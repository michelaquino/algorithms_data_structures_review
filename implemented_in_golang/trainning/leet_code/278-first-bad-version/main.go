package main

import "fmt"

// https://leetcode.com/problems/first-bad-version/

var (
	// n          = 2126753390
	// badVersion = 1702766719
	// n          = 5
	// badVersion = 4
	// n          = 1
	// badVersion = 1
	n          = 3
	badVersion = 1
)

func main() {

	fmt.Println(firstBadVersionLoop(n))
	fmt.Println(firstBadVersion(n))
}

/*
Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.
*/

func firstBadVersionLoop(n int) int {
	firstBadVersion := 0
	for i := 1; i <= n; i++ {
		if isBadVersion(i) {
			return i
		}
	}

	return firstBadVersion
}

func firstBadVersion(n int) int {
	return bynarySearch(0, n)
}

func bynarySearch(startIndex, endIndex int) int {
	for endIndex-startIndex > 1 {

		// middle := startIndex + (endIndex-startIndex)/2
		middle := (endIndex + startIndex) / 2

		/** middle = (start+end)) / 2; **/
		if isBadVersion(middle) {
			endIndex = middle
		} else {
			startIndex = middle
		}
	}

	return endIndex
}

func bynarySearchRecursive(startIndex, endIndex int) int {
	if startIndex > endIndex {
		return -1
	}

	middle := (endIndex + startIndex) / 2
	if isBadVersion(middle) && !isBadVersion(middle-1) {
		return middle
	}

	badIndex := bynarySearch(middle+1, endIndex)
	if badIndex != -1 {
		return badIndex
	}

	return bynarySearch(1, middle-1)
}

func isBadVersion(version int) bool {
	if version >= badVersion {
		return true
	}

	return false
}

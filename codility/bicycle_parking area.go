// You want to park your bicycle in a bicycle parking area where bike racks are aligned in a row. There are already N bikes parked there (each bike is attached to exactly one rack, but a rack can have multiple bikes attached to it). We call racks that already have bikes attached used.
// You want to park your bike in a rack in the parking area according to the following criteria:
// • the chosen rack must lie between the first and the last used racks (inclusive);
// • the distance between the chosen rack and any other used rack is as big as possible.
// A description of the bikes already parked in the racks is given in a non-empty zero-indexed array A: element AlK denotes the position of the rack to which bike number K is attached (for 0 ≤ K < N).
// The central position in the parking area is position 0. A positive value means that the rack is located AlK meters to the right of the central position 0, a negative value means that it's located IA|K|| meters to the left (the absolute value of A[K))
// For example, consider array A such that:
// A101 = 10
// A111 = 0
// A 21 - 8
// A31 = 2
// AI41 =-1
// A151 = 12
// A [6] = 11
// A[71 = 3
// In the figure below, available racks are represented by dots. Bigger dots represent racks with bikes attached to them.

// In the figure below, available racks are represented by dots. Bigger dots represent racks with bikes attached to them.
// 1823:
// 8 10 11 12
// You can attach your bike to any rack between rack - 1 and rack 12 (including these two racks). In order to maximize the distance to any used rack, you should attach your bike either to rack 5 or to rack 6. The resulting distance is 2 meters (from 5 to used rack 3, or from 6 to used rack 8, respectively).
// Write a function:
// fune Solution(A [lint) int
// that, given a non-empty zero-indexed array A of N integers, returns the largest possible distance in meters between the chosen rack and any other used rack.
// Given array A shown above, the function should return 2, as explained above.
// For the following array A:
// A|01 = 5
// AL1I=5
// the function should return 0, as you can attach your bike only to rack 5.
// Write an efficient algorithm for the following assumptions:
// • N is an integer within the range 2.100,000];
// • each element of array A is an integer within the range
// I-1,000,000,000.1,000,000,000.

package solution

// you can also use imports, for example:
import (
	"sort"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func SolutionB(A []int) int {
	sort.Ints(A)
	maxDistance := 0
	for i := 1; i < len(A); i++ {
		distance := A[i] - A[i-1]
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return maxDistance / 2
}

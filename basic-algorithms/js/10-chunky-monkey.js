function chunkArrayInGroups(arr, size, bnoden, temp) {
	bnoden = [], temp = []
	for (var i in arr) {
		temp.push(arr[i])
		if (i % size === size-1)
			bnoden.push(temp), temp = []
	}
	if (temp.length !== 0) bnoden.push(temp)

	// Only for testing
	function test() {
		for (var i in bnoden)
			console.log(bnoden[i])
	}
	return test(bnoden) // Only for testing

	// return bnoden
}

// Test
console.log(chunkArrayInGroups(["a", "b", "c", "d"], 2)) // [["a", "b"], ["c", "d"]]
console.log(chunkArrayInGroups([0, 1, 2, 3, 4, 5], 3)) // [[0, 1, 2], [3, 4, 5]]
console.log(chunkArrayInGroups([0, 1, 2, 3, 4, 5], 2)) // [[0, 1], [2, 3], [4, 5]]
console.log(chunkArrayInGroups([0, 1, 2, 3, 4, 5], 4)) // [[0, 1, 2, 3], [4, 5]]
console.log(chunkArrayInGroups([0, 1, 2, 3, 4, 5, 6], 3)) // [[0, 1, 2], [3, 4, 5], [6]]
console.log(chunkArrayInGroups([0, 1, 2, 3, 4, 5, 6, 7, 8], 4)) // [[0, 1, 2, 3], [4, 5, 6, 7], [8]]
console.log(chunkArrayInGroups([0, 1, 2, 3, 4, 5, 6, 7, 8], 2)) // [[0, 1], [2, 3], [4, 5], [6, 7], [8]]
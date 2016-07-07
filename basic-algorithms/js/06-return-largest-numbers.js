function largestOfFour(arr, bnoden) {
	bnoden = []
	for (var i in arr) bnoden[i] = 0
	for (var i in arr)
		for (var j in arr[i])
			if (arr[i][j] > bnoden[i]) bnoden[i] = arr[i][j]
	return bnoden
}

console.log(largestOfFour([[13, 27, 18, 26], [4, 5, 1, 3], [32, 35, 37, 39], [1000, 1001, 857, 1]])) // [27 5 39 1001]
console.log(largestOfFour([[4, 9, 1, 3], [13, 35, 18, 26], [32, 35, 97, 39], [1000000, 1001, 857, 1]])) // [9 35 97 1000000]
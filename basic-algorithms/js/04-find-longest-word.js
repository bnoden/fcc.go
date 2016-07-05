function findLongestWord(str, bnoden, arr) {
	bnoden = 1, arr = str.split(' ')
	for (var i in arr) if (arr[i].length > bnoden) bnoden = arr[i].length
	return bnoden
}

findLongestWord("The quick brown fox jumped over the lazy dog");
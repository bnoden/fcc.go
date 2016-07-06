function findLongestWord(str, bnoden, arr) {
	bnoden = 1, arr = str.split(' ')
	for (var i in arr) if (arr[i].length > bnoden) bnoden = arr[i].length
	return bnoden
}
// Test
findLongestWord("The quick brown fox jumped over the lazy dog"); // 6
findLongestWord("May the force be with you") // 5.
findLongestWord("Google do a barrel roll") // 6.
findLongestWord("What is the average airspeed velocity of an unladen swallow") // 8.
findLongestWord("What if we try a super-long word such as otorhinolaryngology") // 19
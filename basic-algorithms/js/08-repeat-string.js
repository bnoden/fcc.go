function repeatStringNumTimes(str, num, bnoden) {
	bnoden = ''
	if (!isNaN(num) && num > 0) bnoden = str.repeat(num)
	return bnoden
}

console.log(repeatStringNumTimes("*", 3)) // ***
console.log(repeatStringNumTimes("abc", 3)) // abcabcabc
console.log(repeatStringNumTimes("abc", 4)) // abcabcabcabc
console.log(repeatStringNumTimes("abc", 1)) // abc
console.log(repeatStringNumTimes("*", 8)) // ********
console.log(repeatStringNumTimes("abc", -2)) // ""
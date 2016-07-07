function truncateString(str, num, bnoden) {
	bnoden = str
	if (str.length > num) {
		if (num > 3) { num-=3 }
		bnoden = str.slice(0, num)+'...'
	} return bnoden
}

console.log(truncateString("A-tisket a-tasket A green and yellow basket", 11)) // "A-tisket..."
console.log(truncateString("Peter Piper picked a peck of pickled peppers", 14)) // "Peter Piper..."
console.log(truncateString("A-tisket a-tasket A green and yellow basket", "A-tisket a-tasket A green and yellow basket".length)) // "A-tisket a-tasket A green and yellow basket"
console.log(truncateString("A-tisket a-tasket A green and yellow basket", "A-tisket a-tasket A green and yellow basket".length + 2)) // "A-tisket a-tasket A green and yellow basket"
console.log(truncateString("A-", 1)) // "A..."
console.log(truncateString("Absolutely Longer", 2)) // "Ab..."
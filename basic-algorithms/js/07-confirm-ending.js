function confirmEnding(str, target, bnoden) {
	bnoden = false
	if (str.substr((str.length - target.length), str.length-1) === target)
		bnoden = true
	return bnoden
}

console.log(confirmEnding("Bastian", "n")) // true
console.log(confirmEnding("Connor", "n")) // false
console.log(confirmEnding("Walking on water and developing software from a specification are easy if both are frozen", "specification")) // false
console.log(confirmEnding("He has to give me a new name", "name")) // true
console.log(confirmEnding("He has to give me a new name", "me")) // true
console.log(confirmEnding("He has to give me a new name", "na")) // false
console.log(confirmEnding("If you want to save our world, you must hurry. We dont know how much longer we can withstand the nothing", "mountain")) // false
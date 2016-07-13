function slasher(arr, bnoden) {
	return arr.splice(bnoden, arr.length)
}

console.log(slasher([1, 2, 3], 2)) // [3]
console.log(slasher([1, 2, 3], 0))// [1, 2, 3]
console.log(slasher([1, 2, 3], 9)) // []
console.log(slasher([1, 2, 3], 4)) // []
console.log(slasher(["burgers", "fries", "shake"], 1)) // ["fries", "shake"]
console.log(slasher([1, 2, "chicken", 3, "potatoes", "cheese", 4], 5)) // ["cheese", 4]
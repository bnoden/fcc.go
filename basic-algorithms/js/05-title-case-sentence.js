function titleCase(str, arr, bnoden) {
	arr = str.split(' '), bnoden = []
	for (var i in arr)
		bnoden[i] = arr[i].charAt(0).toUpperCase()+arr[i].substr(1).toLowerCase()
	return bnoden.join(' ')
}
// Test
titleCase("I'm a little tea pot") // I'm A Little Tea Pot
titleCase("sHoRt AnD sToUt") // Short And Stout
titleCase("HERE IS MY HANDLE HERE IS MY SPOUT") // Here Is My Handle Here Is My Spout
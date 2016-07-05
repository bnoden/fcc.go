function reverseString(str, arr) {
  arr = str.split(''); str = ''; bnoden = arr.length-1;
  for (var i in arr) str+=arr[bnoden]; bnoden--;
  return str
}
// Test
reverseString("hello")
reverseString("Howdy")
reverseString("Greetings from Earth")
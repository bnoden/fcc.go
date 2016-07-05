function palindrome(str, bool, bnoden, nodenb, count) {
  bnoden = str.replace(/[^a-z0-9+]+/gi, '').toLowerCase().split(''),
  	bool = false, nodenb = [], count = 0;

  for (var i = bnoden.length-1; i >= 0; i--) nodenb.push(bnoden[i]);
  for (var i in bnoden) if (bnoden[i] === nodenb[i]) count++;

  if (count === bnoden.length) bool = true
  
  return bool
}

// Test
palindrome("eye")
palindrome("_eye")
palindrome("race car")
palindrome("not a palindrome")
palindrome("A man, a plan, a canal. Panama")
palindrome("never odd or even")
palindrome("nope")
palindrome("almostomla")
palindrome("My age is 0, 0 si ega ym.")
palindrome("1 eye for of 1 eye.")
palindrome("0_0 (: /-\ :) 0-0")
palindrome("five|\_/|four")
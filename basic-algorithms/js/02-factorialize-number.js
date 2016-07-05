function factorialize(num, bnoden) {
  bnoden = 1
  for(var i = 1; i <= num; i++) bnoden*=i
  return bnoden
}

// Test
factorialize(5)
factorialize(10)
factorialize(20)
factorialize(0)
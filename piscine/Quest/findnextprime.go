package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	// Мы просто вызываем IsPrime, она уже объявлена в isprime.go
	for i := nb; ; i++ {
		if IsPrime(i) {
			return i
		}
	}
}

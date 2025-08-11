package main

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	eqCharS := make(map[byte]byte)
	eqCharT := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if _, ok := eqCharS[s[i]]; !ok {
			if _, ok := eqCharT[t[i]]; ok {
				return false
			}
			eqCharS[s[i]] = t[i]
			eqCharT[t[i]] = s[i]
		} else if eqCharS[s[i]] != t[i] {
			return false
		}
	}

	return true
}

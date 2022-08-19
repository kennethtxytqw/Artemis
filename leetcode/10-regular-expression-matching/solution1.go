package main

func solution1(s string, p string) bool {
	return match(s, p)
}

func match(s string, p string) bool {

	memo := make([][]bool, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		memo[i] = make([]bool, len(s)+1)
	}

	memo[0][0] = true

	for I := 1; I <= len(p); I++ {
		i := I - 1
		if p[i] == '*' {
			memo[I][0] = memo[I-2][0]
		}
		for J := 1; J <= len(s); J++ {
			j := J - 1
			if p[i] == s[j] || p[i] == '.' {
				memo[I][J] = memo[I-1][J-1]
			} else if p[i] == '*' {
				if memo[I-2][J] {
					memo[I][J] = true
				} else if p[i-1] == '.' || p[i-1] == s[j] {
					memo[I][J] = memo[I-1][J] || memo[I][J-1]
				}
			}
		}
	}
	return memo[len(p)][len(s)]
}

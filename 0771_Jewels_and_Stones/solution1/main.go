package solution1

func numJewelsInStones(J string, S string) int {
	/*
	   暴力
	*/
	ret := 0
	for i := 0; i < len(S); i ++ {
		for j := 0; j < len(J); j++ {
			if S[i] == J[j] {
				ret += 1; break
			}
		}
	}

	return ret
}
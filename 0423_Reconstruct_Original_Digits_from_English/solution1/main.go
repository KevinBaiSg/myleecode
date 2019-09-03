package solution1

func originalDigits(s string) string {
	counter := make([]int, 26)
	for _, c := range s {
		counter[c-'a']++
	}

	number := make([]int, 10)
	number[0] = counter['z'-'a']
	number[2] = counter['w'-'a']
	number[4] = counter['u'-'a']
	number[6] = counter['x'-'a']
	number[8] = counter['g'-'a']

	number[3] = counter['h'-'a'] - number[8]
	number[5] = counter['f'-'a'] - number[4]
	number[7] = counter['s'-'a'] - number[6]

	number[1] = counter['o'-'a'] - number[0] - number[2] - number[4]
	number[9] = counter['i'-'a'] - number[5] - number[6] - number[8]

	ret := make([]byte, 0)
	for i := byte(0); i < 10; i++ {
		for j := 0; j < number[i]; j++ {
			ret = append(ret, i + '0')
		}
	}

	return string(ret)
}

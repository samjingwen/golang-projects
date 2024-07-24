package qns1_3

func URLify(url *[]rune, trueLength int) {
	count := 0
	for i := 0; i < len(*url); i++ {
		if (*url)[i] == ' ' {
			count++
		}
	}

	newLength := trueLength + count*2
	*url = (*url)[:newLength]
	index := newLength - 1

	for i := trueLength - 1; i >= 0; i-- {
		if (*url)[i] == ' ' {
			(*url)[index] = '0'
			(*url)[index-1] = '2'
			(*url)[index-2] = '%'
			index -= 3
		} else {
			(*url)[index] = (*url)[i]
			index--
		}
	}
}

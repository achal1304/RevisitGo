package main

func getMaximumApiSecurity(currentKey string, rotatedKey string) string {
	// Write your code here
	countRot1 := 0
	countRot0 := 0

	for i := 0; i < len(currentKey); i++ {
		if rotatedKey[i] == '1' {
			countRot1++
		} else {
			countRot0++
		}
	}

	res := make([]byte, len(currentKey))

	for i := 0; i < len(currentKey); i++ {
		if currentKey[i] == '0' {
			if countRot1 > 0 {
				res[i] = '1'
				countRot1--
			} else {
				res[i] = '0'
				countRot1--
			}
		} else {
			if countRot0 > 0 {
				res[i] = '1'
				countRot0--
			} else {
				res[i] = '0'
				countRot1--
			}
		}
	}

	return string(res)
}

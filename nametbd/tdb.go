package nametbd

func idkfunc(sl []int) int {
	if len(sl) == 0 {
		return -1
	}
	var (
		counter int
		major   int
	)
	for _, v := range sl {
		if counter == 0 {
			major = v
			counter++
			continue
		}
		if v == major {
			counter++
		} else {
			counter--
		}
	}

	var checkingC int
	for _, v := range sl {
		if v == major {
			checkingC++
		}
	}
	if checkingC <= len(sl)/2 {
		return -1
	}

	return major
}

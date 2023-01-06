package main

func Longest_substring(reqs []int) []int {
	l, r, finL, finR := 0, 1, 0, 1
	flagg := true
	for r < len(reqs) {
		if reqs[r] == reqs[r-1] || (r > 2 && reqs[r] == reqs[r-2]) {
			flagg = true
		}
		if reqs[r] != reqs[r-1] {
			if finR-finL < r-l {
				finR, finL = r, l
			}
			if flagg {
				flagg = false
			} else {
				l = r
			}
		}

		r++
	}
	if finR-finL < r-l {
		finR, finL = r, l
	}

	return reqs[finL:finR]
}

/*	Driver code
	var req, req_num, num int

	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &num)
	for i := 0; i < num; i++ {
		fmt.Fscan(in, &req_num)
		slicc := make([]int, req_num)
		for j := 0; j < req_num; j++ {
			fmt.Fscan(in, &req)
			slicc[j] = req
		}
		// fmt.Println(slicc)
		fmt.Println(len(longest_substring(slicc)))
	}
*/

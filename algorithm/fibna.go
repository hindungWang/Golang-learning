package main

func fib(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
func fibPro(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	res := 1
	pre := 1
	tmp := 0

	for i := 3; i <= n; i++ {
		tmp = res
		res = res + pre
		pre = tmp
	}
	return res
}

func metrixPower(m [2][2]int, p int) [2][2]int {
	res := [2][2]int{}
	for i := 0; i < 2; i++ {
		res[i][i] = 1
	}
	tmp := m
	for ; p != 0; p >>= 1 {
		if (p & 1) != 0 {
			res = mutiMatrix(res, tmp)
		}
		tmp = mutiMatrix(tmp, tmp)
	}
	return res
}
func mutiMatrix(m1, m2 [2][2]int) [2][2]int {
	res := [2][2]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}
func main() {

	// fmt.Println(fibPro(80))
	//fmt.Println(fib(80))

}

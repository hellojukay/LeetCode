package main

func main()

func fx(n int) int {
	// 台阶数是1，只有一种上法
	if n == 1 {
		return
	}
	// 台阶数是2，有两种方法，
	if n == 2 {
		return 2
	}
	// 台阶数大于2，有两种方法，1：消耗一个台阶剩余n-1个台阶，
	//1：消耗一个台阶剩余n-1个台阶
	//1：消耗一个台阶剩余n-2个台阶
	return fx(n-1) + fx(n-2)
}

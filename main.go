package main

import (
	"fmt"
	"math"
)

// TODO: 请用高效算法实现一个计算斐波那契数列第 n 项的函数
// 提示：可以考虑递归、动态规划或矩阵快速幂
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	
	return 0
}

// TODO: 实现一个函数，判断一个整数是否为质数
// 额外要求：使用高效算法（如试除法优化到 sqrt(n)）
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	
	return false // 故意留空
}

// TODO: 实现一个函数，计算圆的面积（接收半径，返回面积）
// 使用 math.Pi，并保留 4 位小数输出
func circleArea(radius float64) float64 {
	return 0.0
}

func main() {
	fmt.Println("Hello from GitHub Copilot demo!")

	fmt.Println("Fibonacci(10):", fibonacci(10))

	fmt.Println("Is 17 prime?", isPrime(17))

	fmt.Println("Area of circle with radius 5:", circleArea(5))
}

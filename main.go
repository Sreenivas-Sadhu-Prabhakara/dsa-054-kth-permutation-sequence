package main

import "fmt"

// KthPermutation returns the k-th (1-based) permutation of the digits 1..n
// in lexicographic order, without generating the earlier ones: each output
// position is chosen by dividing k by the factorial of the remaining slots.
// n is at most 12 (13! overflows the int math long before the digit set runs out).
func KthPermutation(n, k int) string {
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i
	}
	if n < 1 || n > 12 || k < 1 || k > fact[n] {
		return ""
	}
	k-- // 0-based rank
	digits := []byte{}
	for i := 1; i <= n; i++ {
		digits = append(digits, byte('0'+i))
	}
	out := make([]byte, 0, n)
	for i := n; i >= 1; i-- {
		idx := k / fact[i-1]
		k %= fact[i-1]
		out = append(out, digits[idx])
		digits = append(digits[:idx], digits[idx+1:]...)
	}
	return string(out)
}

func main() {
	fmt.Println(KthPermutation(4, 9))
}

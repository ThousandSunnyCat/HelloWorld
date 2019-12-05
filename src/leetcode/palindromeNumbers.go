package leetcode

// 9
// 注释为学习解法
func PalindromeNumber(x int) bool {
	if x < 0 || x % 10 == 0 {
        return false
	}
	
	// p,q:=0,x
	// for p < q {
	// 	p, q = p * 10 + q % 10, q / 10
	// }

	// return p == q || q == p / 10

    k := make([]int, 0)
    for ; x >= 1; x /= 10 {
        f := x % 10
        k = append(k, f)
    }

    p, q := 0, len(k)-1
    for ; p < q; {
        if k[p] != k[q] {
            return false
        }
        p++
        q--
    }

    return true
}
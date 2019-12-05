package leetcode

func IsValid(s string) bool {
	var lk int
	myStact := make([]int32, 0)
	for _, mys := range s {
		
		lk = len(myStact) - 1
		if mys == '(' || mys == '[' || mys == '{' {
			myStact = append(myStact, mys)
			continue
		}
		if lk < 0 {
			return false
		}
		if mys == ')' && myStact[lk] != '(' {
			return false
		}
		if mys == ']' && myStact[lk] != '[' {
			return false
		}
		if mys == '}' && myStact[lk] != '{' {
			return false
		}

		myStact = myStact[:lk]
	}
	
	return len(myStact) == 0
}
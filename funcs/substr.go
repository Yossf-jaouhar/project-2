package mosdef

func SubStr(str, subString string) []int {
	sl:=[]int{}
	for x := range str {
		if x <= len(str)-len(subString) && str[x:x+len(subString)] == subString {
			sl=append(sl, x)
		}
	}
	return sl
}

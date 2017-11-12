package gotrie

func min(l1, l2 int) int {
	if l1 > l2 {
		return l2
	}
	return l1
}

func lcs(s1, s2 string) (result, r1, r2 string) {
	var common int
	length := min(len(s1), len(s2))
	for i := 0; i < length; i, common = i+1, common+1 {
		if s1[i] != s2[i] {
			break
		}
	}
	return s1[:common], s1[common:], s2[common:]
}

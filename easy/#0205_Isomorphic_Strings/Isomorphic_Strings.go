package _0205_Isomorphic_Strings

// 可以把字符串转成数字字符串表示
func isIsomorphic(s string, t string) bool {
	smap := make(map[byte]int)
	tmap := make(map[byte]int)
	ssum, tsum := 0, 0
	for i, _ := range []byte(s) {
		if n, ok := smap[s[i]]; ok {
			ssum = n
		} else {
			ssum = i
			smap[s[i]] = i
		}
		if n, ok := tmap[t[i]]; ok {
			tsum = n
		} else {
			tsum = i
			tmap[t[i]] = i
		}
		if ssum != tsum {
			return false
		}
	}
	return true
}

// 官方解：s 和 t 相互映射
func isIsomorphic2(s string, t string) bool {
	smap := make(map[byte]byte)
	tmap := make(map[byte]byte)
	for i, _ := range s {
		x, y := s[i], t[i]
		if smap[x] > 0 && smap[x] != y || tmap[y] > 0 && tmap[y] != x {
			return false
		}
		smap[x] = y
		tmap[y] = x
	}
	return true
}

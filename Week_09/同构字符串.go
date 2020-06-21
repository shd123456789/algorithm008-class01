func isIsomorphic(s string, t string) bool {
	hash1, hash2 := make([]int, 128), make([]int, 128)
	for i := 0; i < len(s); i++ {
		if hash1[s[i]] != hash2[t[i]] {
			return false
		}
		hash1[s[i]], hash2[t[i]] = i+1, i+1
	}
	return true
}
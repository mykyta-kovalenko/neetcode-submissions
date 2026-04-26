func isAnagram(s string, t string) bool {
if len(s) != len(t) {
    return false
}

counts := make(map[byte]int)

for i := 0; i < len(s); i++ {
    chS := s[i]
    chT := t[i]

    counts[chS]++
    counts[chT]--
}

for _, count := range counts {
    if count != 0 {
        return false
    }
}

return true
}

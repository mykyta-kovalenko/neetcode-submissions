func groupAnagrams(strs []string) [][]string {
    anagramsMap := make(map[[26]int][]string)

    for i := 0; i < len(strs); i++ {
        word := strs[i]
        var signature [26]int
        for _, char := range word {
            signature[char - 'a']++
        }
        anagramsMap[signature] = append(anagramsMap[signature], word)
    }

    result := make([][]string, 0, len(anagramsMap))
    
    for _, value := range anagramsMap {
        result = append(result, value)
    }

    return result
}

// complexity:
//   O(n*k) time and space


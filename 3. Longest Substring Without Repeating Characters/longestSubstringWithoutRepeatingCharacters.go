func lengthOfLongestSubstring(s string) int {
    lastOccurred := make(map[rune] int)
    str := []rune(s)
    var count, max, start int
    
    for i, v := range str {
        idx, ok := lastOccurred[v]
        lastOccurred[v] = i
        if ok == false {
            count++
        } else {
            if idx < start {
                count++
            } else {
                start = idx + 1
                if count > max {
                    max = count
                }
                count = i - start + 1
            }
        }
    }
    
    if count > max {
        return count
    }
    return max
}
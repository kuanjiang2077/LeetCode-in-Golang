func longestPalindrome(s string) string {
    bs := []byte(s)
    max := 0
    str := ""
    for i, _ := range bs {
        s = palindrome(bs, i, i)
        if len(s) > max {
            max, str = len(s), s
        }
        s = palindrome(bs, i, i+1)
        if len(s) > max {
            max, str = len(s), s
        }
    } 
    
    return str
}

func palindrome(b []byte, l int, r int) string {
    for l >=0 && r < len(b) && b[l] == b[r] {
        l--
        r++
    }
    return string(b[l+1:r])
}
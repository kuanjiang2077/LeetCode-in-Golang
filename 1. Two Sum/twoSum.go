func twoSum(nums []int, target int) []int {
    m := make(map[int] int)
    for i, n := range nums {
        val, ok := m[n]
        if ok {
            return []int {val, i}
        } else {
            m[target - n] = i
        }
    }
    return []int {-1, -1}
}
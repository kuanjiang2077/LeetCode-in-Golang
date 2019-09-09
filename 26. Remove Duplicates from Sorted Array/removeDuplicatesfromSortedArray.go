func removeDuplicates(nums []int) int {
  slow, old, fast := 0,0,0
  for fast < len(nums) {
    if nums[fast] != nums[old] {
      slow++
      nums[slow] = nums[fast]
      old = fast
    }
    fast++
  }
  return slow+1
}

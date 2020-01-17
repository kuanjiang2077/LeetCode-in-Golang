func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // x|x 
    // x x x|x x
    // i, j for split position [0, len]
    // total = (m + n + 1) / 2
    // odd return maxOfLeft

    m, n := len(nums1), len(nums2)
    if m > n {
        nums1, nums2, m, n = nums2, nums1, n, m
    }

    iMin, iMax, total := 0, m, (m+n+1)/2 
    for {
        i := iMin + (iMax-iMin) / 2
        j := total - i
        switch
        {
            case i > 0 && nums1[i-1] > nums2[j]:
            {
                iMax = i - 1
            }
            case i < m && nums1[i] < nums2[j-1]:
            {
                iMin = i + 1
            }
            default:
            {
                maxOfLeft, minOfRight:= 0, 0
                
                if i == 0 {
                    maxOfLeft = nums2[j-1]
                } else if j == 0 {
                    maxOfLeft = nums1[i-1]
                } else {
                    maxOfLeft = max(nums1[i-1], nums2[j-1])
                }
                if (m+n)%2 == 1 {
                    return float64(maxOfLeft)
                }

                if i == m {
                    minOfRight = nums2[j]
                } else if j == n {
                    minOfRight = nums1[i]
                } else {
                    minOfRight = min(nums1[i], nums2[j])
                }
        
                return float64(maxOfLeft + minOfRight) / 2.0  
            }
        }
    }

    return -1.0
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:
// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

func twoSum(nums []int, target int) []int {
    // 时间 O(n) 空间需要新开辟内存
    res := make([]int, 2)
    m := make(map[int]int, len(nums))
    for i:=0;i<len(nums);i++{
        idx, exists := m[nums[i]]
        if !exists {
            m[target - nums[i]] = i
        } else {
            res[0] = idx
            res[1] = i
            return res
        }
    }
    return res
}
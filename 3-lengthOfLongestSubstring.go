// Given a string, find the length of the longest substring without repeating characters.

// Examples:

// Given "abcabcbb", the answer is "abc", which the length is 3.

// Given "bbbbb", the answer is "b", with the length of 1.

// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
    max:=0 // 待返回的最大值
    current:=0 // 每次迭代的最大值
    m:=make(map[rune]int) // hashtable 索引表
    for i, v := range s {
        idx, exists := m[v] // idx 是重新计算点
        if exists && i - idx < current{
            current = i-idx
        } else { // 如果不存在或是废弃的记录，直接更新current值
            current += 1
            if current > max {
                max = current
            }
        }
        // 更新 v 的索引
        m[v] = i  
    }

    return max 
}
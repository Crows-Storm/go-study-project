package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// from leetcode
func main() {
	anagrams := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}) // medium
	fmt.Println(anagrams)
	groupAnagrams2([]string{"paper", "apple", "title", "banana", "hello", "hills"}) // medium
	groupAnagrams3([]string{"paper", "apple", "title", "banana", "hello", "hills"}) // medium

	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2})) // medium

	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12})) // simple

	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))  // medium
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // medium

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))  // medium
	fmt.Println(threeSum([]int{0, 0, 0}))              // medium
	fmt.Println(threeSum2([]int{-1, 0, 1, 2, -1, -4})) // medium
	fmt.Println(threeSum2([]int{0, 0, 0, 0}))          // medium
	fmt.Println(threeSum2([]int{0, 0, 0, 0}))          // medium
	fmt.Println(threeSum2([]int{-1, 0, 1, 2, -1, -4})) // medium

	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))  // hard
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                    // hard
	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // hard
	fmt.Println(trap2([]int{4, 2, 0, 3, 2, 5}))                   // hard

	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // medium
	fmt.Println(lengthOfLongestSubstring("bbbbbbb"))  // medium
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // medium

	fmt.Println(findAnagrams("cbaebabacd", "abc"))  // medium
	fmt.Println(findAnagrams("abab", "ab"))         // medium
	fmt.Println(findAnagrams2("cbaebabacd", "abc")) // medium
	fmt.Println(findAnagrams2("abab", "ab"))        // medium

	fmt.Println(subarraySum([]int{1, 1, 1}, 2))                   // medium
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))                   // medium
	fmt.Println(subarraySum([]int{3, 1, 2, 3, 2, 1, 3, 3, 3}, 3)) // medium
	fmt.Println(subarraySum([]int{1, 2, 1, 2, 1}, 3))             // medium
	fmt.Println(subarraySum([]int{1, -1, 0}, 0))                  // medium

}

func subarraySum(nums []int, k int) (count int) {
	fmt.Println("========== subarray Sum 2 ==========")
	sum := 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1 // v = 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if val, exists := prefixSum[sum-k]; exists {
			count += val
		}
		prefixSum[sum]++
	}
	return count
}

func findAnagrams(s string, p string) (sub_str []int) {
	fmt.Println("========== find Anagrams ==========")
	if len(s) < len(p) {
		return []int{}
	}

	p_table, win_table := [26]int{}, [26]int{}

	for i := 0; i < len(p); i++ {
		// conversion to ASCII Code
		p_table[p[i]-'a']++
		win_table[s[i]-'a']++
		fmt.Println("p[i] - 'a': ", p[i]-'a', ", s[i] - 'a': ", s[i]-'a')
	}

	if win_table == p_table { // first win
		fmt.Println("First win end")
		sub_str = append(sub_str, 0)
	}

	for i := len(p); i < len(s); i++ {
		win_table[s[i-len(p)]-'a']--
		fmt.Println("s[i-len(p)]-'a': ", s[i-len(p)]-'a', "need - 1")
		win_table[s[i]-'a']++
		fmt.Println("s[i] - 'a': ", s[i]-'a', "need + 1")

		if p_table == win_table {
			sub_str = append(sub_str, i-len(p)+1)
		}
	}
	return
}

func findAnagrams2(s string, p string) (ans []int) {
	cntP := [26]int{}
	for _, c := range p {
		cntP[c-'a']++
	}

	cntS := [26]int{}
	for right, c := range s {
		cntS[c-'a']++
		left := right - len(p) + 1
		if left < 0 {
			continue
		}
		if cntS == cntP {
			ans = append(ans, left)
		}
		cntS[s[left]-'a']--
	}
	return
}

func lengthOfLongestSubstring(strs string) (max_len int) {
	fmt.Println("========== length Of Longest Substring ==========")

	charSet := make(map[byte]bool)
	left := 0

	for right := 0; right < len(strs); right++ {
		for charSet[strs[right]] {
			delete(charSet, strs[left])
			left++
		}

		charSet[strs[right]] = true

		l := right - left + 1
		if l > max_len {
			max_len = l
		}
	}

	return
}

func trap(height []int) (area int) {
	fmt.Println("========== trap ==========")
	l := len(height) - 1
	for i := 0; i < l; i++ {
		left, right := i, l
		if height[i] == 0 {
			fmt.Println("skip: ", height[i], " i: ", i)
			left++
			continue
		}
		maxRight := 0
		maxLeft := 0
		for left < right {
			fmt.Println("left: ", left, " right: ", right)
			if height[right] > height[i] {
				maxRight = height[right]
			} else if height[left] > height[i] {
				maxLeft = height[left]
			}

			left++
			right--
		}
		fmt.Println("end----", i, height[left], area, maxRight, maxLeft, "\n", left, right)
		if maxRight > maxLeft {
			fmt.Println("maxRight > height[left]: ", maxRight, height[left])
			area += maxRight - maxLeft
			left++
		} else if maxRight == maxLeft {
			fmt.Println("maxRight == height[left]: ", maxRight, height[left])
			area += maxRight
			left++
		} else {
			fmt.Println("maxRight < height[left]: ", maxRight, height[left])
			area += maxLeft - maxRight
			left++
		}
		fmt.Println("area----", area)
	}
	return
}

func trap2(height []int) (area int) {
	fmt.Println("========== trap 2 ==========")
	if len(height) < 3 {
		return
	}
	left, right := 0, len(height)-1
	left_max, right_max := 0, 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= left_max {
				left_max = height[left]
			} else {
				area += left_max - height[left]
			}
			left++
		} else {
			if height[right] >= right_max {
				right_max = height[right]
			} else {
				area += right_max - height[right]
			}
			right--
		}
	}
	return
}

// O(n²)
func threeSum(nums []int) (results [][]int) {
	fmt.Println("========== three Sum ==========")
	left, right := 0, 1

	mapping := make(map[int][]int) // sum : num array
	secondIndex := right + 1
	for left < (len(nums) - 2) {

		first := nums[right]
		second := nums[secondIndex]

		result := nums[left] + first + second
		fmt.Println("result:", result)

		if mapping[result] == nil {
			mapping[result] = append(mapping[result], nums[left], first, second)
		} else if result == 0 {
			fmt.Println("add to mapping: ", result, " because: ", nums[left], first, second, " has: ", mapping[result])
			results = append(results, mapping[result])
			delete(mapping, result)
			results = append(results, []int{nums[left], first, second})
		}

		if secondIndex == len(nums)-1 { // reset pointer, move left pointer
			left++
			secondIndex = left + 2
			right = left + 1
		} else { // move right and secondIndex  pointer
			right++
			secondIndex++
		}

		if left == len(nums)-2 && mapping[0] != nil {
			results = append(results, mapping[result])
			delete(mapping, result)
		}
	}
	return
}

// O(n²)
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	results := [][]int{}
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		target := -nums[i]
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return results
}

func maxArea(height []int) (maxArea int) { // O(N²)
	fmt.Println("========== max Area ==========")
	offset := 0
	for i := 0; i < len(height); i++ { // move long
		for i2 := offset + 1; i2 < len(height); i2++ { // nex line
			low := 0 // low
			if height[offset] > height[i2] {
				low = height[i2]
			} else if offset != i2 {
				low = height[offset]
			}

			area := (i2 - offset) * low
			if area > maxArea {
				maxArea = area
			}
		}
		offset++
	}
	return
}

func maxArea2(height []int) (maxArea int) { // O(N)
	fmt.Println("========== max Area 2 ==========")
	left, rigth := 0, len(height)-1

	for left < rigth { // while
		h := height[left]
		if height[left] > height[rigth] {
			h = height[rigth]
		}
		width := rigth - left
		area := width * h
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[rigth] { // offset
			left++
		} else {
			rigth--
		}
	}
	return
}

/*
from https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked
Input: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Explanation:
No strings in strs can be rearranged to form "bat".
The strings "nat" and "tan" are anagrams because they can be rearranged to form each other.
The strings "ate", "eat", and "tea" are anagrams because they can be rearranged to form each other.
*/
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		// b a t
		chars := strings.Split(s, "")
		// a b t
		sort.Strings(chars)

		// abt
		key := strings.Join(chars, "")

		// [abt:[bat] aet:[eat tea ate] ant:[tan nat]]
		// groups[aet] = [], append groups[eat] = [] + [ate]
		// groups[aet] = [ate], append groups[eat] = [ate] + [eat]
		// groups[aet] = [ate, ate], append groups[eat] = [ate, ate] + [tea]
		// groups[aet] = [ate, ate, tea]
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

// ["paper", "apple", "title", "banana", "hello", "hills"]
func groupAnagrams2(strs []string) (groups map[string][]string) {
	groups = make(map[string][]string)
	fmt.Println("========== group Anagrams ==========")

	result := map[string][]string{}
	for _, str := range strs {
		num := 0
		split := strings.Split(str, "")

		// {id:[str], id:[str]}
		groupMapping := map[string][]int{}
		var groupId []int
		for _, s := range split {
			// {0:p, 1:a, 2:e, 3:r}
			if groupMapping[s] == nil {
				groupMapping[s] = append(groupMapping[s], num)
				num++
			}
			groupId = append(groupId, groupMapping[s]...)
		}
		sb := strings.Builder{}
		for _, g := range groupId {
			sb.WriteString(strconv.Itoa(g))
		}
		id := sb.String()
		result[id] = append(result[id], str)
	}
	for k, v := range result {
		groups[k] = v
	}

	results := make([][]string, 0)
	for _, i := range groups {
		results = append(results, i)
	}

	fmt.Println(results)
	return groups
}

// ["paper", "apple", "title", "banana", "hello", "hills"]
func groupAnagrams3(strs []string) map[string][]string {
	fmt.Println("========== group Anagrams ==========")
	groups := make(map[string][]string)

	for _, str := range strs { // paper
		key := getPatternKey(str)
		groups[key] = append(groups[key], str)
	}

	fmt.Println(groups)
	return groups
}

func getPatternKey(s string) string {
	charIndex := make(map[rune]int)
	var pattern []byte
	nextIdx := 0

	for _, ch := range s {
		if idx, exists := charIndex[ch]; exists {
			pattern = append(pattern, byte('a'+idx))
		} else {
			charIndex[ch] = nextIdx
			pattern = append(pattern, byte('a'+nextIdx))
			nextIdx++
		}
	}

	return string(pattern)
}

func longestConsecutive(nums []int) int {
	fmt.Println("========== longest Consecutive ==========")
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	max := 0
	length := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i]-nums[i-1] == 1 {
			length++
			continue
		}
		if length > max {
			max = length
		}
		length = 1
	}
	if length > max {
		max = length
	}
	return max
}

func moveZeroes(nums []int) []int {
	fmt.Println("========== move Zeroes ==========")
	v := 0 // second pointer, record last 0 element the start index

	for i := range nums {
		if nums[i] != 0 {
			nums[v] = nums[i]
			v++
		}
	}

	for i := v; i < len(nums); i++ { // i = last 0 element the start index, from this pointer onwards, all elements are 0
		nums[i] = 0
	}

	return nums
}

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

	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))          // simple
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))  // medium
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // medium
}

func maxArea(height []int) (maxArea int) {
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

func maxArea2(height []int) (maxArea int) {
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

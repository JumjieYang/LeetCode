/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {

	if len(strs) <= 1 {
		return [][]string{strs}
	} else {

		dict := make(map[string][]string)

		for _, v := range strs {
			string_key := SortString(v)
			if dict_v, ok := dict[string_key]; ok {
				dict[string_key] = append(dict_v, v)
			} else {
				// create a entry
				dict[string_key] = []string{v}
			}
		}
		result := make([][]string, 0, len(dict))

		// returning list of values
		for _, value := range dict {
			result = append(result, value)
		}

		return result
	}
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// @lc code=end


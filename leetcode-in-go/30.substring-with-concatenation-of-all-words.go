/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */

// @lc code=start
type Str struct {
	S        string
	SLen     int
	Ans      []int
	WLen     int
	WordLen  int
	WordsLen int
	M        map[string]int
	Tmp      []Tmp
}

type Tmp struct {
	Index int
	Word  string
}

func findSubstring(s string, words []string) []int {

	//init
	str := new(Str)
	str.S = s
	str.SLen = len(s)
	str.WLen = len(words)
	str.WordLen = len(words[0])
	str.WordsLen = str.WordLen * len(words)

	//exit
	if str.SLen < str.WordsLen {
		return []int{}
	}

	//create map
	str.M = map[string]int{}
	for j := 0; j < str.WLen; j++ {
		str.M[words[j]]++
	}

	//do the work
	for i := 0; i < str.WordLen; i++ {
		str.checkWords(i)
		str.reset()
	}

	//return answer
	return str.Ans
}

func (s *Str) checkWords(start int) {

	for i := start; i <= s.SLen-s.WordLen; i += s.WordLen {
		word := s.S[i : i+s.WordLen]

		//check
		if _, ok := s.M[word]; ok {
			s.M[word]--
			s.Tmp = append(s.Tmp, Tmp{i, word})
			for s.M[word] < 0 {
				s.tmpPoll()
			}
		} else {
			s.reset()
		}

		//add to ans
		if len(s.Tmp) == s.WLen {
			s.Ans = append(s.Ans, s.Tmp[0].Index)
			s.tmpPoll()
		}
	}
}

//reset ( map & tmp )
func (s *Str) reset() {
	for _, w := range s.Tmp {
		s.M[w.Word]++
	}
	s.Tmp = s.Tmp[0:0]
}

//poll from tmp
func (s *Str) tmpPoll() {
	s.M[s.Tmp[0].Word]++
	s.Tmp = s.Tmp[1:]
}

// @lc code=end


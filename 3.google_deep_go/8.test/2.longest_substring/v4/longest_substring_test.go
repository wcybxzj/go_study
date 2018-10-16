package main

import "testing"

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"bbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"b", 1},
		{"abcdef", 6},
		{"这里是米克旺", 6},

		//Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbbb", 1},
		{"abcabcabcabcd", 4},

		//Chinese cases
		{"这里是米克旺", 6},
		{"一二三二一", 3},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubstr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s excepted %d",
				actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkLongestSubstring(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8
	for i := 0; i < b.N; i++ { //运行次数N系统自动给出
		actual := lengthOfNonRepeatingSubstr(s)
		if actual != ans {
			b.Errorf("got %d for input %s excepted %d",
				actual, s, ans)
		}
	}
}

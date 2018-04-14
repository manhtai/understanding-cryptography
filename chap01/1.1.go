package main

import "fmt"
import "unicode"
import "sort"

const cipher = `lrvmnir bpr sumvbwvr jx bpr lmiwv yjeryrkbi jx qmbm wi
bpr xjvni mkd ymibrut jx irhx wi bpr riirkvr jx
ymbinlmtmipw utn qmumbr dj w ipmhh but bj rhnvwdmbr bpr
yjeryrkbi jx bpr qmbm mvvjudwko bj yt wkbrusurbmbwjk
lmird jk xjubt trmui jx ibndt
  wb wi kjb mk rmit bmiq bj rashmwk rmvp yjeryrkb mkd wbi
iwokwxwvmkvr mkd ijyr ynib urymwk nkrashmwkrd bj ower m
vjyshrbr rashmkmbwjk jkr cjnhd pmer bj lr fnmhwxwrd mkd
wkiswurd bj invp mk rabrkb bpmb pr vjnhd urmvp bpr ibmbr
jx rkhwopbrkrd ywkd vmsmlhr jx urvjokwgwko ijnkdhrii
ijnkd mkd ipmsrhrii ipmsr w dj kjb drry ytirhx bpr xwkmh
mnbpjuwbt lnb yt rasruwrkvr cwbp qmbm pmi hrxb kj djnlb
bpmb bpr xjhhjcwko wi bpr sujsru msshwvmbwjk mkd
wkbrusurbmbwjk w jxxru yt bprjuwri wk bpr pjsr bpmb bpr
riirkvr jx jqwkmcmk qmumbr cwhh urymwk wkbmvb`

func percentStats() {
	d := map[string]int{}
	chars := []string{}

	t := 0
	for _, c := range cipher {
		if unicode.IsSpace(c) {
			continue
		}
		s := string(c)

		if d[s] == 0 {
			chars = append(chars, s)
		}

		d[s]++
		t++
	}

	sort.Strings(chars)

	for _, s := range chars {
		fmt.Printf("%s: %.4f\n", s, float64(d[s])/float64(t))
	}

}

func tryMap() {
	plain := ""
	theMap := map[rune]rune{
		'a': 'x',
		'b': 't',
		'c': 'w',
		'd': 'd',
		'e': 'v',
		'f': 'q',
		'g': 'z',
		'h': 'l',
		'i': 's',
		'j': 'o',
		'k': 'n',
		'l': 'b',
		'm': 'a',
		'n': 'u',
		'o': 'g',
		'p': 'h',
		'q': 'k',
		'r': 'e',
		's': 'p',
		't': 'y',
		'u': 'r',
		'v': 'c',
		'x': 'f',
		'y': 'm',
		'z': 'g',
		'w': 'i',
	}

	for _, c := range cipher {
		p := c
		if s, ok := theMap[c]; ok {
			p = s
		}
		plain += string(p)
	}

	fmt.Print(plain)
}

func uc11() {
	percentStats()
	tryMap()
}

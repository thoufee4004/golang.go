package main

import (
	"fmt"
	"math"
)
func main()  {
	var s1,s2 string
	fmt.Scan(&s1,&s2)
	SlidingWindow(s1,s2)

}

/*given two strings s,t find the minimum window in
s which will contain all the charecer in t in complexity o(n)
input: s=adobecodebanc,t=abc
output: banc*/

func SlidingWindow(s1, s2 string) string {
	if len(s1) < len(s2) {
		fmt.Println("no minimum window")
		return ""
	}
	Map := make(map[string]int)
	for _, v := range s2 {
		fmt.Println("v",v)
		Map[string(v)]++
		fmt.Println("Map[string(v)]++",Map[string(v)])
	}
	cnt, l, Size := 0, 0, len(s2)
	res, minLen := "", math.MaxInt32
	currMap := make(map[string]int)
	for r := range s2{
		fmt.Println("r->", r)
		key := string(s1[r])
		fmt.Println("key->", key)
		if Map[key] <= 0 {
			fmt.Println("Map[key]->", Map[key])
			continue
		}
		if currMap[key] < Map[key] {
			fmt.Println("currMap[key]->", currMap[key])
			fmt.Println("Map[key]->", Map[key])
			cnt++
		}
		currMap[key]++
		fmt.Println("currMap[key]++ ->", currMap[key])

		if cnt == Size {
			fmt.Println("cnt", cnt)
			fmt.Println("Size", Size)
			k := string(s1[l])
			fmt.Println("s1[l]", s1[l])
			fmt.Println("k", k)
			for Map[k] <= 0 || currMap[k] > Map[k] {
				if Map[k] > 0 && currMap[k] > Map[k] {
					currMap[k]--
					fmt.Println("currMap[k]--", currMap[k])
				}
				l++
				fmt.Println("l++", l)
				k = string(s1[l])
				fmt.Println("string(s1[l])", k)
			}
			if r-l+1 < minLen {
				minLen = r - l + 1
				fmt.Println("minLen",minLen)
				res = s1[l : r+1]
				fmt.Println("res",res)
			}
		}

	}
	fmt.Println(res)
	return res
}



	/*for i,v := range s1{
		fmt.Println(i,string(v))
	}*/


/*func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	goalMap := make(map[string]int)
	for _, v := range t {
		goalMap[string(v)]++
	}

	res, minLen := "", math.MaxInt32
	l, cnt, golaSize := 0, 0, len(t)
	currMap := make(map[string]int)
	for r := 0; r < len(s); r++ {
		key := string(s[r])
		if goalMap[key] <= 0 {
			continue
		}

		if currMap[key] < goalMap[key] {
			cnt++
		}
		currMap[key]++
		if cnt == golaSize {
			k := string(s[l])
			for goalMap[k] <= 0 || currMap[k] > goalMap[k] {
				if goalMap[k] > 0 && currMap[k] > goalMap[k] {
					currMap[k]--
				}
				l++
				k = string(s[l])
			}
			if r-l+1 < minLen {
				minLen = r - l + 1
				res = s[l : r+1]
			}
		}

	}
	return res
}*/

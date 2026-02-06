package sharecalc

import (
	"sort"
)

func GetSortedMemberBets(memberBet CalMemberBet) []MemberBet {
	keys := make([]int, 0, len(memberBet.MemberBet))
	for k := range memberBet.MemberBet {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	sortedBets := make([]MemberBet, 0, len(keys))
	for _, k := range keys {
		sortedBets = append(sortedBets, memberBet.MemberBet[k])
	}

	return sortedBets
}

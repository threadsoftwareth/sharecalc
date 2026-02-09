package sharecalc

import (
	"sort"
)

func GetSortedMemberBets(memberBet CalMemberBet) []MemberBet {

	sortedBets := make([]MemberBet, 0, len(memberBet.MemberBet))
	for _, v := range memberBet.MemberBet {
		sortedBets = append(sortedBets, v)
	}

	// 2. Sort ตาม Level จากมากไปน้อย (Descending Order)
	sort.Slice(sortedBets, func(i, j int) bool {
		return sortedBets[i].Level > sortedBets[j].Level
	})

	return sortedBets
}

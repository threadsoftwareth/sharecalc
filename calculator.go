package sharecalc

import "fmt"

func CalculateShares(memberBet CalMemberBet) {
	err := Validate(memberBet)
	if err != nil {
		return
	}

	sortedMemberBets := GetSortedMemberBets(memberBet)
	for _, memberBet := range sortedMemberBets {
		fmt.Println(memberBet)
	}
}

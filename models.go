package sharecalc

type CalMemberBet struct {
	MemberID  uint              `json:"member_id"`
	MID       string            `json:"mID"`
	Amount    float64           `json:"amount"`
	MemberBet map[int]MemberBet `json:"member_bet"`
}

type MemberBet struct {
	MemberID         uint    `json:"member_id"`
	ParentID         uint    `json:"parent_id"`
	Level            int     `json:"level"`
	StakePercentTake float64 `json:"stake_percent_take"`
	CommPercentTake  float64 `json:"comm_percent_take"`
	StakePercentBet  float64 `json:"stake_percent_bet"`
	CommPercentBet   float64 `json:"comm_percent_bet"`
}

package sharecalc

type CalMemberBet struct {
	MemberID  uint               `json:"member_id"`
	MID       string             `json:"mID"`
	Amount    float64            `json:"amount"`
	MemberBet map[uint]MemberBet `json:"member_bet"` // key is MemberID
}

type MemberBet struct {
	MemberID     uint    `json:"member_id"`
	ParentID     uint    `json:"parent_id"`
	Level        int     `json:"level"`
	GivePt       float64 `json:"give_pt"`
	KeepPt       float64 `json:"keep_pt"`
	ForcePt      float64 `json:"force_pt"`
	TakeRemainPT float64 `json:"take_remain_pt"`
	Commission   float64 `json:"commission"`
}

type ForPayoutShareBetRequest struct {
	MemberID  uint                   `json:"member_id"`
	MID       string                 `json:"mID"`
	Amount    float64                `json:"amount"`
	Payout    float64                `json:"payout"`
	Winloss   float64                `json:"win_loss"`
	MemberBet map[int]ShareBetResult `json:"member_bet"` // key is Level
}

type ShareBetResult struct {
	MemberID uint `json:"member_id"`
	ParentID uint `json:"parent_id"`
	Level    int  `json:"level"`
	// bet
	StakePercentBet float64 `json:"stake_percent_bet"`
	CommPercentBet  float64 `json:"comm_percent_bet"`
	StakeBet        float64 `json:"stake_bet"`
	CommBet         float64 `json:"comm_bet"`
	WinLossBet      float64 `json:"win_loss_bet"`
	// take
	StakePercentTake float64 `json:"stake_percent_take"`
	CommPercentTake  float64 `json:"comm_percent_take"`
	StakeTake        float64 `json:"stake_take"`
	CommTake         float64 `json:"comm_take"`
	WinLossTake      float64 `json:"win_loss_take"`
}

type ShareBetReportSum struct {
	ID_PARENT   uint    `json:"ID_PARENT"`
	ID_MEMBER   uint    `json:"ID_MEMBER"`
	Oddtype     string  `json:"oddtype"`
	Reportdate  string  `json:"reportdate"`
	Stake       float64 `json:"stake"`
	StakeTake   float64 `json:"stake_take"`
	CommTake    float64 `json:"comm_take"`
	WinlossTake float64 `json:"winloss_take"`
	StakeBet    float64 `json:"stake_bet"`
	CommBet     float64 `json:"comm_bet"`
	WinlossBet  float64 `json:"winloss_bet"`
}

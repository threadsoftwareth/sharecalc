package sharecalc

type MemberBet struct {
	UserID     string  `json:"user_id"`
	ParentID   string  `json:"parent_id"`
	Level      int     `json:"level"`
	Keep       float64 `json:"keep"`
	Give       float64 `json:"give"`
	TakeRemain float64 `json:"take_remain"`
}

type CalculationResult struct {
	UserID      string  `json:"user_id"`
	Level       int     `json:"level"`
	HoldPercent float64 `json:"hold_percent"`
	HoldAmount  float64 `json:"hold_amount"`
}

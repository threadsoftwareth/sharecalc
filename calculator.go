package sharecalc

import (
	"encoding/json"
	"fmt"
)

func CalculateShareBets(memberBetData string, bet_amount float64, MemberID uint) (string, error) {

	var MemberBet map[uint]MemberBet
	if err := json.Unmarshal([]byte(memberBetData), &MemberBet); err != nil {
		return "", fmt.Errorf("failed to unmarshal member bet: %v", err)
	}
	current, ok := MemberBet[MemberID]
	if !ok {
		return "", fmt.Errorf("member not found")
	}

	stake := bet_amount
	pt_remain := 100.00
	pt_nouse := 0.00
	pt_totaltake := 0.00
	attributes := map[int]ShareBetResult{}

	for current.Level > 0 && current.ParentID != 0 {
		member := current
		parent, ok := MemberBet[member.ParentID]
		if !ok {
			return "", fmt.Errorf("parent member not found for member id %d", member.MemberID)
		}

		current = parent
		level := parent.Level

		pt_takeremain := member.TakeRemainPT
		pt_keep := member.KeepPt
		parent_pt_give := parent.GivePt
		parent_pt_force := parent.ForcePt

		pt_takepercent := 0.00
		if level > 0 {
			if pt_takeremain > 0 {
				tempValue := pt_nouse
				if pt_nouse >= pt_takeremain {
					tempValue = pt_takeremain
				}
				pt_takepercent = pt_keep + tempValue
			} else {
				pt_takepercent = pt_keep
			}

			if parent_pt_force > (pt_totaltake + pt_takepercent) {
				pt_takepercent = parent_pt_force - pt_totaltake
			}
		} else {

			pt_takepercent = pt_remain
		}

		pt_totaltake += pt_takepercent
		pt_nouse = (parent_pt_give - pt_totaltake)

		stake_take := pt_takepercent * stake * 0.01
		stake_bet := pt_remain * stake * 0.01

		current_pt_remain := pt_remain
		pt_remain -= pt_takepercent

		if (100.0-pt_remain) > parent_pt_give && level > 0 {
			return "", fmt.Errorf("Incorrect PT!! member id %d", member.MemberID)
		}

		attributes[member.Level] = ShareBetResult{
			MemberID:         member.MemberID,
			ParentID:         member.ParentID,
			Level:            member.Level,
			StakePercentBet:  current_pt_remain,
			CommPercentBet:   member.Commission,
			StakeBet:         stake_bet,
			StakePercentTake: pt_takepercent,
			StakeTake:        stake_take,
			CommPercentTake:  parent.Commission,
		}
	}

	jsonData, _ := json.Marshal(attributes)
	return string(jsonData), nil
}

func CalculateSharePayout(memberBetData string, pay_out float64, validAmount float64, MemberID uint, reportdate string, oddtype string) (string, string, error) {
	var memberBet map[uint]ShareBetResult
	if err := json.Unmarshal([]byte(memberBetData), &memberBet); err != nil {
		return "", "", fmt.Errorf("failed to unmarshal member bet: %v", err)
	}

	win_loss := pay_out - validAmount

	attributes := []ShareBetReportSum{}
	for k, v := range memberBet {

		// payout_member := pay_out * (v.StakePercentBet * 0.01)
		v.WinLossBet += v.StakePercentBet * win_loss * 0.01
		v.WinLossTake += v.StakePercentTake * win_loss * -0.01

		// คิดค่าคอม
		betpercent := v.StakePercentBet * 0.01
		takepercent := v.StakePercentTake * 0.01
		available_bet := validAmount * betpercent
		available_take := validAmount * takepercent

		comm_bet := 0.01 * available_bet * v.CommPercentBet * 0.01
		parent_comm_bet := 0.01 * (available_bet - available_take) * v.CommPercentTake * 0.01

		v.CommBet += comm_bet
		v.CommTake += parent_comm_bet - comm_bet

		memberBet[k] = v

		attributes = append(attributes, ShareBetReportSum{
			ID_PARENT:   v.ParentID,
			ID_MEMBER:   v.MemberID,
			Oddtype:     oddtype,
			Reportdate:  reportdate,
			Stake:       v.StakeBet,
			StakeTake:   v.StakeTake,
			CommTake:    v.CommTake,
			WinlossTake: v.WinLossTake,
			StakeBet:    v.StakeBet,
			CommBet:     v.CommBet,
			WinlossBet:  v.WinLossBet,
		})
	}

	jsonData, _ := json.Marshal(memberBet)
	jsonAttributes, _ := json.Marshal(attributes)
	return string(jsonData), string(jsonAttributes), nil
}

package sharecalc

import (
	"encoding/json"
	"fmt"
)

func CalculateShareBets(memberBet CalMemberBet, MemberID uint) (map[int]ShareBetResult, error) {
	// 1. ดึงสมาชิกเริ่มต้น (Player)
	current, ok := memberBet.MemberBet[MemberID]
	if !ok {
		return nil, fmt.Errorf("member not found")
	}

	stake := memberBet.Amount
	pt_remain := 100.00
	pt_nouse := 0.00
	pt_totaltake := 0.00
	attributes := map[int]ShareBetResult{}

	for current.Level > 0 && current.ParentID != 0 {

		member := current
		parent, ok := memberBet.MemberBet[member.ParentID]
		if !ok {
			return nil, fmt.Errorf("parent member not found for member id %d", member.MemberID)
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
			return nil, fmt.Errorf("Incorrect PT!! member id %d", member.MemberID)
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
		fmt.Println("attributes:", attributes)
	}

	jsonData, _ := json.MarshalIndent(attributes, "", "  ")
	fmt.Println(string(jsonData))
	return attributes, nil
}

func CalculateSharePayouts(pay_out float64, validAmount float64, memberBet ForPayoutShareBetRequest, MemberID uint) (ForPayoutShareBetRequest, error) {
	// stake := memberBet.Amount
	win_loss := pay_out - validAmount

	memberBet.Payout = pay_out
	memberBet.Winloss = win_loss

	for k, v := range memberBet.MemberBet {

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

		memberBet.MemberBet[k] = v
	}

	return memberBet, nil
}

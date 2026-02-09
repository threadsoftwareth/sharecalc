package sharecalc

import (
	"fmt"
	"strings"
)

func Validate(memberBet CalMemberBet) error {
	if len(memberBet.MemberBet) == 0 {
		return fmt.Errorf("member bet is empty")
	}

	if memberBet.Amount == 0 {
		return fmt.Errorf("member bet amount is zero")
	}

	for key, memberBet := range memberBet.MemberBet {
		if memberBet.ParentID == 0 {
			return fmt.Errorf("member bet parent id is zero")
		}
		if memberBet.MemberID == 0 {
			return fmt.Errorf("member bet member id is zero")
		}
		if memberBet.MemberID != uint(key) {
			return fmt.Errorf("member bet memberID %d does not match map key %d", memberBet.Level, key)
		}
	}

	if err := validateMID(memberBet); err != nil {
		return err
	}

	return nil
}

func validateMID(memberBet CalMemberBet) error {
	midList := strings.Split(memberBet.MID, ",")
	parentIDs := make(map[string]bool)

	for _, mb := range memberBet.MemberBet {
		parentIDStr := fmt.Sprintf("%d", mb.ParentID)
		parentIDs[parentIDStr] = true
	}

	if len(midList) != len(parentIDs) {
		return fmt.Errorf("MID count does not match ParentID count")
	}

	for _, mid := range midList {
		mid = strings.TrimSpace(mid)
		if !parentIDs[mid] {
			return fmt.Errorf("MID %s not found in ParentIDs", mid)
		}
	}

	return nil
}

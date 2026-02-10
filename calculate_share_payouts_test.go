package sharecalc

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCalculateSharePayouts(t *testing.T) {
	// Setup test data
	memberBet := ForPayoutShareBetRequest{
		MemberID: 4884314,
		MID:      "TEST001",
		Amount:   100.0,
		MemberBet: map[int]ShareBetResult{
			0: {
				MemberID:         20463,
				ParentID:         1,
				Level:            0,
				StakePercentBet:  10,
				CommPercentBet:   0,
				StakeBet:         10,
				CommBet:          0,
				WinLossBet:       0,
				StakePercentTake: 10,
				CommPercentTake:  0,
				StakeTake:        10,
				CommTake:         0,
				WinLossTake:      0,
			},
			1: {
				MemberID:         4498374,
				ParentID:         20463,
				Level:            1,
				StakePercentBet:  15,
				CommPercentBet:   0,
				StakeBet:         15,
				CommBet:          0,
				WinLossBet:       0,
				StakePercentTake: 5,
				CommPercentTake:  0,
				StakeTake:        5,
				CommTake:         0,
				WinLossTake:      0,
			},
			4: {
				MemberID:         4884314,
				ParentID:         4498374,
				Level:            4,
				StakePercentBet:  100,
				CommPercentBet:   0,
				StakeBet:         100,
				CommBet:          0,
				WinLossBet:       0,
				StakePercentTake: 85,
				CommPercentTake:  0,
				StakeTake:        85,
				CommTake:         0,
				WinLossTake:      0,
			},
		},
	}

	t.Run("Win Scenario", func(t *testing.T) {
		payOut := 200.0
		validAmount := 100.0
		memberID := uint(4884314)

		result, err := CalculateSharePayouts(payOut, validAmount, memberBet, memberID)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if result.Payout != payOut {
			t.Errorf("Expected Payout = %.2f, got %.2f", payOut, result.Payout)
		}

		expectedWinloss := payOut - validAmount
		if result.Winloss != expectedWinloss {
			t.Errorf("Expected Winloss = %.2f, got %.2f", expectedWinloss, result.Winloss)
		}

		jsonData, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println("\n=== Win Scenario Result ===")
		fmt.Println(string(jsonData))

		for level, member := range result.MemberBet {
			fmt.Printf("\nLevel %d (Member ID: %d):\n", level, member.MemberID)
			fmt.Printf("  WinLossBet: %.4f\n", member.WinLossBet)
			fmt.Printf("  WinLossTake: %.4f\n", member.WinLossTake)
			fmt.Printf("  CommBet: %.4f\n", member.CommBet)
			fmt.Printf("  CommTake: %.4f\n", member.CommTake)
		}
	})

	t.Run("Loss Scenario", func(t *testing.T) {
		payOut := 50.0
		validAmount := 100.0
		memberID := uint(4884314)

		result, err := CalculateSharePayouts(payOut, validAmount, memberBet, memberID)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if result.Payout != payOut {
			t.Errorf("Expected Payout = %.2f, got %.2f", payOut, result.Payout)
		}

		expectedWinloss := payOut - validAmount
		if result.Winloss != expectedWinloss {
			t.Errorf("Expected Winloss = %.2f, got %.2f", expectedWinloss, result.Winloss)
		}

		jsonData, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println("\n=== Loss Scenario Result ===")
		fmt.Println(string(jsonData))

		for level, member := range result.MemberBet {
			fmt.Printf("\nLevel %d (Member ID: %d):\n", level, member.MemberID)
			fmt.Printf("  WinLossBet: %.4f\n", member.WinLossBet)
			fmt.Printf("  WinLossTake: %.4f\n", member.WinLossTake)
			fmt.Printf("  CommBet: %.4f\n", member.CommBet)
			fmt.Printf("  CommTake: %.4f\n", member.CommTake)
		}
	})

	t.Run("Break Even Scenario", func(t *testing.T) {
		payOut := 100.0
		validAmount := 100.0
		memberID := uint(4884314)

		result, err := CalculateSharePayouts(payOut, validAmount, memberBet, memberID)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if result.Payout != payOut {
			t.Errorf("Expected Payout = %.2f, got %.2f", payOut, result.Payout)
		}

		expectedWinloss := 0.0
		if result.Winloss != expectedWinloss {
			t.Errorf("Expected Winloss = %.2f, got %.2f", expectedWinloss, result.Winloss)
		}

		jsonData, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println("\n=== Break Even Scenario Result ===")
		fmt.Println(string(jsonData))
	})
}

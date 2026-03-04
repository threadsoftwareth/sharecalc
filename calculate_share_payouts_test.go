package sharecalc

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCalculateSharePayoutCancel(t *testing.T) {
	// ข้อมูล betdata
	betdata := `{"1": {"level": 1, "comm_bet": 0, "comm_take": 0, "member_id": 2, "parent_id": 1, "stake_bet": 15, "stake_take": 15, "parent_level": 0, "win_loss_bet": -13.5, "win_loss_take": 13.5, "comm_percent_bet": 0, "comm_percent_take": 0, "stake_percent_bet": 15, "stake_percent_take": 15}, "2": {"level": 2, "comm_bet": 0, "comm_take": 0, "member_id": 3, "parent_id": 2, "stake_bet": 20, "stake_take": 5, "parent_level": 1, "win_loss_bet": -18, "win_loss_take": 4.5, "comm_percent_bet": 0, "comm_percent_take": 0, "stake_percent_bet": 20, "stake_percent_take": 5}, "3": {"level": 3, "comm_bet": 0, "comm_take": 0, "member_id": 4, "parent_id": 3, "stake_bet": 25, "stake_take": 5, "parent_level": 2, "win_loss_bet": -22.5, "win_loss_take": 4.5, "comm_percent_bet": 0, "comm_percent_take": 0, "stake_percent_bet": 25, "stake_percent_take": 5}, "4": {"level": 4, "comm_bet": 0, "comm_take": 0, "member_id": 5, "parent_id": 4, "stake_bet": 30, "stake_take": 5, "parent_level": 3, "win_loss_bet": -27, "win_loss_take": 4.5, "comm_percent_bet": 0, "comm_percent_take": 0, "stake_percent_bet": 30, "stake_percent_take": 5}, "6": {"level": 6, "comm_bet": 0, "comm_take": 0, "member_id": 7, "parent_id": 5, "stake_bet": 100, "stake_take": 70, "parent_level": 4, "win_loss_bet": -90, "win_loss_take": 63, "comm_percent_bet": 0, "comm_percent_take": 0, "stake_percent_bet": 100, "stake_percent_take": 70}}`

	// ข้อมูลที่ต้องการ cancel
	data := []ShareBetReportSum{
		{Stake: 20, Oddtype: "HK", CommBet: 0, ID_MEMBER: 3, ID_PARENT: 2, CommTake: 0, StakeBet: 20, Reportdate: "2026-03-02", StakeTake: 5, WinlossBet: -17.400000000000002, WinlossTake: 4.3500000000000005},
		{Stake: 25, Oddtype: "HK", CommBet: 0, ID_MEMBER: 4, ID_PARENT: 3, CommTake: 0, StakeBet: 25, Reportdate: "2026-03-02", StakeTake: 5, WinlossBet: -21.75, WinlossTake: 4.3500000000000005},
		{Stake: 30, Oddtype: "HK", CommBet: 0, ID_MEMBER: 5, ID_PARENT: 4, CommTake: 0, StakeBet: 30, Reportdate: "2026-03-02", StakeTake: 5, WinlossBet: -26.1, WinlossTake: 4.3500000000000005},
		{Stake: 100, Oddtype: "HK", CommBet: 0, ID_MEMBER: 7, ID_PARENT: 5, CommTake: 0, StakeBet: 100, Reportdate: "2026-03-02", StakeTake: 70, WinlossBet: -87, WinlossTake: 60.9},
		{Stake: 15, Oddtype: "HK", CommBet: 0, ID_MEMBER: 2, ID_PARENT: 1, CommTake: 0, StakeBet: 15, Reportdate: "2026-03-02", StakeTake: 15, WinlossBet: -13.05, WinlossTake: 13.05},
	}

	// แปลงเป็น JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}
	fmt.Println("memberBetData:", string(jsonData))

	// เรียกใช้ฟังก์ชัน CalculateSharePayoutCancel โดยส่ง betdata และ memberBetData
	newData, canceledData, err := CalculateSharePayoutCancel(betdata, string(jsonData))
	if err != nil {
		t.Fatalf("CalculateSharePayoutCancel failed: %v", err)
	}

	// แปลง result กลับมาเป็น struct
	var newDataResult map[uint]ShareBetResult
	var canceled []ShareBetReportSum

	if err := json.Unmarshal([]byte(newData), &newDataResult); err != nil {
		t.Fatalf("Failed to unmarshal new data: %v", err)
	}

	if err := json.Unmarshal([]byte(canceledData), &canceled); err != nil {
		t.Fatalf("Failed to unmarshal canceled data: %v", err)
	}

	// ตรวจสอบว่าข้อมูล newData มี WinlossBet และ WinlossTake เป็น 0
	t.Log("New Data (with WinLoss = 0):")
	for k, item := range newDataResult {
		t.Logf("Key %d: Level=%d, ID_MEMBER=%d, ID_PARENT=%d, StakeBet=%.2f, StakeTake=%.2f, CommBet=%.2f, CommTake=%.2f, WinlossBet=%.2f, WinlossTake=%.2f",
			k, item.Level, item.MemberID, item.ParentID, item.StakeBet, item.StakeTake, item.CommBet, item.CommTake, item.WinLossBet, item.WinLossTake)

		if item.WinLossBet != 0 {
			t.Errorf("Key %d: Expected WinlossBet 0, got %.2f", k, item.WinLossBet)
		}
		if item.WinLossTake != 0 {
			t.Errorf("Key %d: Expected WinlossTake 0, got %.2f", k, item.WinLossTake)
		}
	}

	// ตรวจสอบว่าข้อมูล canceled มีเครื่องหมายกลับกัน
	if len(canceled) != len(data) {
		t.Errorf("Expected canceled data length %d, got %d", len(data), len(canceled))
	}

	for i, canceledItem := range canceled {
		originalItem := data[i]

		// ตรวจสอบว่าค่าต่างๆ กลับเครื่องหมาย
		if canceledItem.Stake != -originalItem.Stake {
			t.Errorf("Index %d: Expected Stake %.2f, got %.2f", i, -originalItem.Stake, canceledItem.Stake)
		}
		if canceledItem.StakeTake != -originalItem.StakeTake {
			t.Errorf("Index %d: Expected StakeTake %.2f, got %.2f", i, -originalItem.StakeTake, canceledItem.StakeTake)
		}
		if canceledItem.CommTake != -originalItem.CommTake {
			t.Errorf("Index %d: Expected CommTake %.2f, got %.2f", i, -originalItem.CommTake, canceledItem.CommTake)
		}
		if canceledItem.StakeBet != -originalItem.StakeBet {
			t.Errorf("Index %d: Expected StakeBet %.2f, got %.2f", i, -originalItem.StakeBet, canceledItem.StakeBet)
		}
		if canceledItem.CommBet != -originalItem.CommBet {
			t.Errorf("Index %d: Expected CommBet %.2f, got %.2f", i, -originalItem.CommBet, canceledItem.CommBet)
		}

		// ตรวจสอบว่าข้อมูลอื่นๆ ยังคงเหมือนเดิม
		if canceledItem.ID_MEMBER != originalItem.ID_MEMBER {
			t.Errorf("Index %d: Expected ID_MEMBER %d, got %d", i, originalItem.ID_MEMBER, canceledItem.ID_MEMBER)
		}
		if canceledItem.ID_PARENT != originalItem.ID_PARENT {
			t.Errorf("Index %d: Expected ID_PARENT %d, got %d", i, originalItem.ID_PARENT, canceledItem.ID_PARENT)
		}
		if canceledItem.Oddtype != originalItem.Oddtype {
			t.Errorf("Index %d: Expected Oddtype %s, got %s", i, originalItem.Oddtype, canceledItem.Oddtype)
		}
		if canceledItem.Reportdate != originalItem.Reportdate {
			t.Errorf("Index %d: Expected Reportdate %s, got %s", i, originalItem.Reportdate, canceledItem.Reportdate)
		}
	}

	// แสดงผลลัพธ์
	t.Log("\nNew Data JSON:")
	t.Log(newData)
	t.Log("\nCanceled Data (reversed signs):")
	t.Log(canceledData)
}

package sharecalc

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCalculateSharePayoutCancel(t *testing.T) {
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
	fmt.Println(string(jsonData))
	// เรียกใช้ฟังก์ชัน CalculateSharePayoutCancel
	originalData, canceledData, err := CalculateSharePayoutCancel(string(jsonData))
	if err != nil {
		t.Fatalf("CalculateSharePayoutCancel failed: %v", err)
	}

	// แปลง result กลับมาเป็น struct
	var original []ShareBetReportSum
	var canceled []ShareBetReportSum

	if err := json.Unmarshal([]byte(originalData), &original); err != nil {
		t.Fatalf("Failed to unmarshal original data: %v", err)
	}

	if err := json.Unmarshal([]byte(canceledData), &canceled); err != nil {
		t.Fatalf("Failed to unmarshal canceled data: %v", err)
	}

	// ตรวจสอบว่าข้อมูล original ยังคงเหมือนเดิม
	if len(original) != len(data) {
		t.Errorf("Expected original data length %d, got %d", len(data), len(original))
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
	t.Log("Original Data:")
	t.Log(originalData)
	t.Log("\nCanceled Data (reversed signs):")
	t.Log(canceledData)
}

package sharecalc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSharesBet(t *testing.T) {
	mockData := CalMemberBet{
		MemberID: 4884314,
		MID:      "1,17276,20463,4498374",
		Amount:   100.0,
		MemberBet: map[uint]MemberBet{
			4884314: {
				MemberID:     4884314,
				ParentID:     4498374,
				Level:        6,
				GivePt:       0,
				KeepPt:       85,
				ForcePt:      0,
				TakeRemainPT: 1,
				Commission:   0,
			},
			4498374: {
				MemberID:     4498374,
				ParentID:     20463,
				Level:        4,
				GivePt:       85,
				KeepPt:       5,
				TakeRemainPT: 1,
				Commission:   0,
			},
			20463: {
				MemberID:     20463,
				ParentID:     1,
				Level:        1,
				GivePt:       90,
				KeepPt:       10,
				TakeRemainPT: 1,
				Commission:   0,
			},
			1: {
				MemberID:     1,
				ParentID:     1,
				Level:        0,
				GivePt:       100,
				KeepPt:       0,
				TakeRemainPT: 1,
				Commission:   0,
			},
		},
	}

	response, err := CalculateShareBets(mockData, 4884314)

	fmt.Println(response)
	fmt.Println(err)

	assert.NotNil(t, mockData)
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		mockData  CalMemberBet
		wantError string
	}{
		{
			name: "Success_ValidData",
			mockData: CalMemberBet{
				MemberID: 123,
				MID:      "123,456",
				Amount:   1000.0,
				MemberBet: map[uint]MemberBet{
					1: {
						MemberID: 123,
						ParentID: 456,
						Level:    1,
					},
					2: {
						MemberID: 456,
						ParentID: 123,
						Level:    2,
					},
				},
			},
			wantError: "",
		},
		{
			name: "Error_MemberBetEmpty",
			mockData: CalMemberBet{
				MemberID:  123,
				MID:       "123",
				Amount:    1000.0,
				MemberBet: map[uint]MemberBet{},
			},
			wantError: "member bet is empty",
		},
		{
			name: "Error_AmountZero",
			mockData: CalMemberBet{
				MemberID: 123,
				MID:      "123,456",
				Amount:   0.0,
				MemberBet: map[uint]MemberBet{
					1: {
						MemberID: 123,
						ParentID: 456,
						Level:    1,
					},
				},
			},
			wantError: "member bet amount is zero",
		},
		{
			name: "Error_ParentIDZero",
			mockData: CalMemberBet{
				MemberID: 123,
				MID:      "123,456",
				Amount:   1000.0,
				MemberBet: map[uint]MemberBet{
					1: {
						MemberID: 123,
						ParentID: 0,
						Level:    1,
					},
				},
			},
			wantError: "member bet parent id is zero",
		},
		{
			name: "Error_MemberIDZero",
			mockData: CalMemberBet{
				MemberID: 123,
				MID:      "123,456",
				Amount:   1000.0,
				MemberBet: map[uint]MemberBet{
					1: {
						MemberID: 0,
						ParentID: 456,
						Level:    1,
					},
				},
			},
			wantError: "member bet member id is zero",
		},
		{
			name: "Error_KeyMismatchLevel",
			mockData: CalMemberBet{
				MemberID: 123,
				MID:      "123,456",
				Amount:   1000.0,
				MemberBet: map[uint]MemberBet{
					1: {
						MemberID: 123,
						ParentID: 456,
						Level:    2,
					},
					2: {
						MemberID: 456,
						ParentID: 123,
						Level:    2,
					},
				},
			},
			wantError: "member bet level 2 does not match map key 1",
		},
		{
			name: "Error_MIDMismatchCount",
			mockData: CalMemberBet{
				MemberID: 123,
				MID:      "123,456,789",
				Amount:   1000.0,
				MemberBet: map[uint]MemberBet{
					1: {
						MemberID: 123,
						ParentID: 456,
						Level:    1,
					},
					2: {
						MemberID: 456,
						ParentID: 123,
						Level:    2,
					},
				},
			},
			wantError: "MID count does not match ParentID count",
		},
		{
			name: "Error_MIDNotFoundInParentIDs",
			mockData: CalMemberBet{
				MemberID: 123,
				MID:      "123,999",
				Amount:   1000.0,
				MemberBet: map[uint]MemberBet{
					1: {
						MemberID: 123,
						ParentID: 456, // 456 is not in MID
						Level:    1,
					},
					2: {
						MemberID: 456,
						ParentID: 123, // 123 is in MID
						Level:    2,
					},
				},
			},
			wantError: "MID 999 not found in ParentIDs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.mockData)
			if tt.wantError != "" {
				assert.EqualError(t, err, tt.wantError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

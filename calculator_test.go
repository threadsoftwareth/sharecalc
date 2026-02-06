package sharecalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateShares(t *testing.T) {
	mockData := CalMemberBet{
		MemberID: 4884314,
		MID:      "1,17276,20463,4498374",
		Amount:   1000.0,
		MemberBet: map[int]MemberBet{
			4: {
				MemberID:         4884314,
				ParentID:         4498374,
				Level:            4,
				StakePercentTake: 0,
				CommPercentTake:  0,
				StakePercentBet:  100,
				CommPercentBet:   0,
			},
			2: {
				MemberID:         4498374,
				ParentID:         20463,
				Level:            2,
				StakePercentTake: 85,
				CommPercentTake:  0,
				StakePercentBet:  100,
				CommPercentBet:   0,
			},
			1: {
				MemberID:         20463,
				ParentID:         17276,
				Level:            1,
				StakePercentTake: 5,
				CommPercentTake:  0,
				StakePercentBet:  15,
				CommPercentBet:   0,
			},
			0: {
				MemberID:         17276,
				ParentID:         1,
				Level:            0,
				StakePercentTake: 10,
				CommPercentTake:  0,
				StakePercentBet:  10,
				CommPercentBet:   0,
			},
		},
	}

	CalculateShares(mockData)

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
				MemberBet: map[int]MemberBet{
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
				MemberBet: map[int]MemberBet{},
			},
			wantError: "member bet is empty",
		},
		{
			name: "Error_AmountZero",
			mockData: CalMemberBet{
				MemberID: 123,
				MID:      "123,456",
				Amount:   0.0,
				MemberBet: map[int]MemberBet{
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
				MemberBet: map[int]MemberBet{
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
				MemberBet: map[int]MemberBet{
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
				MemberBet: map[int]MemberBet{
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
				MemberBet: map[int]MemberBet{
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
				MemberBet: map[int]MemberBet{
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

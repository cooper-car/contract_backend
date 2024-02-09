package dto

import (
	"github.com/shopspring/decimal"
)

type UpdateMemberRequestDTO struct {
	Username string `json:"username"`
}

type PostMemberRequestDTO struct {
	Username string `json:"username"`
}

type GetMembersTransactionsYearlyResponseDTO struct {
	Username            string          `json:"username"`
	Type                int             `json:"type"`
	BorrowFee           decimal.Decimal `json:"borrowFee"`
	BorrowFeeCreateTime string          `json:"borrowFeeCreateTime"`
	UserCreateTime      string          `json:"userCreateTime"`
}

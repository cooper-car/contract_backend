package memberService

import (
	"contract/internal/data/dto"
	"fmt"
	"time"
)

func (m MemberService) GetMembersTransactions(
	id string,
	startDateTime string,
	endDateTime string,
) []dto.GetMembersTransactionsResponseDTO {

	var result []dto.GetMembersTransactionsResponseDTO

	// 時間轉換(起始)
	parseStartDateTime, err := time.Parse("2006-01-02T15:04:05", startDateTime)
	if err != nil {
		fmt.Println(err.Error())
	}
	formatStartDateTime := parseStartDateTime.Format("2006-01-02 15:04:05")

	// 時間轉換(結束)
	parseEndDateTime, err := time.Parse("2006-01-02T15:04:05", endDateTime)
	if err != nil {
		fmt.Println(err.Error())
	}
	formatEndDateTime := parseEndDateTime.Format("2006-01-02 15:04:05")

	if id == "all" {
		result = m.memberRepo.GetAllMembersTransactionsByDateRange(formatStartDateTime, formatEndDateTime)
	} else {
		result = m.memberRepo.GetMemberTransactionsByDateRangeAndPk(id, formatStartDateTime, endDateTime)
	}

	return result
}

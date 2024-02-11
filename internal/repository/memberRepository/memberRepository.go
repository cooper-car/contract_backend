package memberRepository

import (
	"contract/internal/data/dto"
	"contract/internal/data/po"
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

func (m MemberRepository) CreateByBath(data []po.MemberPO) {
	batchSize := 10000
	result := m.db.Table("member").CreateInBatches(data, batchSize)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}
}

func (m MemberRepository) GetAllPk() []int {
	var pks []int

	sql := `SELECT pk FROM member`
	result := m.db.Raw(sql).Scan(&pks)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	return pks
}

func (m MemberRepository) UpdateMember(id string, memberPO po.MemberPO) {
	sql := `UPDATE member SET username = ? WHERE pk = ?`
	result := m.db.Exec(sql, memberPO.Username, id)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}
}

func (m MemberRepository) CreateMember(memberPO po.MemberPO) {
	sql := `INSERT INTO member (username, create_time) VALUES (?, ?)`
	result := m.db.Exec(sql, memberPO.Username, memberPO.CreateTime)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}
}

func (m MemberRepository) GetMembers() []dto.GetMembersResponseDTO {
	var members []dto.GetMembersResponseDTO
	sql := `SELECT pk, username FROM member`
	result := m.db.Raw(sql).Scan(&members)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}

	return members
}

func (m MemberRepository) GetAllMembersTransactionsByDateRange(start string, end string) []dto.GetMembersTransactionsResponseDTO {
	sql := `
		SELECT m.username, b.type, b.borrow_fee, b.create_time as borrow_fee_create_time, m.create_time as user_create_time 
		FROM borrow_fee as b JOIN member as m ON b.member_fk = m.pk
		WHERE b.create_time BETWEEN ? AND ? 
		ORDER BY b.create_time ASC;
	`

	var result []dto.GetMembersTransactionsResponseDTO

	rows, err := m.db.Raw(sql, start, end).Rows()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	for rows.Next() {
		var username string
		var types int
		var borrowFee decimal.Decimal
		var borrowFeeCreateTime string
		var userCreateTime string

		rows.Scan(
			&username,
			&types,
			&borrowFee,
			&borrowFeeCreateTime,
			&userCreateTime,
		)

		borrowFeeTime, _ := time.Parse("2006-01-02T15:04:05-07:00", borrowFeeCreateTime)
		borrowFeeDateTime := borrowFeeTime.Format("2006-01-02 15:04:05")

		userTime, _ := time.Parse("2006-01-02T15:04:05-07:00", userCreateTime)
		userDateTime := userTime.Format("2006-01-02 15:04:05")

		result = append(result, dto.GetMembersTransactionsResponseDTO{
			Username:            username,
			Type:                types,
			BorrowFee:           borrowFee,
			BorrowFeeCreateTime: borrowFeeDateTime,
			UserCreateTime:      userDateTime,
		})
	}

	return result
}

func (m MemberRepository) GetMemberTransactionsByDateRangeAndPk(id string, start string, end string) []dto.GetMembersTransactionsResponseDTO {
	sql := `
		SELECT m.username, b.type, b.borrow_fee, b.create_time as borrow_fee_create_time, m.create_time as user_create_time 
		FROM borrow_fee as b JOIN member as m ON b.member_fk = m.pk
		WHERE m.pk = ? AND b.create_time BETWEEN ? AND ? 
		ORDER BY b.create_time ASC;
	`

	var result []dto.GetMembersTransactionsResponseDTO

	rows, err := m.db.Raw(sql, id, start, end).Rows()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	for rows.Next() {
		var username string
		var types int
		var borrowFee decimal.Decimal
		var borrowFeeCreateTime string
		var userCreateTime string

		rows.Scan(
			&username,
			&types,
			&borrowFee,
			&borrowFeeCreateTime,
			&userCreateTime,
		)

		borrowFeeTime, _ := time.Parse("2006-01-02T15:04:05-07:00", borrowFeeCreateTime)
		borrowFeeDateTime := borrowFeeTime.Format("2006-01-02 15:04:05")

		userTime, _ := time.Parse("2006-01-02T15:04:05-07:00", userCreateTime)
		userDateTime := userTime.Format("2006-01-02 15:04:05")

		result = append(result, dto.GetMembersTransactionsResponseDTO{
			Username:            username,
			Type:                types,
			BorrowFee:           borrowFee,
			BorrowFeeCreateTime: borrowFeeDateTime,
			UserCreateTime:      userDateTime,
		})
	}

	return result
}

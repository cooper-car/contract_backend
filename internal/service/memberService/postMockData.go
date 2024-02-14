package memberService

import (
	"contract/internal/data/po"
	"fmt"
	"math/rand"
	"time"
)

func (m MemberService) CreateMockDataForMembers() {

	// 取得現在時間的 Unix - 2 小時
	nowUnix := time.Now().Unix() - 7200

	// 產生 1000 筆 member 資料
	membersData, nowUnix := makeMemberData(nowUnix)

	// 建立 1000 筆 member 資料
	m.memberRepo.CreateByBath(membersData)

	// 取得所有 member 的 pk
	pks := m.memberRepo.GetAllPk()

	// 使用 pks 隨機產生 5000 筆 borrow_fee 資料
	borrowFeeData := makeBorrowFeeData(pks, nowUnix)

	// 建立 5000 筆 borrow_fee 資料
	m.borrowFeeRepo.CreateByBatch(borrowFeeData)
}

func makeBorrowFeeData(pks []int, nowUnix int64) []po.BorrowFeePO {
	var borrowFeeData []po.BorrowFeePO
	for i := 0; i < 5000; i++ {
		types := randType()
		fee := randBorrowFee()
		fk := randMemberFk(pks)
		datetime := time.Unix(nowUnix, 0)

		borrowFeeData = append(borrowFeeData, po.BorrowFeePO{
			MemberFk:   fk,
			Type:       types,
			BorrowFee:  fee,
			CreateTime: datetime.Format("2006-01-02 15:04:05"),
		})

		nowUnix++
	}
	return borrowFeeData
}

func makeMemberData(nowUnix int64) ([]po.MemberPO, int64) {
	var data []po.MemberPO
	for i := 1; i <= 1000; i++ {
		name := fmt.Sprintf("TestUser_%d", i)
		datetime := time.Unix(nowUnix, 0)
		data = append(data, po.MemberPO{
			Username:   name,
			CreateTime: datetime.Format("2006-01-02 15:04:05"),
		})
		nowUnix++
	}
	return data, nowUnix
}

func randType() int {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)

	types := []int{0, 1}
	randType := types[randGen.Intn(len(types))]

	return randType
}

func randBorrowFee() string {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)

	integerPart := randGen.Intn(101)
	decimalPart := randGen.Intn(100)

	fee := fmt.Sprintf("%d.%02d", integerPart, decimalPart)

	return fee
}

func randMemberFk(pks []int) int {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)

	fk := pks[randGen.Intn(len(pks))]
	return fk
}

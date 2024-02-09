package memberRepository

import (
	"contract/internal/data/po"
	"fmt"
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

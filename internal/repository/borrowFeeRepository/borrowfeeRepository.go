package borrowFeeRepository

import (
	"contract/internal/data/po"
	"fmt"
)

func (b BorrowFeeRepository) CreateByBatch(data []po.BorrowFeePO) {
	batchSize := 10000
	result := b.db.Table("borrow_fee").CreateInBatches(data, batchSize)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}
}

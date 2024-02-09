package borrowFeeRepository

import "gorm.io/gorm"

type BorrowFeeRepository struct {
	db *gorm.DB
}

func NewBorrowFeeRepository(db *gorm.DB) *BorrowFeeRepository {
	return &BorrowFeeRepository{
		db: db,
	}
}

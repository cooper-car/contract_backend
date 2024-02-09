package po

type BorrowFeePO struct {
	MemberFk   int    `gorm:"column:member_fk"`
	Type       int    `gorm:"column:type"`
	BorrowFee  string `gorm:"column:borrow_fee"`
	CreateTime string `gorm:"column:create_time"`
}

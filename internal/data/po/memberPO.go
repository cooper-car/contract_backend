package po

type MemberPO struct {
	Username   string `gorm:"column:username"`
	CreateTime string `gorm:"column:create_time"`
}

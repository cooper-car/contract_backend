package memberService

import (
	"contract/internal/data/dto"
	"contract/internal/data/po"
	"time"
)

func (m MemberService) CreateMember(request dto.PostMemberRequestDTO) {
	memberPO := createToMemberPO(request)
	m.memberRepo.CreateMember(memberPO)
}

func createToMemberPO(request dto.PostMemberRequestDTO) po.MemberPO {
	createTime := time.Now().Format("2006-01-02 15:04:05")

	return po.MemberPO{
		Username:   request.Username,
		CreateTime: createTime,
	}
}

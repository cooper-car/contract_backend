package memberService

import (
	"contract/internal/data/dto"
	"contract/internal/data/po"
)

func (m MemberService) UpdateMember(id string, dto dto.UpdateMemberRequestDTO) {
	memberPO := updateToMemberPO(dto)
	m.memberRepo.UpdateMember(id, memberPO)
}

func updateToMemberPO(requestDTO dto.UpdateMemberRequestDTO) po.MemberPO {
	return po.MemberPO{
		Username: requestDTO.Username,
	}
}

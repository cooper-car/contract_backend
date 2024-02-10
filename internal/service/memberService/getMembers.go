package memberService

import "contract/internal/data/dto"

func (m MemberService) GetMembers() []dto.GetMembersResponseDTO {

	return m.memberRepo.GetMembers()
}

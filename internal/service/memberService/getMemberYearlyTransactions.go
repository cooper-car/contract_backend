package memberService

import "contract/internal/data/dto"

func (m MemberService) GetMembersTransactionsYearly() []dto.GetMembersTransactionsYearlyResponseDTO {

	return m.memberRepo.GetMembersTransactionsYearly()
}

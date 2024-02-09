package memberService

import (
	"contract/internal/repository/borrowFeeRepository"
	"contract/internal/repository/memberRepository"
)

type MemberService struct {
	memberRepo    *memberRepository.MemberRepository
	borrowFeeRepo *borrowFeeRepository.BorrowFeeRepository
}

func NewMemberService(
	memberRepo *memberRepository.MemberRepository,
	borrowFeeRepo *borrowFeeRepository.BorrowFeeRepository,
) *MemberService {
	return &MemberService{
		memberRepo:    memberRepo,
		borrowFeeRepo: borrowFeeRepo,
	}
}

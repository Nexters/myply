package member

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(NewMemberService)

type MemberService interface {
	SignUp(deviceToken string, name string) error
}

type memberService struct {
	repo MemberRepository
}

func NewMemberService(repo MemberRepository) MemberService {
	return &memberService{repo: repo}
}

func (ms *memberService) SignUp(deviceToken string, name string) error {
	entity := Member{DeviceToken: deviceToken, Name: name}

	return ms.repo.Create(entity)
}

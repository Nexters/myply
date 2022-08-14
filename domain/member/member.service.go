package member

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(NewMemberService)

type MemberService interface {
	SignUp(deviceToken string, name string, keywords []string) error
	Update(
		deviceToken string,
		name *string,
		keywords []string,
	) error
}

type memberService struct {
	repo MemberRepository
}

func NewMemberService(repo MemberRepository) MemberService {
	return &memberService{repo: repo}
}

func (ms *memberService) SignUp(deviceToken string, name string, keywords []string) error {
	entity := Member{DeviceToken: deviceToken, Name: name, Keywords: keywords}

	return ms.repo.Create(entity)
}

func (ms *memberService) Update(
	deviceToken string,
	name *string,
	keywords []string,
) error {
	return ms.repo.Update(deviceToken, name, keywords)
}

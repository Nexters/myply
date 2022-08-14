package member

import (
	"strings"

	"github.com/google/uuid"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewMemberService)

type MemberService interface {
	SignUp(deviceToken *string, name string, keywords []string) (*Member, error)
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

func (ms *memberService) SignUp(deviceToken *string, name string, keywords []string) (*Member, error) {
	var token string
	if deviceToken == nil {
		token = strings.ReplaceAll(uuid.New().String(), "-", "")
		exist, err := ms.repo.Get(token)
		if err != nil {
			return nil, err
		}
		if exist != nil {
			return ms.SignUp(deviceToken, name, keywords)
		}
	} else {
		token = *deviceToken
	}

	entity := Member{DeviceToken: token, Name: name, Keywords: keywords}
	if err := ms.repo.Create(entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (ms *memberService) Update(
	deviceToken string,
	name *string,
	keywords []string,
) error {
	return ms.repo.Update(deviceToken, name, keywords)
}

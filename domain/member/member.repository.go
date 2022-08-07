package member

type MemberRepository interface {
	Create(entity Member) error
}

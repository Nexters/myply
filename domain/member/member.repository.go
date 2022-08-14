package member

type MemberRepository interface {
	Get(deviceToken string) (*Member, error)
	Create(entity Member) error
	Update(
		deviceToken string,
		name *string,
		keywords []string,
	) (*Member, error)
}

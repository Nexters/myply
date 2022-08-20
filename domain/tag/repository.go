package tag

type TagRepository interface {
	Recommend() (*Tags, error)
}

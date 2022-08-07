package tag

import "github.com/google/wire"

var Set = wire.NewSet(NewTagService)

type TagService interface {
	Recommend() (*Tags, error)
}

type tagService struct {
	repo TagRepository
}

func NewTagService(repo TagRepository) TagService {
	return &tagService{repo: repo}
}

func (ts tagService) Recommend() (*Tags, error) {
	return ts.repo.Recommend()
}

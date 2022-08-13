package persistence

import (
	"regexp"

	"github.com/Nexters/myply/domain/tag"
	"github.com/Nexters/myply/infrastructure/persistence/data"
)

type tagRepository struct{}

func NewTagRepository() tag.TagRepository {
	return &tagRepository{}
}

func (tr tagRepository) Recommend() (*tag.Tags, error) {
	str := data.RecommendTags
	re := regexp.MustCompile("[\n|,]")
	labels := re.Split(str, -1)
	tags := tag.NewTagsByLabel(labels)

	return tags, nil
}

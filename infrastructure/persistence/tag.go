package persistence

import (
	"path/filepath"

	"github.com/Nexters/myply/domain/tag"
	"github.com/Nexters/myply/infrastructure/persistence/fs"
)

type tagRepository struct{}

func NewTagRepository() tag.TagRepository {
	return &tagRepository{}
}

func (tr tagRepository) Recommend() (*tag.Tags, error) {
	filePath, err := filepath.Abs("infrastructure/persistence/data/tags_recommend.csv")
	if err != nil {
		return nil, err
	}

	csvManager := fs.CSVManger{
		FilePath: filePath,
	}
	labels, err := csvManager.Data()
	if err != nil {
		return nil, err
	}

	tags := tag.NewTagsByLabel(labels)

	return tags, nil
}

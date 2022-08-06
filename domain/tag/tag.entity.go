package tag

import (
	"context"
	"strings"

	rxgo "github.com/reactivex/rxgo/v2"
)

type Tag struct {
	Label string
}

type Tags struct {
	value []Tag
}

func NewTagsByLabel(labels []string) *Tags {
	tags := make([]Tag, len(labels))
	for i := range labels {
		tags[i] = Tag{Label: labels[i]}
	}

	return &Tags{value: tags}
}

func (tags *Tags) Labels() []string {
	var labels []string

	observable := rxgo.
		Just(tags.value)().
		Map(func(_ context.Context, item interface{}) (interface{}, error) {
			return strings.TrimSpace(item.(Tag).Label), nil
		}).
		Filter(func(item interface{}) bool {
			return len(item.(string)) > 0
		})

	for label := range observable.Observe() {
		labels = append(labels, label.V.(string))
	}

	return labels
}

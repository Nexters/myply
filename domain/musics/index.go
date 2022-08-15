package musics

import (
	"fmt"
)

func GenerateCacheKey(query, token string) string {
	return fmt.Sprintf("%s:%s", query, token)
}

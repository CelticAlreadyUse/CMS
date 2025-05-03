package helper

import (
	"fmt"

)
const NewsBucketKey = "news"

func NewNewsByIDCacheKey(id int64) string {
	return fmt.Sprintf("news:%d", id)
}
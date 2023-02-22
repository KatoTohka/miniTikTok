package cache

import (
	"fmt"
	"miniTikTok/pkg/constants"
	"os"
)

func Cache(data []byte, key int64) error {
	err := os.WriteFile(constants.TempCachePlace+fmt.Sprintf("%v.mp4", key), data, 0664)
	return err
}

package cache

import (
	"fmt"
	"os"
)

func Cache(data []byte, key int64) error {
	err := os.WriteFile(fmt.Sprintf("/home/shs/miniTikTok/videoImage/%v.mp4", key), data, 0664)
	return err
}

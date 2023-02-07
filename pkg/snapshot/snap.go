package snapshot

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"miniTikTok/pkg/constants"
	"os"
)

// GetSnapshot https://github.com/u2takey/ffmpeg-go#task-frame-from-video
func GetSnapshot(inFileName string, frameNum int, imgName string) ([]byte, error) {
	snapshotPath := constants.TempCachePlace + imgName
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Println("frame get err", err)
		return []byte{}, err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Println("snapshot get err", err)
		return []byte{}, err
	}
	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("save snapshot err：", err)
		return []byte{}, err
	}
	data, err := os.ReadFile(snapshotPath + ".png")
	if err != nil {
		return nil, err
	}
	return data, nil
}

package snapshot

//package main
//
//import (
//	"bytes"
//	"fmt"
//	"github.com/disintegration/imaging"
//	ffmpeg "github.com/u2takey/ffmpeg-go"
//	"log"
//	"os"
//)
//
//// GetSnapshot https://github.com/u2takey/ffmpeg-go#task-frame-from-video
//func GetSnapshot(inFileName string, frameNum int, imgName string) (string, error) {
//	snapshotPath := "../..videoImage/" + imgName
//	buf := bytes.NewBuffer(nil)
//	err := ffmpeg.Input(inFileName).
//		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
//		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
//		WithOutput(buf, os.Stdout).
//		Run()
//	if err != nil {
//		panic(err)
//	}
//	img, err := imaging.Decode(buf)
//	if err != nil {
//		log.Fatal(err)
//		return "", err
//	}
//
//	err = imaging.Save(img, snapshotPath+".png")
//	if err != nil {
//		log.Fatal(err)
//		return "", err
//	}
//
//	imgPath := snapshotPath + ".png"
//
//	return imgPath, nil
//}
//
//func main() {
//	_, err := GetSnapshot("./bear.mp4", 1, "bear")
//	if err != nil {
//		return
//	}
//}

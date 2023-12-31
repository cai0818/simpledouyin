package Tool

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"strings"
)

func main() {
	name, err := GetSnapshot("./static/1.mp4", "testImage", 1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("return:", name)
}

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	snapshotPath = "static/picture/" + snapshotPath
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	fmt.Println("--snapshotPath--", snapshotPath)
	// --snapshotPath-- ./assets/testImage

	names := strings.Split(snapshotPath, "\\")
	fmt.Println("----names----", names)
	// ----names---- [./assets/testImage]
	// 这里把 snapshotPath 的 string 类型转换成 []string

	snapshotName = names[len(names)-1] + ".png"
	fmt.Println("----snapshotName----", snapshotName)
	// ----snapshotName---- ./assets/testImage.png

	return snapshotName, nil
}

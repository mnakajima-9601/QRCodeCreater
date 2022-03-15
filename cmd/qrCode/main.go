package main

import (
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/google/uuid"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Conf struct {
	Out      string
	ListFile string
	CsvFile  string
	Size     int
	TtfFile  string
}

var count int
var dir string
var conf Conf
var codeList []string
var noList []string
var title1List []string
var title2List []string
var title3List []string
var information1List []string
var information2List []string
var information3List []string
var margin int
var qrX int
var qrY int
var outPut string
var csvPath string
var uuFlg string
var uu string
var qrCode string
var fileName string

const (
	x             = 0
	y             = 0
	fontSize      = 15  // point
	fontSize2     = 20  // point
	imageWidth    = 320 // pixel
	imageHeight   = 516 // pixel
	textTopMargin = 80  // fixed.I
	qrDefaultX    = 61
	qrDefaultY    = 160
)

func main() {
	if len(os.Args) > 1 {
		outPut = os.Args[1]
		csvPath = os.Args[2]
	}
	if len(os.Args) == 4 {
		uuFlg = os.Args[3]
	}
	// ユーザディレクトリ取得
	p, _ := os.UserHomeDir()
	dir = p + "/Desktop/"
	//パラメーターをxmlファイルから取得する
	getConf()
	// uuid用のファイル名を作成する
	createFileName()
	//CSVを読み込む
	getCsv()
	// 画像作成
	createImg()
}

func getConf() {
	p, _ := os.Getwd()
	confname := ""
	confname = p + "/conf.xml"
	data, _ := ioutil.ReadFile(confname)
	err := xml.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	conf.CsvFile = dir + conf.CsvFile
	conf.Out = dir + conf.Out
}

func createFileName() {
	if uuFlg == "1" {
		if outPut == "" {
			outPut = conf.Out
		}
		day := time.Now()
		var layout = "20060102"
		fileName = outPut + "_" + day.Format(layout) + "_作成UUID一覧.txt"
	}
}

func getCsv() {

	if csvPath == "" {
		csvPath = conf.CsvFile
	}
	file, err := os.Open(csvPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string

	for {
		line, err = reader.Read()
		if err != nil {
			break
		}
		if uuFlg == "1" {
			createUUID()
			qrCode = line[0] + uu
			writeFile()
		} else {
			qrCode = line[0]
		}
		codeList = append(codeList, qrCode)
		noList = append(noList, line[1])
		title1List = append(title1List, line[2])
		title2List = append(title2List, line[3])
		title3List = append(title3List, line[4])
		information1List = append(information1List, line[5])
		information2List = append(information2List, line[6])
		information3List = append(information3List, line[7])
	}

}
func createUUID() {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	uu = u.String()
}

func writeFile() {
	// ファイルを書き込み用にオープン
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// テキストを書き込む
	_, err = file.WriteString(uu + "\n")
	if err != nil {
		panic(err)
	}
}

func createImg() {
	f, err := os.Open(conf.TtfFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ttfBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	ft, err := truetype.Parse(ttfBytes)
	if err != nil {
		panic(err)
	}

	opt := truetype.Options{
		Size:              fontSize,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	opt2 := truetype.Options{
		Size:              fontSize2,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	count = 0
	for _, no := range noList {

		if no == "" {
			continue
		}

		face := truetype.NewFace(ft, &opt)
		face2 := truetype.NewFace(ft, &opt2)

		img := image.NewRGBA(image.Rect(x, y, imageWidth, imageHeight))
		// 短形に色を追加
		for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
			for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
				img.Set(j, i, color.RGBA{255, 255, 255, 255})
			}
		}

		// --------------文字入力開始-------------------

		// タイトル１
		dr1 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr1.Dot.X = (fixed.I(imageWidth) - dr1.MeasureString(title1List[count])) / 2
		if title3List[count] == "" || title2List[count] == "" {
			dr1.Dot.Y = fixed.I(100)
		} else if title3List[count] == "" && title2List[count] == "" {
			dr1.Dot.Y = fixed.I(130)
		} else {
			dr1.Dot.Y = fixed.I(70)
		}
		dr1.DrawString(title1List[count])

		// タイトル２
		dr2 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr2.Dot.X = (fixed.I(imageWidth) - dr2.MeasureString(title2List[count])) / 2
		if title3List[count] == "" {
			dr2.Dot.Y = fixed.I(130)
		} else {
			dr2.Dot.Y = fixed.I(100)
		}
		dr2.DrawString(title2List[count])

		// タイトル３
		dr3 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr3.Dot.X = (fixed.I(imageWidth) - dr3.MeasureString(title3List[count])) / 2
		dr3.Dot.Y = fixed.I(130)
		dr3.DrawString(title3List[count])

		// 整理番号
		dr4 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face2,
			Dot:  fixed.Point26_6{},
		}

		dr4.Dot.X = (fixed.I(imageWidth) - dr4.MeasureString("整理番号："+no)) / 2
		dr4.Dot.Y = fixed.I(395)
		dr4.DrawString("整理番号：" + no)

		// 情報１
		dr5 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr5.Dot.X = (fixed.I(imageWidth) - dr5.MeasureString(information1List[count])) / 2
		dr5.Dot.Y = fixed.I(435)
		dr5.DrawString(information1List[count])

		// 情報２
		dr6 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr6.Dot.X = (fixed.I(imageWidth) - dr6.MeasureString(information2List[count])) / 2
		if information1List[count] == "" {
			dr6.Dot.Y = fixed.I(435)
		} else {
			dr6.Dot.Y = fixed.I(460)
		}
		dr6.DrawString(information2List[count])

		// 情報３
		dr7 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr7.Dot.X = (fixed.I(imageWidth) - dr7.MeasureString(information3List[count])) / 2
		if information1List[count] == "" && information2List[count] == "" {
			dr7.Dot.Y = fixed.I(435)
		} else if information1List[count] == "" || information2List[count] == "" {
			dr7.Dot.Y = fixed.I(460)
		} else {
			dr7.Dot.Y = fixed.I(485)
		}
		dr7.DrawString(information3List[count])

		// QRコード
		dr8 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		margin = 0
		qrCode, _ := qr.Encode(codeList[count], qr.M, qr.Auto)
		qrCode, _ = barcode.Scale(qrCode, conf.Size, conf.Size)

		rect := qrCode.Bounds()
		qrX = qrDefaultX
		qrY = qrDefaultY
		for y := rect.Min.Y - margin; y < rect.Max.Y+margin; y++ {
			qrY = qrY + 1
			qrX = qrDefaultX
			for x := rect.Min.X - margin; x < rect.Max.X+margin; x++ {
				if rect.Min.X <= x && x < rect.Max.X &&
					rect.Min.Y <= y && y < rect.Max.Y &&
					qrCode.At(x, y) == color.Black {
					dr8.Dot.X = fixed.I(qrX)
					dr8.Dot.Y = fixed.I(qrY)
					dr8.DrawString(".")
					// img.Set(qrX, qrY, color.RGBA{0, 0, 0, 0})
				} else {
					dr8.Dot.X = fixed.I(qrX)
					dr8.Dot.Y = fixed.I(qrY)
					dr8.DrawString("")
					// img.Set(qrX, qrY, color.RGBA{255, 255, 255, 255})
				}
				qrX = qrX + 1
			}
		}
		qrX = qrDefaultX
		qrY = qrDefaultY

		// --------------文字入力終了-------------------

		buf := &bytes.Buffer{}
		err = png.Encode(buf, img)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if outPut == "" {
			outPut = conf.Out
		}

		file, err := os.Create(outPut + "/" + no + ".png")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer file.Close()

		file.Write(buf.Bytes())

		count = count + 1

	}
}

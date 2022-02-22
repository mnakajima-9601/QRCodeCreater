package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"runtime"

	// "log"
	// "os"
	// "runtime"
	// "time"

	// "gocv.io/x/gocv"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

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
var conf Conf
var codeList []string
var noList []string
var title1List []string
var title2List []string
var title3List []string
var information1List []string
var information2List []string
var information3List []string

const (
	x             = 0
	y             = 0
	fontSize      = 30  // point
	fontSize2     = 40  // point
	imageWidth    = 600 // pixel
	imageHeight   = 800 // pixel
	textTopMargin = 80  // fixed.I
)

func main() {

	//パラメーターをxmlファイルから取得する
	// getConf()
	//QRコードにする文字のリストを読み込む
	// getList()
	//QRコード生成 boombuler/barcodeパッケージ使用
	// createCode()

	//パラメーターをxmlファイルから取得する
	getConf()
	//CSVを読み込む
	getCsv()
	//QRコード生成 boombuler/barcodeパッケージ使用
	createCode()
	// 画像作成
	createImg()

}

func getConf() {
	p, _ := os.Getwd()
	fmt.Println(p)
	confname := ""
	if runtime.GOOS == "windows" {
		confname = p + "\\conf.xml"
	} else {
		confname = p + "/conf.xml"
	}
	data, _ := ioutil.ReadFile(confname)
	err := xml.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
}

func getList() {
	fp, err := os.Open(conf.ListFile)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		codeList = append(codeList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func createCode() {
	for _, code := range codeList {
		if code == "" {
			continue
		}
		qrCode, _ := qr.Encode(code, qr.M, qr.Auto)
		qrCode, _ = barcode.Scale(qrCode, conf.Size, conf.Size)
		file, _ := os.Create(conf.Out + "test.png")
		defer file.Close()
		png.Encode(file, qrCode)
	}
}

func getCsv() {
	file, err := os.Open(conf.CsvFile)
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
		codeList = append(noList, line[0])
		noList = append(noList, line[1])
		title1List = append(title1List, line[2])
		title2List = append(title2List, line[3])
		title3List = append(title3List, line[4])
		information1List = append(information1List, line[5])
		information2List = append(information2List, line[6])
		information3List = append(information3List, line[7])
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

		// b, a := font.BoundString(face, text)
		b, a := font.BoundString(face, no)
		w := b.Max.X - b.Min.X + fixed.I(1)
		h := b.Max.Y - b.Min.Y + fixed.I(1)

		img := image.NewRGBA(image.Rect(x, y, imageWidth, imageHeight))
		// 短形に色を追加
		for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
			for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
				img.Set(j, i, color.RGBA{255, 255, 255, 255})
			}
		}

		// --------------文字入力開始-------------------

		// 整理番号
		dr := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face2,
			Dot:  fixed.Point26_6{},
		}

		dr.Dot.X = fixed.I(230)
		dr.Dot.Y = fixed.I(600)
		dr.DrawString(no)

		// タイトル１
		dr1 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr1.Dot.X = fixed.I(60)
		dr1.Dot.Y = fixed.I(100)
		dr1.DrawString(title1List[count])

		// タイトル２
		dr2 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr2.Dot.X = fixed.I(150)
		dr2.Dot.Y = fixed.I(140)
		dr2.DrawString(title2List[count])

		// タイトル３
		dr3 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr3.Dot.X = fixed.I(250)
		dr3.Dot.Y = fixed.I(180)
		dr3.DrawString(title3List[count])

		// 情報１
		dr4 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr4.Dot.X = fixed.I(235)
		dr4.Dot.Y = fixed.I(650)
		dr4.DrawString(information1List[count])

		// 情報２
		dr5 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr5.Dot.X = fixed.I(235)
		dr5.Dot.Y = fixed.I(690)
		dr5.DrawString(information2List[count])

		// 情報３
		dr6 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr6.Dot.X = fixed.I(210)
		dr6.Dot.Y = fixed.I(730)
		dr6.DrawString(information3List[count])

		// QRコード
		dr7 := &font.Drawer{
			Dst:  img,
			Src:  image.Black,
			Face: face,
			Dot:  fixed.Point26_6{},
		}

		dr7.Dot.X = (w - a) / 2
		dr7.Dot.Y = h - b.Max.Y

		qrCode, _ := qr.Encode(codeList[count], qr.M, qr.Auto)
		qrCode, _ = barcode.Scale(qrCode, conf.Size, conf.Size)

		// dr7.DrawString(qrCode)

		// --------------文字入力終了-------------------

		buf := &bytes.Buffer{}
		err = png.Encode(buf, img)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		file, err := os.Create(conf.Out + no + ".png")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer file.Close()

		file.Write(buf.Bytes())

		count = count + 1

	}
}

package util

import (
	"flag"
	"github.com/golang/freetype"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")                   //设置分辨率
	fontFile = flag.String("fontfile", "./static/font/simhei.ttf", "filename of the ttf font") //设置字体
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
)

//图片合并
func ImageMerge(backgroundImgName, centerImgName, createImgName string, x, y float64) error {
	//背景图
	backgroundImg, err := ImageDecode(backgroundImgName)
	if err != nil {
		log.Panicln("ImageMerge ImageDecode error is " + err.Error())
		return err
	}
	backgroundBound := backgroundImg.Bounds()
	//x轴坐标总数
	backgroundX := backgroundBound.Size().X
	//y轴坐标总数
	backgroundY := backgroundBound.Size().Y

	//添加图
	centerImg, err := ImageDecode(centerImgName)
	if err != nil {
		log.Panicln("ImageMerge ImageDecode error is " + err.Error())
		return err
	}
	centerBound := centerImg.Bounds()
	//x轴坐标总数
	centerX := centerBound.Size().X
	//y轴坐标总数
	centerY := centerBound.Size().Y

	//坐标偏差，x轴y轴 计算
	newImgX := float64(backgroundX-centerX) / x
	newImgY := float64(backgroundY-centerY) / y
	offset := image.Pt(int(newImgX), int(newImgY))
	//x轴坐标总数
	img := image.NewRGBA(backgroundBound)
	draw.Draw(img, backgroundBound, backgroundImg, image.ZP, draw.Src)
	draw.Draw(img, centerBound.Add(offset), centerImg, image.ZP, draw.Over)

	//保存到新文件中
	err = ImageEncode(img, createImgName)
	if err != nil {
		log.Panicln("ImageMerge ImageEncode error is " + err.Error())
		return err
	}

	return nil
}

//图片解析
func ImageDecode(imgPath string) (image.Image, error) {
	//如果是windows 换成c:/1.jpg
	imgFile, err := os.Open(imgPath)
	if err != nil {
		log.Panicln("ImageDecode os.Open error is " + err.Error())
		return nil, err
	}
	defer imgFile.Close()

	var img image.Image
	imageType := path.Ext(imgPath)
	switch imageType {
	case ".png":
		img, err = png.Decode(imgFile)
		if err != nil {
			log.Panicln("ImageDecode png.Decode error is " + err.Error())
			return nil, err
		}
	case ".jpg":
		img, err = jpeg.Decode(imgFile)
		if err != nil {
			log.Panicln("ImageDecode jpeg.Decode error is " + err.Error())
			return nil, err
		}
	}

	return img, nil
}

//图片编码
func ImageEncode(imgInfo image.Image, imgPath string) error {
	//如果是windows 换成c:/1.jpg
	imgFile, err := os.Create(imgPath)
	if err != nil {
		log.Panicln("ImageEncode os.Create error is " + err.Error())
		return err
	}
	defer imgFile.Close()

	imageType := path.Ext(imgPath)
	switch imageType {
	case ".png":
		err := png.Encode(imgFile, imgInfo)
		if err != nil {
			log.Panicln("ImageEncode png.Encode error is " + err.Error())
			return err
		}
	case ".jpg":
		err := jpeg.Encode(imgFile, imgInfo, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			log.Panicln("ImageEncode jpeg.Encode error is " + err.Error())
			return err
		}
	}

	return nil
}

//图片调整大小
func ImageResize(qrCodeName, qrCodeCompressName string, width, height uint) error {
	newImg, err := ImageDecode(qrCodeName)
	if err != nil {
		log.Panicln("ImageResize ImageDecode error is " + err.Error())
		return err
	}

	compressImg := resize.Resize(width, height, newImg, resize.Lanczos3)
	if err != nil {
		log.Panicln("ImageResize resize.Resize error is " + err.Error())
		return err
	}

	err = ImageEncode(compressImg, qrCodeCompressName)
	if err != nil {
		log.Panicln("ImageResize ImageEncode error is " + err.Error())
		return err
	}

	return nil
}

type FontStruct struct {
	Content   string     `json:"content"`
	Body      []string   `json:"body"`
	FontSize  float64    `json:"font_size"`
	FontX     int        `json:"font_x"`
	FontY     int        `json:"font_y"`
	FontColor color.RGBA `json:"font_color"`
}

//图片渲染字体
func ImageFontRender(centerImgName string, param *FontStruct) error {
	//需要加水印的图片
	imgFile, err := ImageDecode(centerImgName)
	if err != nil {
		log.Panicln("ImageFontRender ImageDecode error is " + err.Error())
		return err
	}

	img := image.NewRGBA(imgFile.Bounds())

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, imgFile.At(x, y))
		}
	}

	//读取字体数据
	fontBytes, err := ioutil.ReadFile(*fontFile)
	if err != nil {
		log.Panicln("ImageFontRender ioutil.ReadFile error is " + err.Error())
		return err
	}
	//载入字体数据
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Panicln("ImageFontRender freetype.ParseFont error is " + err.Error())
		return err
	}

	f := freetype.NewContext()
	f.SetDPI(*dpi)                //设置分辨率
	f.SetFont(font)               //设置字体
	f.SetFontSize(param.FontSize) //设置尺寸
	f.SetClip(imgFile.Bounds())
	f.SetDst(img)                               //设置输出的图片
	f.SetSrc(image.NewUniform(param.FontColor)) //设置字体颜色

	//设置字体的位置
	pt := freetype.Pt(param.FontX, param.FontY+int(f.PointToFixed(param.FontSize)>>6))
	if param.Content != "" {
		_, err = f.DrawString(param.Content, pt)
		if err != nil {
			log.Panicln("ImageFontRender f.DrawString error is " + err.Error())
			return err
		}
		pt.Y += f.PointToFixed(param.FontSize * *spacing)
	}
	if param.Body != nil {
		for _, val := range param.Body {
			_, err = f.DrawString(val, pt)
			if err != nil {
				log.Panicln("ImageFontRender f.DrawString error is " + err.Error())
				return err
			}
			pt.Y += f.PointToFixed(param.FontSize * *spacing)
		}
	}

	//保存到新文件中
	err = ImageEncode(img, centerImgName)
	if err != nil {
		log.Panicln("ImageFontRender ImageEncode error is " + err.Error())
		return err
	}

	return nil
}

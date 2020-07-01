package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"time"
)

func main()  {
	imgfile, _ := os.Open("123.jpg")
	defer imgfile.Close()
	jpgimg, _ := jpeg.Decode(imgfile)
	img := image.NewNRGBA(jpgimg.Bounds())

	xMiddle := img.Bounds().Dx()/2
	yMiddle := img.Bounds().Dy()/2
	xEnd := img.Bounds().Dx()
	yEnd := img.Bounds().Dy()

	//1  上下翻转
	for x:= 0 ; x<xMiddle;x++{
		for y:= 0;y<yMiddle;y++{
			img.Set(x,yMiddle-y-1,jpgimg.At(x,y))
		}
	}

	//2  左右翻转
	for x:=xMiddle;x<xEnd;x++ {
		for y:=0;y<yMiddle;y++ {
			img.Set(xEnd - (x-xMiddle)-1,y,jpgimg.At(x,y))
		}
	}

	//3 左右翻转
	for x:=0;x<xMiddle ;x++ {
		for y:=yMiddle;y<yEnd ;y++  {
			img.Set(xMiddle-x-1,y,jpgimg.At(x,y))
		}
	}

	//4上下翻转
	for x := xMiddle;x<xEnd ;x++  {
		for y:=yMiddle;y<yEnd;y++ {
			img.Set(xEnd - (x-xMiddle)-1,yEnd-(y-yMiddle)-1,jpgimg.At(x,y))
		}
	}


	outImg := image.NewNRGBA(img.Bounds())
	for x:=0;x<img.Bounds().Dx() ;x++  {
		for y:=0;y<img.Bounds().Dy() ;y++  {
			outImg.Set(x,y,img.At(x,y))
		}
	}

	const size = 40
	xCount := outImg.Bounds().Dx()/size
	yCount := outImg.Bounds().Dy()/size

	for xc := 0;xc<xCount;xc++{
		xStart := xc * size
		xE := xStart + size
		for yc:=0;yc<yCount;yc++  {
			yStart := yc * size
			yE := yStart+size
			//原图 0、4倍数列，第一行
			if yc % 4 == 0 && xc % 2 == 0 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++ {
						outImg.Set(x,y,img.At(x,y))
					}
				}
			}

			// 上下 0、4倍数列，第二行
			if yc % 4 == 0 && xc % 2 == 1 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++ {
						outImg.Set(x,yE-(y-yStart)-1,img.At(x,y))
					}
				}
			}

			//上下左右   5、9等对4取余的列，第一行
			if yc % 4 ==1 && xc % 2 ==0 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++  {
						outImg.Set(xE-(x-xStart)-1,yE-(y-yStart)-1,img.At(x,y))
					}
				}
			}

			//左右  1、5、9等对4取余的列，第二行
			if yc % 4 ==1 && xc % 2 == 1 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++  {
						outImg.Set(xE-(x-xStart)-1,y,img.At(x,y))
					}
				}
			}

			//上下  2、6、10等对4取余为2的列  第一行
			if yc % 4 ==2 && xc % 2 == 0 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++  {
						outImg.Set(x,yE - (y-yStart)-1,img.At(x,y))
					}
				}
			}

			//左右  2、6、10等对4取余为2的列  第二行
			if yc % 4 ==2 && xc % 2 == 1 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++  {
						outImg.Set(xE-(x-xStart)-1,y,img.At(x,y))
					}
				}
			}

			//上下左右 3、7、11等对4取余为3的列，第一行
			if yc % 4 == 3 && xc % 2 ==0 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++  {
						outImg.Set(xE-(x-xStart)-1,yE-(y-yStart)-1,img.At(x,y))
					}
				}
			}

			//左右 3、7、11等对4取余为3的列，第二行
			if yc % 4 == 3 && xc % 2 == 1 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++ {
						outImg.Set(xE-(x-xStart)-1,y,img.At(x,y))
					}
				}
			}
		}
	}


	outImg2 := image.NewNRGBA(img.Bounds())

	for x:=0;x<outImg2.Bounds().Dx() ;x++  {
		for y:=0;y<outImg2.Bounds().Dy() ;y++  {
			outImg2.Set(x,y,outImg.At(x,y))
		}
	}

	const size2 = 30
	xCount2 := outImg2.Bounds().Dx()/size2
	yCount2 := outImg2.Bounds().Dy()/size2

	for xc := 0;xc<xCount2;xc++{
		xStart := xc * size2
		xE := xStart + size2
		for yc:=0;yc<yCount2;yc++  {
			yStart := yc * size2
			yE := yStart+size2
			//原图 0、4倍数列，第一行
			if yc % 4 == 0 && xc % 2 == 0 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++ {
						outImg2.Set(x,y,outImg.At(x,y))
					}
				}
			}

			// 上下 0、4倍数列，第二行
			if yc % 4 == 0 && xc % 2 == 1 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++ {
						outImg2.Set(x,yE-(y-yStart)-1,outImg.At(x,y))
					}
				}
			}

			//上下左右   5、9等对4取余的列，第一行
			if yc % 4 ==1 && xc % 2 ==0 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++  {
						outImg2.Set(xE-(x-xStart)-1,yE-(y-yStart)-1,outImg.At(x,y))
					}
				}
			}

			//左右  1、5、9等对4取余的列，第二行
			if yc % 4 ==1 && xc % 2 == 1 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++  {
						outImg2.Set(xE-(x-xStart)-1,y,outImg.At(x,y))
					}
				}
			}

			//上下  2、6、10等对4取余为2的列  第一行
			if yc % 4 ==2 && xc % 2 == 0 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++  {
						outImg2.Set(x,yE - (y-yStart)-1,outImg.At(x,y))
					}
				}
			}

			//左右  2、6、10等对4取余为2的列  第二行
			if yc % 4 ==2 && xc % 2 == 1 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++  {
						outImg2.Set(xE-(x-xStart)-1,y,outImg.At(x,y))
					}
				}
			}

			//上下左右 3、7、11等对4取余为3的列，第一行
			if yc % 4 == 3 && xc % 2 ==0 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++  {
						outImg2.Set(xE-(x-xStart)-1,yE-(y-yStart)-1,outImg.At(x,y))
					}
				}
			}

			//左右 3、7、11等对4取余为3的列，第二行
			if yc % 4 == 3 && xc % 2 == 1 {
				for x:=xStart;x<xE ;x++ {
					for y:=yStart;y<yE ;y++ {
						outImg2.Set(xE-(x-xStart)-1,y,outImg.At(x,y))
					}
				}
			}
		}
	}

	timeString := strconv.FormatInt(time.Now().Unix(),10)
	newfile, _ := os.Create(timeString+".jpg")
	defer newfile.Close()
	err := jpeg.Encode(newfile, outImg2, &jpeg.Options{100})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timeString)

}

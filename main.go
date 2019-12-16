package main

import (
  "image"
    "image/color"
    "image/png"
  "log"
    "math"
    "os"
)

func main() {
 // 图片大小
 const size = 900

 // 根据size创建灰度图
 pic := image.NewGray(image.Rect(0,0,size,size))


 // 遍历每个像素
 for x := 0; x < size; x++ {
     for y := 0; y < size; y++ {
         pic.SetGray(x, y, color.Gray{255})
     }
 }

 // 从0到最大像素生成x坐标
 for x := 0; x < size; x++ {
     s := float64(x) * 2 * math.Pi/size
     y := size/2 - math.Sin(s)*size/2
     pic.SetGray(x, int(y), color.Gray{0})
 }

 // 文件操作
 file, err := os.Create("sin.png")
 if err != nil {
   log.Fatal(err)
 }

 // 灰度图写入文件
 png.Encode(file,pic)

 file.Close()
}
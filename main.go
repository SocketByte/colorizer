package main

import (
	"github.com/kataras/iris"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"mime/multipart"
)

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	web := iris.New()
	engine := iris.Django("./views", ".html").Binary(Asset, AssetNames)
	engine.Reload(true)

	web.RegisterView(engine)

	web.Get("/", func(ctx iris.Context) {
		ctx.ViewData("result", false)
		ctx.View("index.html")
	})

	web.Get("/uploads/{file:string}", func(ctx iris.Context) {
		ctx.ServeFile("./uploads/"+ctx.Params().Get("file"), false)
	})

	web.Post("/upload", func(ctx iris.Context) {
		ctx.UploadFormFiles("./uploads", func(ctx iris.Context, header *multipart.FileHeader) {
			file, err := header.Open()
			if err != nil {
				log.Println(err)
				return
			}
			defer file.Close()

			pixels, width, height := GetPixels(file)

			pixelDominationMap := make(map[Pixel]int)
			for x := 0; x < width; x++ {
				for y := 0; y < height; y++ {
					pixel := pixels[y][x]

					if val, ok := pixelDominationMap[pixel]; ok {
						pixelDominationMap[pixel] = val + 1
					} else {
						pixelDominationMap[pixel] = 0
					}
				}
			}
			pairsPtr := SortPixels(pixelDominationMap, 10)
			pairs := *pairsPtr

			ctx.ViewData("result", true)
			ctx.ViewData("pairs", pairs)
			ctx.ViewData("file", header.Filename)
			ctx.View("index.html")
		})
	})

	web.Run(iris.Addr(":8080"), iris.WithCharset("utf-8"))
}

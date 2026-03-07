---
title: Creating a Slack Emoji Auto-Generator Bot with Golang, chromedp, and Slack Bot
slug: golang-chromedp-slack-bot
date: 2020-08-11T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Slack
  - Chrome
  - chromedp
  - Emoji
  - Slack Bot
description: A weekend project to create a Slack bot that automatically generates emoji images using Golang, chromedp, and a headless browser.
translation_key: golang-chromedp-slack-bot
---

# Overview
I created a Slack bot that automatically generates emoji images using Golang, chromedp, and Slack Bot.

# What I Made
This bot generates an image when you mention it in Slack with specific parameters. Internally, it creates an image using a canvas based on the parameters, takes a screenshot with a headless browser, saves the image, and posts it to Slack.

[github.com - emoji-generator-slack-app](https://github.com/bmf-san/emoji-generator-slack-app)

You can get a rough idea of how to use it by checking the README...

Since I made this over a weekend hackathon-style, there are still some bugs left...
https://github.com/bmf-san/emoji-generator-slack-app/issues/1

# Image Generation in Golang
Golang has a standard package called `image` that provides a full range of image processing features.

You can easily perform tasks like applying mosaic effects, combining images, cropping, and drawing text (at least from what I’ve seen).

While it’s not very practical, for example, you can create a solid-colored image with just a few lines of code:

```go
package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	x, y := 0, 0
	width, height := 400, 400
	quality := 100

	img := image.NewRGBA(image.Rect(x, y, width, height))
	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, color.RGBA{255, 255, 255, 255})
		}
	}

	file, err := os.Create("sample.jpg")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	if err = jpeg.Encode(file, img, &jpeg.Options{quality}); err != nil {
		log.Println(err)
	}
}
```

If you want to make it more interesting by drawing text on the image, you can use the following code:

```go
package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func main() {
	baseFile, err := os.Open("./image/base.jpg")
	if err != nil {
		log.Println(err)
	}
	defer baseFile.Close()
	baseImage, _, err := image.Decode(baseFile)
	if err != nil {
		log.Println(err)
	}

	fontFile, err := ioutil.ReadFile("./font/Roboto-Regular.ttf")
	if err != nil {
		log.Println(err)
	}
	parsedFont, err := truetype.Parse(fontFile)
	if err != nil {
		log.Println(err)
	}

	r := baseImage.Bounds()
	rgbaImage := image.NewRGBA(image.Rect(0, 0, r.Dx(), r.Dy()))
	draw.Draw(rgbaImage, rgbaImage.Bounds(), baseImage, r.Min, draw.Src)
	drawer := font.Drawer{
		Dst: rgbaImage,
		Src: image.Black,
	}
	drawer.Face = truetype.NewFace(parsedFont, &truetype.Options{
		Size: 20,
		DPI:  350,
	})
	drawText := "Hello World"
	drawer.Dot = fixed.Point26_6{
		X: (fixed.I(r.Dx()) - drawer.MeasureString(drawText)) / 2,
		Y: fixed.I(r.Dy() / 2),
	}

	file, err := os.Create("sample_text.jpg")
	if err != nil {
		log.Println(err)
	}
	drawer.DrawString(drawText)
	if err = jpeg.Encode(file, drawer.Dst, &jpeg.Options{Quality: 100}); err != nil {
		log.Println(err)
	}
}
```

By carefully adjusting the drawing parameters as shown in the code above, you can generate custom images. If you want to create complex geometric patterns, you’ll need to adjust more parameters, which can be quite a challenge.

While it seems possible to create Slack emojis using the `image` package, I found adjusting the parameters to be tedious. So, I looked for a simpler solution and came across an article introducing image generation using a headless browser. I decided to use a headless browser for this project.

cf. [note.com - Dynamic Image Generation Using Headless Browser in Go](https://note.com/timakin/n/n55d483d11b22)

From the article above, I learned that the `image` package only supports the TrueType font format. While I wasn’t too concerned about this since I wasn’t aiming for a highly designed output, it’s something to keep in mind if you want to customize fonts for your service.

# Image Generation Using a Headless Browser
The idea is to use a headless browser to take a screenshot and generate an image.

Compared to using the `image` package, this method allows you to adjust the image on the frontend rather than the server-side. This provides greater design flexibility, such as adjusting the image with CSS or using any font supported by the browser.

Additionally, this method is highly versatile as it can also be used for web scraping or generating Open Graph Protocol (OGP) images.

To use a headless browser (Chrome) with Golang, I used the `chromedp` package.
[github.com - chromedp](https://github.com/chromedp/chromedp)

`chromedp` supports the Chrome DevTools Protocol, allowing you to control Chrome with or without a UI, without relying on external tools like Selenium or PhantomJS.
[chrome devtools protocol](https://chromedevtools.github.io/devtools-protocol/)

Here’s an example of how to take a screenshot using `chromedp` in headless mode:

```go
// Extracted from https://github.com/bmf-san/emoji-generator-slack-app
ctx, cancel := chromedp.NewContext(context.Background())
defer cancel()

var buf []byte
if err := chromedp.Run(ctx, chromedp.Tasks{
    chromedp.Navigate(`http://localhost:9999/generator?` + query.Encode()),
    chromedp.Sleep(2 * time.Second),
    chromedp.WaitVisible(`#target`, chromedp.ByID),
    chromedp.Screenshot(`#target`, &buf, chromedp.NodeVisible, chromedp.ByID),
}); err != nil {
    log.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Failed to take a screen shot."))
    return
}
```

After taking the screenshot with `chromedp`, you can save it to a file and post it to Slack to complete the bot’s workflow.

# Creating the Slack Emoji Auto-Generator Bot
I’ll skip the details of Slack bot development.

This article explains it well:
cf. [qiita.com - Creating a Slack Bot with Go (March 2020 Edition)](https://qiita.com/frozenbonito/items/cf75dadce12ef9a048e9)

For this project, I created a bot that responds to mentions with parameters, generates an image based on the parameters, and posts the image. I didn’t use features like dialogs or slash commands, only event subscriptions.

Using dialogs might improve the user experience, but it seemed complicated and there weren’t many examples available, so I implemented a bot that simply responds to mentions.

Here’s the final result before diving into the implementation:

![Screen Shot 2020-08-11 at 14 49 46](https://user-images.githubusercontent.com/13291041/89861979-f3a4a680-dbe1-11ea-8c93-7c118c89e813.png)
![Screen Shot 2020-08-11 at 14 49 40](https://user-images.githubusercontent.com/13291041/89861975-f1dae300-dbe1-11ea-8e59-10ef38800cce.png)

When you mention the bot like this, it generates a Slack emoji image (128px × 128px) with one or two lines of text and responds.

The bot takes the following input:
`@botname [color] [bgColor] [line1] [line2(optional)]`

The input data is used to generate the image. I used a canvas for the image generation. Initially, I wanted to achieve this with just CSS, but I couldn’t figure out how to take a screenshot without including extra margins. Using a canvas worked as expected, so I decided to go with it.

The input data is passed into a template file (tpl), which the canvas uses to generate the image:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head> 
<body onLoad="draw()">
<canvas id="target" height="128" width="128"></canvas>
<script>
function draw() {
    var element = document.getElementById("target");
    var context = element.getContext("2d");
    var maxWidth = element.width;

    element.style.background = {{.BgColor}};
    context.clearRect(0, 0, element.width, element.height);
    context.textAlign = "center";
    context.font = "bold 64px Arial";
    context.fillStyle = {{.Color}};
    context.fillText({{.Line1}}, element.width*0.5, 56, maxWidth);
    context.fillText({{.Line2}}, element.width*0.5, 115, maxWidth);
}
</script>
</body>
</html>
```

There’s no specific limit on the number of characters, so I adjusted the x-axis to fit the width nicely. For the y-axis, I found a suitable value and set it (I should probably think about how to calculate it properly, but I didn’t bother...).

Even though I ended up dealing with x and y axes using the canvas, which made me think I could have just used the `image` package, I still feel this method was relatively easier to implement.

Once the template for image generation is ready, I set up an API endpoint that generates images using query strings.
For example: http://localhost:9999/generator?color=red&bgColor=green&line1=foo&line2=bar

Finally, I wrote the code to respond to bot mentions, extract parameters from the mention, launch a headless browser with `chromedp`, hit the image generation endpoint, create the image, and post it to Slack.

Here’s an excerpt of the code:

```golang
ctx, cancel := chromedp.NewContext(context.Background())
defer cancel()

var buf []byte
// Take a screenshot
if err := chromedp.Run(ctx, chromedp.Tasks{
    chromedp.Navigate(`http://localhost:9999/generator?` + query.Encode()),
    chromedp.Sleep(2 * time.Second),
    chromedp.WaitVisible(`#target`, chromedp.ByID),
    chromedp.Screenshot(`#target`, &buf, chromedp.NodeVisible, chromedp.ByID),
}); err != nil {
    log.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Failed to take a screen shot."))
    return
}

// Write the image to a file
if err := ioutil.WriteFile("result.png", buf, 0644); err != nil {
    log.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Failed to take a screen shot."))
    return
}

// Post the image
r := bytes.NewReader(buf)
_, err = api.UploadFile(
    slack.FileUploadParameters{
        Reader:   r,
        Filename: "upload file name",
        Channels: []string{event.Channel},
    })
if err != nil {
    log.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Failed to post a image."))
    return
}
```

With this, I was able to create a bot that responds to mentions and generates images. However, unfortunately, there are still some bugs left...

https://github.com/bmf-san/emoji-generator-slack-app/issues/1

# Thoughts
I feel like I didn’t really need to use `chromedp`. Now I’m wondering how to fix the bugs...

# References
- [note.com - Dynamic Image Generation Using Headless Browser in Go](https://note.com/timakin/n/n55d483d11b22)
- [qiita.com - Creating a Slack Bot with Go (March 2020 Edition)](https://qiita.com/frozenbonito/items/cf75dadce12ef9a048e9)
- [qiita.com - Creating an Interactive Slack Bot with Go (May 2020 Edition)](https://qiita.com/frozenbonito/items/1df9bb685e6173160991#%E3%81%BE%E3%81%A8%E3%82%81)
- [dev.to - Creating a Bot That Responds with Images in Slack](https://dev.to/amotarao/slackbot-376)
- [lab.syncer.jp - How to Draw Multi-Line Text](https://lab.syncer.jp/Web/JavaScript/Canvas/8)
- [Stackoverflow - Size to fit font on a canvas](https://stackoverflow.com/questions/20551534/size-to-fit-font-on-a-canvas)

---
title: Creating a Slack Emoji Auto-Generation Bot with Golang, chromedp, and Slack
description: A step-by-step guide on Creating a Slack Emoji Auto-Generation Bot with Golang, chromedp, and Slack, with practical examples and configuration tips.
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
translation_key: golang-chromedp-slack-bot
---

# Overview
I created a Slack emoji auto-generation bot using Golang, chromedp, and Slack.

# What I Made
When you send a mention with parameters to the bot in Slack, it generates an image. Internally, it generates an image based on the parameters using a canvas, takes a screenshot with a headless browser, saves the image, and posts it to Slack.

[github.com - emoji-generator-slack-app](https://github.com/bmf-san/emoji-generator-slack-app)

You should be able to understand how to use it by looking at the README...

Since I made it in a hackathon-like spirit over the weekend, there are some bugs left...
https://github.com/bmf-san/emoji-generator-slack-app/issues/1

# Image Generation with Golang
Golang has a standard package called image that is rich in image processing features.

You can relatively easily apply mosaic processing, composite images, crop, and draw text. (At least, that's what I've seen.)

Although it has little practical use, if you want to generate a solid color image, you can achieve it with a few lines of code like this:

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
def	er file.Close()

	if err = jpeg.Encode(file, img, &jpeg.Options{quality}); err != nil {
		log.Println(err)
	}
}
```

A solid color image is boring, so if you want to draw text on an image, you can achieve it with the following code:

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
def	er baseFile.Close()
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

By skillfully adjusting the drawing parameters as shown in the code above, you can generate any image. If you want to create complex geometric patterns, the number of parameters to adjust will increase, and the calculations will likely become quite a challenge.

For creating emojis for Slack, it seems feasible to achieve it with the image package, but since adjusting parameters seemed cumbersome, I was looking for a more straightforward way to implement it. I came across an article introducing image generation using a headless browser, so I decided to implement it this way this time.

cf. [note.com - Dynamic Image Generation Using Headless Browser in Go](https://note.com/timakin/n/n55d483d11b22)

From the article, I learned that the font formats that the image package can draw are only supported in truetype.

This time, since I didn't want to focus on design, I didn't pay much attention to it, but if you want to adjust the font according to the service, you need to be careful.

# Image Generation Using a Headless Browser
In short, the method is to start a headless browser and take a screenshot to generate the image.

Compared to the method using the image package, you can adjust the image on the front end rather than the server side, allowing for high design flexibility by adjusting the image with CSS and freely using fonts supported by the browser.

Additionally, since scraping can also be done, it has high versatility. It seems to be a good match for automatic OGP generation.

For using a headless browser (Chrome) in Golang, I utilized the chromedp package.
[github.com - chromedp](https://github.com/chromedp/chromedp)

chromedp supports the Chrome DevTools Protocol, allowing you to operate Chrome with or without a UI without external dependencies like Selenium or PhantomJS.
[Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/)

The code to take a screenshot headlessly using chromedp can be written like this:

```go
// Extracted part of the code from https://github.com/bmf-san/emoji-generator-slack-app
ctx, cancel := chromedp.NewContext(context.Background())
def	er cancel()

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

After that, you just need to write the screenshot taken with chromedp to a file and post the data to Slack to complete the bot's job.

# Creating a Slack Emoji Auto-Generation Bot
I will skip the development of the Slack bot.

This article is easy to understand.
cf. [qiita.com - Creating a Slack Bot with Go (March 2020 Edition)](https://qiita.com/frozenbonito/items/cf75dadce12ef9a048e9)

This time, when a mention with parameters comes in, I created a bot that simply takes the parameters as input, generates the image, and posts the image.

I used only event subscription without using features like dialogs or slash commands.

I thought using a dialog would be good for UX, but it seemed quite cumbersome and there were few samples, so it would take time, so I implemented it as a bot that only responds to mentions.

Let me introduce the completed version before touching on the implementation.

![Screen Shot 2020-08-11 at 14 49 46](/assets/images/posts/golang-chromedp-slack-bot/89861979-f3a4a680-dbe1-11ea-8c93-7c118c89e813.png)
![Screen Shot 2020-08-11 at 14 49 40](/assets/images/posts/golang-chromedp-slack-bot/89861975-f1dae300-dbe1-11ea-8e59-10ef38800cce.png)

When you send a mention like this, the bot generates a Slack emoji image (128px × 128px) in one or two lines and responds.

The input that the bot receives is as follows:
`@botname [color] [bgColor] [line1] [line2(optional)]`

Based on this input data, the image is generated using a canvas. I actually wanted to do it nicely with just CSS without using a canvas, but since the screenshot included margins (I couldn't find a way to take a screenshot without margins like a selection area), I tried using a canvas, and it turned out as expected, so I decided to implement it with a canvas.

The input data is fed into a template file (tpl) to generate the image using the canvas.
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

Since there are no specific restrictions on the number of characters, I adjusted the x-axis to make it look good. For the y-axis, I found a nice value and set it (I think it would be better to consider how to calculate it, but it was cumbersome...).

In the end, since I was facing the x-axis and y-axis using the canvas, I sometimes thought about using the image package, but it seems that I was able to achieve it relatively easily.

Once the template for image generation was created, I prepared an endpoint as an API that generates images via query strings.
ex. http://localhost:9999/generator?color=red&bgColor=green&line1=foo&line2=bar

After that, I just needed to write code to respond to the bot's mention, read the parameters from the mention, start the headless browser with chromedp, hit the endpoint for image generation, create the image, and post the created image to Slack.

I will omit various details, but here is an excerpt of the code.
```golang
ctx, cancel := chromedp.NewContext(context.Background())
def	er cancel()

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

// Write the image
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

With this, I have created a bot that generates images in response to mentions, but unfortunately, there are still bugs left...

https://github.com/bmf-san/emoji-generator-slack-app/issues/1

# Thoughts
I feel like I didn't need to use chromedp. I wonder how I should fix the bugs...

# References
- [note.com - Dynamic Image Generation Using Headless Browser in Go](https://note.com/timakin/n/n55d483d11b22)
- [qiita.com - Creating a Slack Bot with Go (March 2020 Edition)](https://qiita.com/frozenbonito/items/cf75dadce12ef9a048e9)
- [qiita.com - Creating an Interactive Slack Bot with Go (May 2020 Edition)](https://qiita.com/frozenbonito/items/1df9bb685e6173160991#%E3%81%BE%E3%81%A8%E3%82%81)
- [dev.to - Created a bot that returns text sent on Slack as an image](https://dev.to/amotarao/slackbot-376)
- [lab.syncer.jp - How to Draw Multiple Lines of Text](https://lab.syncer.jp/Web/JavaScript/Canvas/8)
- [Stackoverflow - Size to fit font on a canvas](https://stackoverflow.com/questions/20551534/size-to-fit-font-on-a-canvas)
---
title: Golang×chromedp×slack botでslackの絵文字自動生成ボットをつくってみた
description: Golang×chromedp×slack botでslackの絵文字自動生成ボットをつくってみたについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: golang-chromedp-slack-bot
date: 2020-08-11T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - Slack
  - chrome
  - chromedp
  - emoji
  - slack-bot
translation_key: golang-chromedp-slack-bot
---


# 概要
Golang×chromedp×slack botでslackの絵文字自動生成ボットをつくってみた。

# 作ったもの
slackでbotにパラメータを付けたメンションを飛ばすと画像を生成してくれるだけのもの。
内部的には、パラメータを元にcanvasで画像を生成、ヘッドレスブラウザでスクショを撮って画像を保存、slackに投稿、といった感じ。

[github.com - emoji-generator-slack-app](https://github.com/bmf-san/emoji-generator-slack-app)

使い方等はREADMEを見ればなんとなくわかるはず...

週末にハッカソン的ノリでつくったため、バグが残ってしまっている...
https://github.com/bmf-san/emoji-generator-slack-app/issues/1

# Golangでの画像生成
Golangには画像処理の一通りの機能が充実しているimageという標準パッケージがある。

モザイク処理を施したり、画像を合成したり、トリミングをしたり、テキストを描画したりといったことが比較的に簡単にできる。（はず。色々みた限りでは。）

ほとんど実用性がないが、例えばベタ塗りされた画像を生成したいなら次のような数行のコードで実現できる。

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

ベタ塗りの画像ではつまらないので、画像にテキストを描画したければ、次のようなコードで実現できる。

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

上記のコードのように描画のパラメータを上手に調整することで任意の画像を生成することができる。
複雑な幾何学的な模様を作ってみたければ、調整すべきパラーメータの数は増えるし、計算するのも一苦労になるだろうと思う。

slack用のemojiをつくるくらいだったら、imageパッケージでも十分実現できそうな気はするが、パラメータの調整が面倒な気がしたので、もっとわかりやすい形で実現する方法を探っていたところ、ヘッドレスブラウザを用いた画像生成を紹介している記事を見かけたので、ヘッドレスブラウザを使った形で今回は実現してみることにした。

cf. [note.com - Goでheadless browserを用いた動的画像生成](https://note.com/timakin/n/n55d483d11b22)

上記の記事を見て知ったのだが、imageパッケージが描画できるfont形式はtruetypeしかサポートされていないらしい。

今回はデザインに凝りたいわけではないのであまり気に留めていないが、サービスに合わせてフォントを調整したい場合は注意が必要。

# ヘッドレスブラウザを用いた画像生成
要は、ヘッドレスブラウザを起動してスクショを撮って画像生成とする方法。

imageパッケージを使った方法と比べて、サーバーサイドではなくフロントエンドで画像の調整ができるため、CSSで画像を調整したり、ブラウザが対応するfontを自由に使用したりとデザインの柔軟性が高い。

あとはスクレイピングとかもできるので汎用性が高い。OGP自動生成とか相性が良さそう。

Golangでヘッドレスブラウザ（chrome）を使うためのパッケージとして、今回は、chromedpを利用した。
[github.com - chromedp](https://github.com/chromedp/chromedp)

chromedpはchromeを操作するためのプロトコルであるchrome devtools protocolをサポートしており、SeleniumやPhantomJSといった外部依存なしに、UI付きでもヘッドレスでもchromeを操作することができるパッケージ。
[chrome devtools protocol](https://chromedevtools.github.io/devtools-protocol/)

chromedpを使ってスクリーンショットをヘッドレスで撮影するコードはこんな感じに書くことできる。

```go
// https://github.com/bmf-san/emoji-generator-slack-appの一部のコードを抜粋
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

後はchromedpを使って撮影したスクショをファイルに書き込んで、データをslackに投稿すればbotの一連の仕事ができあがる。

# Slack絵文字自動生成botをつくる
slack botの開発については割愛する。

こちらの記事が分かりやすい。
cf. [qiita.com - Go で Slack Bot を作る (2020年3月版)](https://qiita.com/frozenbonito/items/cf75dadce12ef9a048e9)

今回はパラメータ付きのメンションが来たら、パラメータを入力として受け取って、画像生成、画像を投稿するだけのbotを作成した。

dialogやslash commandといった機能は使わず、event subscriptionだけ。

dialogを使うのがUX的に良さそうだと思ったが、中々に面倒かつサンプルも少ないので時間が掛かりそうだったのでメンションに反応するだけのbotという形で実装した。

先に完成形を紹介してから実装について触れる。

![Screen Shot 2020-08-11 at 14 49 46](/assets/images/posts/golang-chromedp-slack-bot/89861979-f3a4a680-dbe1-11ea-8c93-7c118c89e813.png)
![Screen Shot 2020-08-11 at 14 49 40](/assets/images/posts/golang-chromedp-slack-bot/89861975-f1dae300-dbe1-11ea-8e59-10ef38800cce.png)

こんな感じにメンションを飛ばすと、1行、もしくは2行の形のslack絵文字画像（128px×128px）を生成してレスポンスしてくれるだけのbot。

botが受け取る入力は以下の通り。
`@botname [color] [bgColor] [line1] [line2(optional)]`

この入力データを元に画像生成するのだが、画像の生成にはcanvasを使用した。
本当はcanvasを使わずcssだけで上手いことやりたかったのだが、スクショに余白が含まれてしまう（余白を含まずエリア選択のような形でスクショを取る方法がわからなかった）のでcanvasを使ってみたら期待通りの形になったのでcanvasで実装することにした。

入力データはテンプレートファイル(tpl)に流し込み、canvasが画像生成するようにする。
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

文字数は特に指定しないので横幅をいい感じになるようにx軸を調整。y軸についてはいい感じの値を見つけて設定した（どんな感じで計算したら良いか考えたほうが良いと思うが面倒だった...）。

結局canvasを使ってx軸y軸と向き合ってしまっているので、imageパッケージでも・・と思わないこともないが、比較的に楽に実現できたようには思う。

画像生成用のテンプレートができたら、クエリストリングで画像が生成されるAPIとしてエンドポイントを用意する。
ex. http://localhost:9999/generator?color=red&bgColor=green&line1=foo&line2=bar

後は、botのメンションに反応して、メンションからパラーメータを読み取り、chromedpでヘッドレスブラウザを起動、画像生成用のエンドポイントを叩いて画像作成、作成された画像をslackに投稿するコードを書くだけ。

諸々省略するが、コードを一部抜粋。
```golang
ctx, cancel := chromedp.NewContext(context.Background())
defer cancel()

var buf []byte
// スクショを取る
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

// 画像書き込み
if err := ioutil.WriteFile("result.png", buf, 0644); err != nil {
    log.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Failed to take a screen shot."))
    return
}

// 画像投稿 
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

これでメンションに反応して画像を生成してくれるbotができたわけだが、残念がバグが残ってしまっている。。。

https://github.com/bmf-san/emoji-generator-slack-app/issues/1

# 所感
chromedp使わなくても良かった気はする。バグはどうやって直したものか...

# 参考
- [note.com - Goでheadless browserを用いた動的画像生成](https://note.com/timakin/n/n55d483d11b22)
- [qiita.com - Go で Slack Bot を作る (2020年3月版)](https://qiita.com/frozenbonito/items/cf75dadce12ef9a048e9)
- [qiita.com - Go で Interactive な Slack Bot を作る (2020年5月版)](https://qiita.com/frozenbonito/items/1df9bb685e6173160991#%E3%81%BE%E3%81%A8%E3%82%81)
- [dev.to - Slackで送った文字を画像で返すbot作った](https://dev.to/amotarao/slackbot-376)
- [lab.syncer.jp - 複数行のテキストを描く方法](https://lab.syncer.jp/Web/JavaScript/Canvas/8)
- [Stackoverflow - Size to fit font on a canvas](https://stackoverflow.com/questions/20551534/size-to-fit-font-on-a-canvas)

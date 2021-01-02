# dm42-misaki-font

![Alt text](/github/screenshot-demo.png?raw=true "Screenshot Demo")

##  使い方

### mk42font
```bash
go run mk42font/mk42font.go mk42font/misaki_gothic_min.jpg > mk42font/misaki_gothic_min.hp42s
```
フォントデータをFOCAL言語に変換する。出力を電卓で実行すると、フォントデータを`fntDt`変数でロードされた。

### SKANA.hp42s
`ST X`番のグリフを`curX`x`curY`で表示する。`curX`を右に８ピクセルを前進する。

### SJISX.hp42s
`ST X`で`aabb`（a: jis X 0208 区コード,b: jis X 0208 点コード）のようなコードから変換して、`SKANA`を実行する。

### DEMO.hp42s
プログラム例。「こんにちは！」を表示する。
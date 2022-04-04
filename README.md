# ISO 8583 Message Challenge

[Go Conference Online 2022 Spring](https://gocon.jp/2022spring/) の開催を記念したプログラミングクイズです。Go を使って ISO 8583 メッセージのパースに挑戦してみましょう。

## 問題

カンムが提供する[バンドルカード](https://vandle.jp/)では、クレジットカードの決済処理に ISO 8583 の仕様に準拠したメッセージを利用しています。

このリポジトリには ISO 8583 メッセージのサンプルとなるバイナリファイル `message.bin` と、それをパースする Go のコードの雛形が含まれています。`parse.go` には ISO 8583 メッセージのバイナリを受け取って、メッセージの構造体を返す関数 `Parse()` が定義されています。この関数の処理を実装して、ISO 8583 メッセージのパースを成功させてください。次のようにコードを実行すると正しくパースできているかが確認できます。`VALID` と表示されたら正解です。

```
$ go run .
```

## ヒント

- ISO 8583 メッセージのパース方法は @hiroakis の builderscon tokyo 2019 での[発表資料](https://speakerdeck.com/hiroakis/kurezitutokadofalsetong-xin-purotokoru-iso8583-tozhan-u)を参照してください
- ISO 8583 メッセージのフォーマットは [Wikipedia](https://en.wikipedia.org/wiki/ISO_8583) でも参照できます

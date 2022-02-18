# Goのおもちゃ箱 〜プログラミング言語Go入門〜

## 解説資料

* TBA

## ハンズオンのやりかた

`skeleton`ディレクトリ以下に問題があり、8個のステップに分けられています。
STEP01からSTEP08までステップごとに進めていくことで、Goのパッケージ分けやテストの方法が学べます。

各ステップに、READMEが用意されていますので、まずは`README`を読みます。
`README`には、そのステップを理解するための解説が書かれています。

`README`を読んだら、ソースコードを開き`TODO`コメントが書かれている箇所をコメントに従って修正して行きます。
`TODO`コメントをすべて修正し終わったら、`README`に書かれた実行例に従ってプログラムをコンパイルして実行します。

途中でわからなくなった場合は、`solution`ディレクトリ以下に解答例を用意していますので、そちらをご覧ください。

`macOS`の動作結果をもとに解説しています。
`Windows`の方は、パスの区切り文字やコマンド等を適宜読み替えてください。

## ソースコードの取得

```
$ go env GOPATH
$ cd ↑のディレクトリに移動
$ mkdir -p src/github.com/gohandson/
$ cd src/github.com/gohandson
$ git clone https://github.com/gohandson/toybox-ja
$ cd toybox-ja
```

## ソースコードの編集

`skeleton`ディレクトリ以下のソースコードを編集する際にはセクションごとにブランチを作って作業するとよいでしょう。
以下の例は、Section 01を編集するための`fix-section01`ブランチを作成しています。

```
$ git checkout -b fix-section01
```

作業にひと区切りがついたら以下のように作業内容をコミットしてください。

```
$ git add 編集したファイル
$ git commit -m "変更の概要"
```

## ソースコードのアップデート

ハンズオン資料が更新された場合は以下のように更新してください。
なお、編集中のものがある場合はコミットしておきましょう。

```
$ git fetch -p
$ git merge origin/main 
```

アップデートの内容によっては編集中の内容とコンフリクトを起こす可能性があります。

## ライセンス

<a href="https://creativecommons.org/licenses/by-nc/4.0/legalcode.ja">
	<img width="200" src="by-nc.eu.png">
</a>

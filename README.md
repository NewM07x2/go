# go
dev go

## goコマンドのまとめ

### go run

binaryファイルを作成せずにgoファイルを実行する

```
go run main.go
```

WEBサーバーが立ち上がる。終了するときは「Ctrl+C」

```
go build
```

binaryファイルを作成する
```
go build main.go
```
main.exeが作成される
実行するとWEBサーバーが立ち上がる

```
go clean
```
go buildやgo.testで出力されたファイルを削除する

```
go clean
```

```
go get package
```
パッケージをインストールする
go get package名

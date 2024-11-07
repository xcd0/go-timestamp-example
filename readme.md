# golangでタイムスタンプを取得するサンプルコマンド

標準ライブラリを使用するとOS毎に対応が必要で大変。  
https://github.com/djherbis/times を使用する。  

## 実行例

下記はwslで実行。
同じファイルでもOSが異なると`log.Printf("%v", t)` で出力される小数の精度が異なる模様。

### windows

```sh
$ ./go-timestamp-example.exe main.go
13:25:15 main.go:42:
        atime: 2024-11-07 13:13:06.9354402 +0900 JST
        mtime: 2024-11-07 13:13:04.4458878 +0900 JST
        ctime: 2024-11-07 13:13:04.4658842 +0900 JST
        btime: 2024-11-07 13:13:04.4458878 +0900 JST
```

### linux

```sh
$ ./go-timestamp-example main.go
13:25:21 main.go:42:
    atime: 2024-11-07 13:13:06.935440127 +0900 JST
    mtime: 2024-11-07 13:13:04.445887755 +0900 JST
    ctime: 2024-11-07 13:13:04.46588416 +0900 JST
    btime: 2024-11-07 13:13:04.445887755 +0900 JST
```


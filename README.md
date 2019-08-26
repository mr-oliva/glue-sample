# Glue サンプル

## 概要
[How Uber "Go"es](https://speakerdeck.com/lelenanam/how-uber-go-es) の内容を参考につくったサンプルアプリケーション.

* uber-go/fx を用いた dependency injection
* Glueに載ったフレームワーク
* Monorepoで複数サービス運用

を考慮したサンプルとなっている

## 実行方法

```
$ cd $GOPATH/github.com/bookun/glue-sample
$ cd service/serviceA
$ go run main.go
```

## uber-go/fx
`di` ディレクトリに logger, configについて記述.    
fx.go には基本となるmodule群をまとめたものをExportしている

### config
uber-go/config を用いて .env から環境変数を取得

### logger
uber-go/zap を初期化

## Glue
各サービスは `services` 配下にディレクトリを作り作成    
各サービスは Glue に則り
* handler
* controller
* repository
* gateway

からなる

## エンドポイント
### Baseエンドポイント
default) http://localhost:8080

|Request Path|Contents|
|:---|:---:|
|/users| user一覧を表示|
|/user/[1-2]| user1を表示 |
|/ip|自身のグローバルIPを表示 |
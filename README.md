# practical-go-programming
『 [実用 Go言語](https://www.oreilly.co.jp/books/9784873119694/) 』のサポートリポジトリです。

![image](https://user-images.githubusercontent.com/14310228/165003929-7074bdab-ff44-4476-8367-f1bd09f310cb.png)



# 参考になるところまとめ
## ch1　Goらしさ
定数、変数
constant variable

iota
enumdef

error
error

opt引数
optionparams

メモリの最適化　スライス、マップ  deferの落とし穴
allocate

文字列の結合
stringjoin.go

日時の取り扱い
adddate
customtype


## ch2 定義型
メソッドを追加して組み込み型を拡張
extend

スライス、値、列挙、構造体の型定義
none

型の変換
cast


## ch3 構造体
基本
structsyntax

構造体のインスタンス化 3pattern
createstruct

method
method

embed
structembed

tag
structtag

構造体の設計のポイント
ポインタとして扱うケース
copyguard
ミュータブルな構造体、イミュータブルな構造体
imutable
ゼロ値
型をポインタにするかどうかで制御。

空の構造体でゴルーチン間の送受信
emptystrcut

構造体のメモリを割り当てを高速化
allocate2

構造体とオブジェクト指向の違い
oop
テンプレートパターン
ベースを作成し、変更する部分だけオーバーライドする
ストラテジーパターン
インターフェースを満たす形で実装し、置き換える


## ch4 インターフェース
基本
interfacesyntax
interfacebasic

any
emptyinterface

インターフェースのキャスト
スライスのキャストなど
interfacecast

インターフェースの合成

実装を切り替える方法
digsample


## ch5 エラーハンドリング
エラーの書き方いろいろ
ラップ、アンラップ
is, as で比較　文字列比較は×
スタックトレース
writing

エラーハンドリングのテクニック
呼び出し元に返す
ログ出力して処理を継続
リトライ
リソースをクローズ
複数のエラーをまとめる
practice

エラーのチェック忘れを予防
other


## ch6 パッケージ、モジュール
親パッケージの中の子パッケージはファイルを分ける


## ch7 環境整備
IDEの機能

Linter
linter

buildflag
バイナリサイズを小さくする
ビルド時のパスを削除
CGOを制御　不要なケースでは統一化のため無効にしておく
OSを指定
buildopt


CI
githuboptions
circleci
ci

makefile
makefile



## ch8 csv json excel
csv
csv

json
json


## ch9 RDB
基本
intro

トランザクション
transaction

コネクションプール
接続数などの設定

クエリーのキャンセル
context
cancel

ログ
ドライバーを使ってクエリーをロギング
ラッパーのドライバーを使ってクエリーをロギング
logging

大量データのバッチインサート
プリペアードステートメント
バッチインサート
DB組み込み関数
batchinsert

共通カラムをうまく使う
systemcolumns

DBのテスト
本物のDBを使う、モックを使う。
基本は本物を使うほうがいい
dtest

サードパーティ
thirdparty


## ch10 HTTPサーバー
HTTPサーバーの実装
Jsonの読み書き
バリデーション
intro

リクエストのパース
parse_query

ルーターの基本
router
router_chi

middleware
ステータスコードをキャプチャ
パニック防御
DBトランザクション
タイムアウト
middlewareの作り方

レートリミット　流動制限
ratelimit

SPA
URLの共有に対応する
spa


## ch11
HTTPクライアント

基本
intro

RoundTripで処理を分離
ロギング
認証認可
リトライ
roundtrip

リトライ時に考慮すべき点
retry

プロキシー
proxy


## ch12 ログとオブザーバビリティ
ログに出力すべき内容
どんな入力があって、どんな出力を得たか

構造化ログ
ミドルウェアでログ出力 zlog.go

エラーとログ出力
fatalとpanicは使わない。
コマンドラインツールでは、使われたりする。
使ってもいい場所もある。main関数など

net/httpのエラーログをカスタマイズする
errlog

分散システムの動作確認
分散トレース
otel

## ch13 テスト
基礎
basic_test

table driven test
cookbooktest

ヘルパー関数
helper

WebServerのハンドラのテスト
httptest
handlertest

testify
スライスと構造体の比較など
testify-test
go-cmp

テストが書きにくいいテスト
シンプルな入出力
    関数をio.Writerに対して実装するようにする
    入出力先を変更できるようにする
時刻のテスト

品質保証のテスト
ペアワイズ
カバレッジ
プロパティベーステスト pbd 入力バリエーションの網羅

benchmark
cookbooktest

exampletest
cookbooktest

## ch14 クラウド
DockerFileの構成
マルチステージビルド
goのimage
ECA lambda cloudrun


## ch15　クラウドストレージ
S3
DynamoDB


## ch16 並行処理
ゴルーチンプール
errgoroup
future/promise
webserviceのセッション
ゴルーチンリークを見つける



# 特に参考になる
エラー
テスト
ログ
ミドルウェア
CI


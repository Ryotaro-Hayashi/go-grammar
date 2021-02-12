# go-grammar

## 文法

Qiita記事①
https://qiita.com/k-penguin-sato/items/1d0e1c6b4bf937996cd3

- 頭文字を大文字にすると外部パッケージから参照できる（外部パッケージから参照するものは頭文字が大文字）
- 関数内ではvarによる変数の宣言を「:」で省略できる
- <要基本型の違い>
- 型変換を行うときは変数を括弧で囲んで手前に型名を書く
- 定数は関数内でも「:」で宣言できない

Qiita記事②
https://qiita.com/k-penguin-sato/items/31b45fd3914797b654f0

- for文の初期化ステートメントと後処理ステートメントは省略可
- if文の条件式の手前で, スコープ内のみで有効なステートメントを書ける
- 遅延処理deferに渡した関数は呼び出し元の関数の終わりまで遅延される
- 複数のdeferはスタックされる

Qiita記事③
https://qiita.com/k-penguin-sato/items/62dfe0f93f56e4bf9157
https://qiita.com/Sekky0905/items/447efa04a95e3fec217f
https://qiita.com/kotaonaga/items/4a93ec40718c279154f5

#### ポインタについて
- 初期値を与える場合は, 明示的に「*int」を付ける必要はない（変数で明示的に付ける必要がないのと同様）
- 文字列はシングルクオーテーションで定義できない（コード書いてて発見）
- ポインタにするときは「&変数」, ポインタ型変数の値を取り出すには「*ポインタ型変数」
- 関数の引数に値を渡すとコピーされる（渡す前のポインターと渡したあとの関数内でのポインターは異なる）
- 変数に代入するときもコピーされる（値渡し）
- intのポインタはPrintlnで普通に出力できるが, 構造体のポインタは「%p」などでフォーマット指定しないと「&構造体名」で出力される

#### 構造体について
- 構造体の初期化は①定義後にフィールドに代入する②（フィールド名を「:」で指定して）「{}」でフィールドの値を渡す③初期化関数を作成して初期化
- 構造体のフィールドは構造体のポインタからアクセスできる（「(*person).age」みたいな感じ）
- クラス⇨構造体とメソッド, 継承⇨構造体の埋め込み（User構造体の中にPerson構造体を入れるなど）

Qiita記事④
https://qiita.com/k-penguin-sato/items/daad9986d6c42bdcde90

[]byte型とString型の変換
https://dev.classmethod.jp/articles/struct-json/

- <配列とスライスは必要になった時に再度確認>
- []byte型をstringにキャストするとJSON
- 構造体⇆JSONの変換は構造体⇆[]byte型（プラス []byte型をstringに変換）
- []byte(<string型の変数>) は文字列1字1字がbyte型に変換されて1文字ずつスライスの要素に入っている

Qiita記事⑤
https://qiita.com/k-penguin-sato/items/a320072fa09502bde3e9

- <マップも必要になった時に再度確認>
- for文のrange節は, indexとindexの場所の要素の2つを返す
- forループを途中で抜けるにはbreak, そのループをスキップして次のループに進むにはcontinue

Qiita記事⑥
https://qiita.com/k-penguin-sato/items/885a61d819cc431304f5

#### インターフェースの使い方
記事：初心者に送りたいinterfaceの使い方[Golang]

- 同じようなメソッド（名前からメアドを取得するメソッド）を持つ2つの構造体（生徒, 先生）があって, その構造体を引数に持つ関数（メールを送信する関数）を定義したいとき, そのままでは引数の構造体のみが異なる同じような関数を2つ作らなくてはいけなくなる. そういう場合に, interfaceを定義して構造体を抽象化し, 関数にinterfaceを渡すことで1つの関数定義で済む
- 構造体の初期化関数で返り値をinterfaceにしてモック化する場合にも使う
- interfaceは頭文字大文字でinterfaceを満たすstructは頭文字小文字だと良いかも（interfaceは他のパッケージからも参照するし区別にもなる）

Qiita記事⑦
https://qiita.com/k-penguin-sato/items/5b09fa89d8d231bcdac8

- <並行処理は必要になった時に再度確認>

## テスト

- テストファイルの名前の末尾は"~_test.go"
- 関数名の末尾は"~Test.go"
- 必要に応じて引数に「*testing.T」型の変数を渡す（エラーハンドリングのためなど）

## パッケージ

- init() はmain()が呼ばれる前に実行される

## エイリアス

- 「import . "fmt"」でインポートするとパッケージ名を省略して関数を呼び出せる（「Println("fmt.は必要なくなる")」）
- エイリアスを付けてインポートすると, エイリアスでのアクセスが可能になる
- インポートしたパッケージの関数は使わないけど、関数内のinit関数を呼び出す必要があるとき（依存関係などによって）、アンダーバーを付けてインポートする

## レイヤードアーキテクチャ

- UI層（interfaces）, Infrastructure層(infrastructure), Application層(usecase), Domain層(domain)
- リクエストがinterfacesに入り, handlerでリクエストやレスポンスの処理をしたり, エラーハンドリングを行う. リクエストの処理ではusecaseを呼び出す. usecaseはdomainのrepositoryやserviceに依存するが, domainにはビジネスロジックのみを記述するので, interfaceを定義し, DIPで実装はinfrastructureに書く.

## byte型とは

## []byte型の振る舞い
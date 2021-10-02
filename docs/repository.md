# リポジトリ

## リポジトリとは

リポジトリはデータの保管庫です。データの永続化と再構築を直接行うのではなく、リポジトリを経由して行います。

## リポジトリの責務

リポジトリの責務はドメインオブジェクトの永続化や再構築を行うことです。永続化するデータストアが RDB か NoSQL かファイルなのかドメインにとっては重要ではありません。

## リポジトリのインターフェース

リポジトリはインターフェースで定義します。リポジトリの責務はあくまでもオブジェクトを永続化することです。

```go
type Repository interface {
	SaveUser(user *model.User) error
	FindUserByName(name string) (*model.User, error)
}
```

## SQL を利用したリポジトリを作成する

インターフェースをうまく活用することで、クラス上には具体的な永続化にまつわる処理を記述せずにデータストアにインスタンスを永続化できるようになります。

## テスト用のリポジトリを作成する

## リポジトリに定義されるふるまい

永続化のふるまいは永続化を行うオブジェクトを引数にとります。

**良い例**

```go
type Repository interface {
	SaveUser(user *model.User) error
	DeleteUser(user *model.User) error
}
```

**悪い例**

```go
type Repository interface {
	UpdateUserByName(id, name string) error
	UpdateUserByEmail(id, email string) error
	UpdateUserByAddress(id, address string) error
}
```

## まとめ

- リポジトリを利用するとデータの永続化にまつわる処理が抽象化できる
- ドメインのルールに比べると、データストアが何であるかは些末な問題である

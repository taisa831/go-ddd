# ファクトリ

## ファクトリとは

## 採番処理をファクトリに実装した例

可能であればテスト用に気軽にインスタンスを生成したいときは適当な ID を振り、そうでない場合はデータベース接続して採番を行えるようにする

```go
type UserFactory interface {
	Create(name string) User
}
```

```go
type UserFactory struct {
}

func Create(name string) User {

}
```

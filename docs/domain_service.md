# ドメインサービス

## ドメインサービスとは

システムには値オブジェクトやエンティティに記述すると不自然になってしまうふるまいが存在します。ドメインサービスはそういった不自然さを解決するオブジェクトです。

例えば、ユーザの重複確認をする処理をする場合、生成したオブジェクト自身に問い合わせをすることになるので、ユーザ自身に重複するかを確認するのは不自然な振る舞いとなります。

```go
package model

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID   string
	Name string
}

func NewUser(name string) (*User, error) {
	u := &User{
		ID: uuid.NewString(),
	}
	err := u.ChangeName(name)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func Exists(user *User) bool {
	//重複を確認するコード
	// ...
	return false
}
```

## 不自然さを解決する

この不自然さを解決するのがドメインサービスです。

```go
package service

import (
	"github.com/taisa831/go-ddd/domain/model"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
)

type UserService struct {
	r repository.Repository
}

func NewUserService(r repository.Repository) service.UserService {
	return &UserService{
		r: r,
	}
}

func (s *UserService) Exists(user *model.User) (bool, error) {
	_, err := s.r.FindUserByID(user.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
```

## 可能な限りドメインサービスを避ける

ドメインモデルの処理もドメインサービスに書けてしまいますが、無機質な Getter と Setter だけを持った、何も語らないドメインモデルとなるため（ドメインモデル貧血症）、可能な限りドメインサービスの利用は避けるようにします。

> ドメインモデル貧血症
> ドメインオブジェクトに本来記述されるべき知識やふるまいが、ドメインサービスやアプリケーションサービスに記述され、語るべきことを何も語っていないドメインオブジェクトの状態をドメインモデル貧血症といいます。

## エンティティや値オブジェクトと共にユースケースを組み立てる

実際に SQLite3 を利用してクリーンアーキテクチャでドメインモデル・ドメインサービス・ユースケースを構成してみました。

```
tree
.
├── application
│   └── usecase
│       └── user_usecase.go
├── docker-compose.yml
├── domain
│   ├── model
│   │   └── user.go
│   ├── repository
│   │   └── repository.go
│   └── service
│       └── user_service.go
├── go.mod
├── go.sum
├── gorm.db
├── infrastructure
│   ├── repository
│   │   ├── repository.go
│   │   └── user_repository.go
│   └── service
│       └── user_service.go
├── interfaces
│   ├── handler
│   │   └── user_handler.go
│   └── response
│       └── user_response.go
└── main.go
```

## ユーザー作成処理

### ドメインモデル

```go
package model

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID   string
	Name string
}

func NewUser(name string) (*User, error) {
	u := &User{
		ID: uuid.NewString(),
	}
	err := u.ChangeName(name)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (m *User) ChangeName(name string) error {
	if name == "" {
		return fmt.Errorf("ユーザ名は必須です。")
	}
	if len(name) < 3 {
		return fmt.Errorf("ユーザ名は3文字以上です。%s", name)
	}
	m.Name = name
	return nil
}
```

### ドメインサービス

```go
package service

import "github.com/taisa831/go-ddd/domain/model"

type UserService interface {
	Exists(user *model.User) (bool, error)
}
```

```go
package service

import (
	"github.com/taisa831/go-ddd/domain/model"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
)

type UserService struct {
	r repository.Repository
}

func NewUserService(r repository.Repository) service.UserService {
	return &UserService{
		r: r,
	}
}

func (s *UserService) Exists(user *model.User) (bool, error) {
	_, err := s.r.FindUserByID(user.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
```

### リポジトリ

```go
package repository

import "github.com/taisa831/go-ddd/domain/model"

type Repository interface {
	FindUsers() ([]*model.User, error)
	FindUserByID(id string) (*model.User, error)
}
```

```go
package repository

import (
	"github.com/taisa831/go-ddd/domain/repository"
	"gorm.io/gorm"
)

type rdbRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.Repository {
	return &rdbRepository{
		db: db,
	}
}
```

```go
package repository

import (
	"github.com/taisa831/go-ddd/domain/model"
)

func (r *rdbRepository) FindUsers() ([]*model.User, error) {
	var users []*model.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *rdbRepository) FindUserByID(id string) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
```

### ユースケース

```go
package usecase

import (
	"fmt"

	"github.com/taisa831/go-ddd/application/input"
	"github.com/taisa831/go-ddd/domain/model"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
)

type UserUsecase struct {
	r  repository.Repository
	us service.UserService
}

func NewUserUsecase(r repository.Repository, us service.UserService) *UserUsecase {
	return &UserUsecase{
		r:  r,
		us: us,
	}
}
func (u *UserUsecase) Primitive() {
	fullName := "taro suzuki"
	fmt.Println(fullName)
}

func (u *UserUsecase) Create(in input.UserCreateInput) error {
	b, err := u.us.Exists(in.Name)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("%s は存在します。", in.Name)
	}

	conf := model.UserCreateConfig{
		Name: in.Name,
	}
	user, err := model.NewUser(conf)
	if err != nil {
		return err
	}
	if err := u.r.CreateUser(user); err != nil {
		return err
	}

	return nil
}
```

### ハンドラー

```go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taisa831/go-ddd/application/usecase"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
	"github.com/taisa831/go-ddd/interfaces/response"
)

type UserHandler struct {
	u *usecase.UserUsecase
}

func NewUserHandler(r repository.Repository, us service.UserService) UserHandler {
	return UserHandler{
		u: usecase.NewUserUsecase(r, us),
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	err := h.u.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}
```

### ルーティング

```go
package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taisa831/go-ddd/infrastracture/repository"
	"github.com/taisa831/go-ddd/infrastracture/service"
	"github.com/taisa831/go-ddd/interfaces/handler"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	router := gin.Default()

	r := repository.NewRepository(db)
	us := service.NewUserService(r)

	uh := handler.NewUserHandler(r, us)
	router.POST("/users", uh.Create)
	router.Run()
}
```

## まとめ

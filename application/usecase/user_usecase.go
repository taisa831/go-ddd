package usecase

import (
	"fmt"

	"github.com/taisa831/go-ddd/domain/model"
	"github.com/taisa831/go-ddd/domain/repository"
	"github.com/taisa831/go-ddd/domain/service"
	"github.com/taisa831/go-ddd/interfaces/request"
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

func (u *UserUsecase) Create(req request.UserCreateRequest) error {
	b, err := u.us.Exists(req.Name)
	if err != nil {
		return err
	}
	if b {
		return &model.UserExistsError{}
	}

	conf := model.UserCreateConfig{
		Name: req.Name,
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

func (u *UserUsecase) List() ([]*model.User, error) {
	return u.r.FindUsers()
}

func (u *UserUsecase) GetUser() error {
	fullName, err := model.NewFullName("taro", "suzuki")
	if err != nil {
		return err
	}
	fmt.Println(fullName.FirstName())
	// 代入によって交換する
	fullName, err = model.NewFullName("taro", "sato")
	if err != nil {
		return err
	}
	fmt.Println(fullName.FirstName())
	return nil
}

// func (u *UserUsecase) Compare() (bool, error) {
// 	fullNameA, err := model.NewFullName("taro", "suzuki")
// 	if err != nil {
// 		return false, err
// 	}

// 	fullNameB, err := model.NewFullName("taro", "suzuki")
// 	if err != nil {
// 		return false, err
// 	}

// 	ret := fullNameA == fullNameB
// 	fmt.Println(ret)

// 	userA, err := model.NewUser("taro")
// 	if err != nil {
// 		return false, err
// 	}

// 	userB, err := model.NewUser("taro")
// 	if err != nil {
// 		return false, err
// 	}
// 	return userA == userB, nil
// }

func (u *UserUsecase) Money() error {
	myMoney := model.NewMoney(1000, "JPY")
	allowance := model.NewMoney(3000, "JPY")
	result, err := myMoney.Add(*allowance)
	if err != nil {
		return err
	}
	fmt.Println(result.Amount())

	jpy := model.NewMoney(1000, "JPY")
	usd := model.NewMoney(10, "USD")
	result2, err := jpy.Add(*usd)
	if err != nil {
		return err
	}
	fmt.Println(result2.Amount())
	return nil
}

func (u *UserUsecase) MoneyPrimitive() error {
	myMoney := 1000
	allowance := 3000
	result := myMoney + allowance
	fmt.Println(result)
	return nil
}

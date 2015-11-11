package Env

import "github.com/FanszHub/test-site/Data"

type Env struct {
	BlahRepository Data.BlahRepository `inject:""`
	UserRepository Data.UserRepository `inject:""`
}
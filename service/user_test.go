package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
)

func TestUser(t *testing.T) {
	dto.InitConfig()
	db.Init()
	followModel := &model.User{
		ID:              12,
		Name:            "nyf",
		Password:        "1234567",
		FollowCount:     12,
		FollowerCount:   12,
		BackgroundImage: "http://tse2-mm.cn.bing.net/th/id/OIP-C.sDoybxmH4DIpv033-wQEPgHaEq?pid=ImgDetars=1",
	}
	model.CreateUser(context.Background(), followModel)
	//res, total, _ := model.QueryFollow(context.Background(), "223", 1, 1, 10)
	//fmt.Println(res, total)
	//re, _ := model.QueryUserById(context.Background(), "08dc2b99ef974d47a2554ed3dea73ea0")
	//for index, value := range re {
	//	fmt.Println("index=", index, "value=", value)
	//}
	//fmt.Println(*re[0])
	//userModel := &model.User{
	//	Name:     "nyf123456",
	//	PassWord: "12345678",
	//}
	//model.CreateUser(context.Background(), userModel)
	res, m := model.QueryUserByName(context.Background(), "nyf")
	fmt.Println(res)
	fmt.Println(m)
}

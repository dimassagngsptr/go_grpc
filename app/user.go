package app

import (
	"fmt"
	"go_grpc/model"
)

func (s *service) GET_INFO_USER(params string) (resp model.User, err error) {
	fmt.Println("params", params)
	resp, err = s.Store().GetInfoUser(params)
	if err != nil {
		fmt.Printf("Error getting info user: %v", err.Error())
		//Rollback database
		return
	}
	fmt.Println("resp", resp)

	//Commit database
	return
}

package app

import "go_grpc/model"

//set interface port for APP
type ReadApp interface {
	GET_INFO_USER(params string) (resp model.User, err error)
}

type CreateApp interface {
}

type UpdateApp interface{}

type DeleteApp interface{}

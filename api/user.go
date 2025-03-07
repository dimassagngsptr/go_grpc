package api

import (
	"context"
	"fmt"
	"go_grpc/model"
	pb "go_grpc/proto"

	"github.com/jinzhu/copier"
)

func (g *go_grpcAPI) GET_INFO_USER(ctx context.Context, request *pb.ExampleGetRequest) (response *pb.GET_USER_RESULT, err error) {

	var resApp model.User
	response = &pb.GET_USER_RESULT{}
	data := &pb.ExampleGetResult{}
	resApp, err = g.app.Read().GET_INFO_USER(request.FullName)
	if err != nil {
		fmt.Printf("Error getting info user app: %v", err.Error())
		return
	}
	fmt.Println("response", resApp)
	//copy response from app into proto
	err = copier.Copy(&data, resApp)
	if err != nil {
		fmt.Printf("Error copying info user app: %v", err.Error())
		return
	}
	response.ResponseCode = "200"
	response.ResponseMsg = "Berhasil"
	response.ResponseData = data
	return
}

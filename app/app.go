package app

import (
	"go_grpc/model"
)

type AppSvc interface {
	Create() CreateApp
	Read() ReadApp
	Update() UpdateApp
	Delete() DeleteApp
	Store() model.ModelsPorts
}

var _ AppSvc = (*service)(nil)

type service struct {
	createSvc CreateApp
	readSvc   ReadApp
	updateSvc UpdateApp
	deleteSvc DeleteApp
	storeSvc  model.ModelsPorts
}

func (s *service) Create() CreateApp {
	return s.createSvc
}

func (s *service) Read() ReadApp {
	return s.readSvc
}

func (s *service) Update() UpdateApp {
	return s.updateSvc
}

func (s *service) Delete() DeleteApp {
	return s.deleteSvc
}

func (s *service) Store() model.ModelsPorts {
	return s.storeSvc
}

func AppService(model model.ModelsPorts) AppSvc {
	var svc = &service{}
	svc.storeSvc = model
	svc.readSvc = svc
	svc.updateSvc = svc
	svc.createSvc = svc
	svc.deleteSvc = svc
	return svc
}

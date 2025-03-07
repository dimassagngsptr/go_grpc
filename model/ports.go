package model

type ModelsPorts interface{
	//set port interface function for manage CRUD 
	GetInfoUser(fullName string) (result User, err error)
}
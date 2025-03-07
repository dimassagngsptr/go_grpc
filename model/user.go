package model


func (d *Database) GetInfoUser(fullName string) (result User, err error){
	//process get into database
	user := User{
		FullName: "Dimas Ageng Saputro",
		Age: 24,
	}

	result = user
	return
}
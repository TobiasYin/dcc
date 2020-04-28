package sql


func createUser(name string, pass string) (*User, error) {
	u := User{
		Name: name,
		Password: pass,
	}
	db, err := GetDBConnect()
	defer PutDBConnect(db)
	if err != nil {
		return nil, err
	}
	db.Create(&u)
	return &u, nil
}

func saveUser(u *User) error {
	db, err := GetDBConnect()
	defer PutDBConnect(db)
	if err != nil {
		return err
	}
	db.Save(&u)
	return nil
}

func getUser(name string) (*User, error) {
	db, err := GetDBConnect()
	defer PutDBConnect(db)
	if err != nil {
		return nil, err
	}
	var u User
	err = db.Where("name = ?", name).First(&u).Error
	return &u, err
}

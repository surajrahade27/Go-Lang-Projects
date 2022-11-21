package datastore

import (
	"ASSIGNMENT/model"
	"log"
)

func InsertDB(usr *model.User) (string, error) {
	m := DBCONN()
	tx, err := m.Begin()
	if err != err {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT into users(UserId,FirstName,LastName,Email,Phone,Status,Password,LastUpdatedAt) value(?,?,?,?,?,?,?,?)")
	if err != err {
		log.Fatal(err)
	}
	_, err = stmt.Exec(usr.UserID, usr.FirstName, usr.LastName, usr.Email, usr.Phone, usr.Status, usr.Password, usr.LastUpdatedAt)
	if err != err {
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != err {
		log.Fatal(err)
	}

	return "User Created Successfully", nil

}

func GetDBUsers(usr *model.User) ([]model.User, error) {
	m := DBCONN()
	GetUserResult := make([]model.User, 0)
	stmt, err := m.Prepare("SELECT UserId,FirstName,LastName,Email,Phone,Status,Password,LastUpdatedAt FROM USERS where UserID=?")
	if err != err {
		log.Fatal(err)
	}
	row, err := stmt.Query(usr.UserID)
	if err != err {
		log.Fatal(err)
	}
	for row.Next() {
		u := model.User{}
		err := row.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Email, &u.Phone, &u.Status, &u.Password, &u.LastUpdatedAt)
		if err != err {
			log.Fatal(err)
		}
		GetUserResult = append(GetUserResult, u)
	}
	return GetUserResult, nil
}
func GetAllDBUsers(usr *model.User) ([]model.User, error) {
	m := DBCONN()
	GetUserResult := make([]model.User, 0)
	stmt, err := m.Prepare("SELECT UserId,FirstName,LastName,Email,Phone,Status,Password,LastUpdatedAt FROM USERS")
	if err != err {
		log.Fatal(err)
	}
	row, err := stmt.Query()
	if err != err {
		log.Fatal(err)
	}
	for row.Next() {
		u := model.User{}
		err := row.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Email, &u.Phone, &u.Status, &u.Password, &u.LastUpdatedAt)
		if err != err {
			log.Fatal(err)
		}
		GetUserResult = append(GetUserResult, u)
	}
	return GetUserResult, nil
}

func UpdateDB(usr *model.User) error {
	m := DBCONN()

	tx, err := m.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("UPDATE users SET FirstName=?,LastName=?,Email=?,Phone=?,Status=?,Password=?,LastUpdatedAt=? WHERE UserID=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(usr.FirstName, usr.LastName, usr.UserID, usr.Email, usr.Phone, usr.Status, usr.Password, usr.LastUpdatedAt)
	if err != nil {
		log.Fatal()
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func DeleteDBUsers(usr *model.User) error {
	m := DBCONN()

	tx, err := m.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("DELETE from USERS WHERE UserID=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(usr.UserID)
	if err != nil {
		log.Fatal()
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func VerifyUser(usr *model.User) (string, error) {
	m := DBCONN()
	stmt, err := m.Prepare("SELECT UserId,FirstName,LastName,Email,Phone,Status,Password,LastUpdatedAt FROM USERS where Email=?")
	if err != err {
		log.Fatal(err)
	}
	row, err := stmt.Query(usr.Email)
	if err != err {
		log.Fatal(err)
	}
	for row.Next() {
		u := model.User{}
		err := row.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Email, &u.Phone, &u.Status, &u.Password, &u.LastUpdatedAt)
		if err != err {
			log.Fatal(err)
		}

		if usr.Password == u.Password {
			return "Login Successfull !", nil
		}

		if usr.Password != u.Password {
			return "Login Failed !", nil
		}
	}

	return " ", nil
}

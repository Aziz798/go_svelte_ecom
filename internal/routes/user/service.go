package user

import (
	"database/sql"
	"fmt"
	"go_ecom/internal/models"
	"go_ecom/internal/utils"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User, db *sql.DB) (string, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	res, err := db.Exec("INSERT INTO Users (Username,Email,Password,Role) VALUES(?,?,?,?);", user.Username, user.Email, hashedPassword, user.Role)
	if err != nil {
		return string(""), err
	}
	userId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	jwtToken, err := utils.CreateToken(userId, user.Role)
	return jwtToken, err
}

func GetUserById(id int, db *sql.DB) (models.User, error) {
	var user models.User
	dbRes := db.QueryRow("SELECT * FROM Users WHERE ID =?;", id)
	err := dbRes.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string, db *sql.DB) (models.User, error) {
	var user models.User
	dbRes := db.QueryRow("SELECT * FROM Users WHERE Email =?;", email)
	err := dbRes.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return user, err
	}
	return user, nil
}

func LoginUser(user models.LoginUser, db *sql.DB) (string, error) {
	var userFromDb models.User
	dbRes := db.QueryRow("SELECT * FROM Users WHERE Email =?;", user.Email)
	err := dbRes.Scan(&userFromDb.ID,
		&userFromDb.Username,
		&userFromDb.Email,
		&userFromDb.Password,
		&userFromDb.Role,
		&userFromDb.CreatedAt,
		&userFromDb.UpdatedAt)
	fmt.Println(userFromDb)
	if err != nil {
		return "", err
	}
	comparePassword := bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password))
	if comparePassword != nil {
		return "", comparePassword
	}
	jwtToken, err := utils.CreateToken(int64(userFromDb.ID), userFromDb.Role)
	return jwtToken, err
}

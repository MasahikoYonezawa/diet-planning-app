package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type User struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Gender        int       `json:"gender"`
	Age           int       `json:"age"`
	Height        float64   `json:"height"`
	Weight        float64   `json:"weight"`
	ActivityLevel int       `json:"activity_level"`
	BMR           int       `json:"bmr"`
	TDEE          int       `json:"tdee"`
	TargetWeight  float64   `json:"target_weight"`
	Term          int       `json:"term"`
	TermType      int       `json:"term_type"`
	Protein       float64   `json:"protein"`
	Fat           float64   `json:"fat"`
	Carbohydrate  float64   `json:"carbohydrate"`
	CreatedAT     time.Time `json:"created_at"`
	UpdatedAT     time.Time `json:"updated_at"`
}

var DB *sql.DB
var DBUser string
var DBPassword string
var DBName string

func init() {
	loadEnv()
	var err error
	dbconf := DBUser + ":" + DBPassword + "@tcp(db:3306)/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	DB, err = sql.Open("mysql", dbconf)
	if err != nil {
		fmt.Println("データベース　オープン失敗", err)
		panic(err)
	}
	//defer DB.Close()
	//
	err = DB.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗", err)
		return
	} else {
		fmt.Println("データベース接続成功")
	}
}

func loadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	DBUser = os.Getenv("MYSQL_USER")
	DBPassword = os.Getenv("MYSQL_PASSWORD")
	DBName = os.Getenv("MYSQL_DATABASE")
}

func getUsers(limit int) (posts []User, err error) {
	stmt := "SELECT * FROM `users` LIMIT ?"
	rows, err := DB.Query(stmt, limit)
	if err != nil {
		return
	}

	var users []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Gender,
			&user.Age,
			&user.Height,
			&user.Weight,
			&user.BMR,
			&user.TDEE,
			&user.ActivityLevel,
			&user.TargetWeight,
			&user.Term,
			&user.TermType,
			&user.Protein,
			&user.Fat,
			&user.Carbohydrate,
			&user.CreatedAT,
			&user.UpdatedAT)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

func retrieve(id int) (user User, err error) {
	user = User{}
	stmt := "SELECT * FROM users WHERE id = ?"
	err = DB.QueryRow(stmt, id).Scan(&user.ID,
		&user.Name,
		&user.Gender,
		&user.Age,
		&user.Height,
		&user.Weight,
		&user.ActivityLevel,
		&user.BMR,
		&user.TDEE,
		&user.TargetWeight,
		&user.Term,
		&user.TermType,
		&user.Protein,
		&user.Fat,
		&user.Carbohydrate,
		&user.CreatedAT,
		&user.UpdatedAT)
	return
}

func (u *User) create() (err error) {
	stmt := "INSERT INTO users (name, gendar, age, height, weight, bmr, tdee, activity_level, target_weight, term, term_type, protein, fat, carbohydrate) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	err = DB.QueryRow(stmt,
		u.Name,
		u.Gender,
		u.Age,
		u.Height,
		u.Weight,
		u.ActivityLevel,
		u.BMR,
		u.TDEE,
		u.TargetWeight,
		u.Term,
		u.TermType,
		u.Protein,
		u.Fat,
		u.Carbohydrate).Scan(&u.ID)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (u *User) update() (err error) {
	stmt := "UPDATE users set name = ?, gendar = ?, age = ?, height = ?, weight = ?, bmr = ?, tdee = ?, activity_level = ?, target_weight = ?, term = ?, term_type = ?, protein = ?, fat = ?, carbohydrate = ? WHERE id = ?"
	_, err = DB.Exec(stmt,
		u.Name,
		u.Gender,
		u.Age,
		u.Height,
		u.Weight,
		u.ActivityLevel,
		u.BMR,
		u.TDEE,
		u.TargetWeight,
		u.Term,
		u.TermType,
		u.Protein,
		u.Fat,
		u.Carbohydrate,
		u.ID)
	return
}

func (u *User) delete() (err error) {
	stmt := "DELETE FROM users WHERE id = ?"
	_, err = DB.Exec(stmt, u.ID)
	return
}

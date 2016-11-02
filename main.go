package main

import (
	"log"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id        int64  `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
}

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Open("mssql",
		"server=192.168.0.148;Database=database;encrypt=disable;user id=user;password=changeme")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// db, err := sqlx.Open("mssql",
	// 	"server=192.168.0.148;Database=Kobaly;encrypt=disable;user id=sa;password=Kobaly!123")
	// defer db.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/sql", GetTest)
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.POST("/users", PostUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
	}

	r.Run(":8080")
}

type Application struct {
	ApplicationID   int    `db:"applicationID"`
	ApplicationName string `db:"applicationName"`
	CurrentVersion  string `db:"currentVersion"`
}

func GetTest(c *gin.Context) {
	//var db *sqlx.DB
	apps := []Application{}

	err := db.Select(&apps, "SELECT applicationID,applicationName,currentVersion FROM Applications")
	if err != nil {
		c.JSON(500, err.Error())
		return
		//log.Fatal(err)
	}
	c.JSON(200, apps)
}

func GetUsers(c *gin.Context) {
	type Users []User

	var users = Users{
		User{Id: 1, Firstname: "Oliver", Lastname: "Queen"},
		User{Id: 2, Firstname: "Malcom", Lastname: "Merlyn"},
	}

	c.JSON(200, users)

	// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user_id, _ := strconv.ParseInt(id, 0, 64)

	if user_id == 1 {
		content := gin.H{"id": user_id, "firstname": "Oliver", "lastname": "Queen"}
		c.JSON(200, content)
	} else if user_id == 2 {
		content := gin.H{"id": user_id, "firstname": "Malcom", "lastname": "Merlyn"}
		c.JSON(200, content)
	} else {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
	}

	// curl -i http://localhost:8080/api/v1/users/1
}

func PostUser(c *gin.Context) {
	// The futur code…
}

func UpdateUser(c *gin.Context) {
	// The futur code…
}

func DeleteUser(c *gin.Context) {
	// The futur code…
}

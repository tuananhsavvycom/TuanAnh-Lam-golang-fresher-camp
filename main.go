package main

import (
	"Project/component"
	"Project/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type User struct {
	User_id  int    `json:"user_id,omitempty" gorm:"column:user_id"`
	Name     string `json:"name" gorm:"column:name"`
	Password string `json:"password" gorm:"column:password"`
	Email    string `json:"email" gorm:"column:email"`
	Phone    int    `json:"phone" gorm:"column:phone"`
	Address  string `json:"address" gorm:"column:address"`
}
type UserUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
}

func (User) TableName() string {
	return "users"
}
func (UserUpdate) TableName() string {
	return User{}.TableName()
}
func main() {

	//dung bien moi truong
	dsn := os.Getenv("DBConnectionString")

	//ket noi vao DB thong qua thu vien gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//neu co loi thi log ra thoat luon
	if err != nil {
		log.Fatalln(err)
	}
	//va neu khong co loi thi ta se lam gi do
	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}
func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	//khai bao 1 API /Ping khi co request toi /Ping thi sever se tra ve Json :200
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//Crud
	users := r.Group("/users")
	{
		users.POST("", ginuser.CreateUser(appCtx))
		users.GET("/:user_id", ginuser.GetUser(appCtx))
		users.GET("", ginuser.ListUser(appCtx))
		users.PATCH("/:user_id", ginuser.UpdateUser(appCtx))
		users.DELETE("/:user_id", ginuser.DeleteUser(appCtx))
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

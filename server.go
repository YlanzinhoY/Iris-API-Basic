package main

import (
	"github.com/kataras/iris/v12"
	"log"
	"myapp/model"
	"net/http"
)

func createUser(ctx iris.Context) {
	var user model.User

	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		return
	}

	result := model.Db.Create(&user)
	if result.Error != nil {
		log.Fatal("erro em criar")
		return
	}

	err := ctx.JSON(user)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		return
	}

	ctx.StatusCode(http.StatusCreated)
}

func getById(ctx iris.Context) {
	var user model.User
	id := ctx.Params().Get("id")

	result := model.Db.First(&user, id)

	if result.Error != nil {
		ctx.StatusCode(http.StatusNotFound)
		return
	}
	err := ctx.JSON(user)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		return
	}

	ctx.StatusCode(http.StatusOK)

}

func getAll(ctx iris.Context) {
	var user []model.User

	result := model.Db.Find(&user)

	if result.Error != nil {
		log.Fatal("erro em criar")
		return
	}
	err := ctx.JSON(user)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		return
	}

	ctx.StatusCode(http.StatusOK)

}

func putById(ctx iris.Context) {
	var user model.User
	id := ctx.Params().Get("id")

	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		return
	}

	result := model.Db.Model(&model.User{}).Where("id =?", id).Updates(user)
	defer model.Db.Save(&user)

	if result.Error != nil {
		log.Fatal("erro em criar")
		return
	}
	err = ctx.JSON(&user)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		return
	}

	ctx.StatusCode(http.StatusAccepted)
}

func deleteById(ctx iris.Context) {
	var user model.User
	id := ctx.Params().Get("id")
	result := model.Db.Delete(&user, id)

	if result.Error != nil {
		log.Fatal("erro em criar")
		return
	}
	ctx.StatusCode(http.StatusAccepted)

}

func main() {
	app := iris.Default()
	model.DbConnect()
	app.Post("/user", createUser)
	app.Get("/user", getAll)
	app.Get("/user/{id}", getById)
	app.Put("/user/{id}", putById)
	app.Delete("/user/{id}", deleteById)

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}

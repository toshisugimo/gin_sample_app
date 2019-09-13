package controller

import (
	"app/db"
	"app/model"
)

func List() []model.Todo {
	db := db.Connection()
	var todos []model.Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

func Get(id int) model.Todo {
	db := db.Connection()
	var todo model.Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

func Create(text string, status string) {
	db := db.Connection()
	db.Create(&model.Todo{Text: text, Status: status})
	defer db.Close()
}

func Update(id int, text string, status string) {
	db := db.Connection()
	var todo model.Todo
	db.First(&todo, id).Update(&model.Todo{Text: text, Status: status})
	db.Close()
}

func Delete(id int) {
	db := db.Connection()
	var todo model.Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

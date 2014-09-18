package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
)

type PersonCard struct {
	CardId int `json:"cardid" binding:"required"`
	Amount float64
}

//注意大写字母开头为可导出，名字小写是无法和json转换的
type Person struct {
	Name    string      `json:"name"`
	Age     int         `json:"age"`
	Address *string     //= nil
	Card    *PersonCard //非指针型不能为nil
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"hello": "world"})
	})

	m.Get("/person", func(r render.Render) {
		r.JSON(200, Person{
			Name: "newegg",
			Age:  30,
			Card: GetCard(),
		})
	})

	//get route param
	m.Get("/order/:orderid", func(params martini.Params) string {
		return "this.is " + params["orderid"]
	})

	//get query string
	m.Get("/product", func(r *http.Request) string {
		qs := r.URL.Query()
		return qs.Get("productid")
	})

	//反序列化json
	m.Post("/card", binding.Bind(PersonCard{}), func(card PersonCard) string {
		return fmt.Sprintf("CardId: %d", card.CardId)
	})

	//读取并输出json, martini会通过Injector来匹配HandlerFunc的参数类型
	m.Put("/card", binding.Bind(PersonCard{}), func(r render.Render, card PersonCard) {
		r.JSON(http.StatusOK, Person{
			Name: "newegg",
			Age:  30,
			Card: GetCard(),
		})
	})

	//m.Run()
	http.ListenAndServe(":8060", m)
}

func GetCard() *PersonCard {
	return &PersonCard{CardId: 100}
}

package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("tic-tac-toe", func() {
	Title("Tic Tac Toe")
	Description("The infamous war-game.")
	Version("v1")
	Server("tictactoe", func() {
		Description("tictactoe hosts the tictactoe game server.")
		Services("game")

		Host("development", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("game", func() {
	Description("A small tic-tac-toe game server.")

	Method("new", func() {
		Description("Initialize a new board")
		Result(func() {
			Field(1, "id", String)
			Required("id")
		})

		HTTP(func() {
			POST("/")
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("get", func() {
		Description("Obtain a board by ID.")
		Payload(func() {
			Field(1, "board", String, "Board ID")
			Required("board")
		})
		Result(func() {
			Field(1, "a1", String, "Field A1")
			Field(2, "a2", String, "Field A2")
			Field(3, "a3", String, "Field A3")
			Field(4, "b1", String, "Field B1")
			Field(5, "b2", String, "Field B2")
			Field(6, "b3", String, "Field B3")
			Field(7, "c1", String, "Field C1")
			Field(8, "c2", String, "Field C2")
			Field(9, "c3", String, "Field C3")
			Field(10, "winner", String)
		})
		Error("NotFound")

		HTTP(func() {
			GET("/{board}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("NotFound", CodeNotFound)
		})
	})

	Method("move", func() {
		Payload(func() {
			Field(1, "board", String)
			Field(2, "square", String)
			Required("board", "square")
		})
		Result(func() {
			Field(1, "a1", String, "Field A1")
			Field(2, "a2", String, "Field A2")
			Field(3, "a3", String, "Field A3")
			Field(4, "b1", String, "Field B1")
			Field(5, "b2", String, "Field B2")
			Field(6, "b3", String, "Field B3")
			Field(7, "c1", String, "Field C1")
			Field(8, "c2", String, "Field C2")
			Field(9, "c3", String, "Field C3")
			Field(10, "winner", String)
		})
		Error("NotFound")
		Error("BadRequest")

		HTTP(func() {
			POST("/{board}/{square}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("BadRequest", StatusBadRequest)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("NotFound", CodeNotFound)
			Response("BadRequest", CodeUnknown)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

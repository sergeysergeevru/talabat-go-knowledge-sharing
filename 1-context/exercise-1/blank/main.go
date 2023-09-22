package main

import (
	"context"
	"fmt"
	"time"
)

type DB struct {
}

type User struct {
	Name string
}

func (d *DB) SelectUser(ctx context.Context, email string) (User, error) {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-timer.C:
		return User{Name: "Go"}, nil
	case // 1 - complete receiving the signal of context canceling
		return User{}, fmt.Errorf("context canceled")
	}
}

type Handler struct {
	db *DB
}

type Request struct {
	Email string
}

type Response struct {
	User User
}

func (h *Handler) HandleAPI(ctx context.Context, req Request) (Response, error) {
	u, err := h.db.SelectUser(ctx, req.Email)
	if err != nil {
		return Response{}, err
	}

	return Response{User: u}, nil
}

func main() {
	db := DB{}
	handler := Handler{db: &db}
	ctx, cancel := context.WithCancel(context.Background())

	// 2 - complete code that canceling context withing 500 ms

	// if the code executes correctly change duration from 500 ms to 2000 ms

	req := Request{Email: "test@test.com"}
	resp, err := handler.HandleAPI(ctx, req)
	fmt.Println(resp, err)
}
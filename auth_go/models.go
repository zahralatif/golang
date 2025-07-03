package main

type Request struct {
	Command  string `json:"command"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"` 
}
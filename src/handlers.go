package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{

		Todo{Name: "One"},
		Todo{Name: "Two"},
	}
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}

}

func TodoShow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	todoId := vars["todoId"]
	session, _ := store.Get(r, "session-name")
	val := session.Values["foo"]
	fmt.Fprintln(w, "TodoShow: val session", val, todoId)

}

var store = sessions.NewCookieStore([]byte("userinfo"))

func Login(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")
	session, _ := store.Get(r, "session-name")
	session.Options = &sessions.Options{
		Path: "/",
	}
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// Save it.
	session.Save(r, w)
	if username == "username" && password == "password" {
		//fmt.Fprintln(w,"Login successful:",username,password)
		st := Status{Value: "Success"}
		//http.Redirect(w, r, "/views/User.html", http.StatusMovedPermanently)
		if err := json.NewEncoder(w).Encode(st); err != nil {
			panic(err)
		}

	} else {
		fmt.Fprintln(w, "Login failed:", username, password)
	}

}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
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

type WrapHTTPHandler struct {
	m *http.Handler
}

func LoginGet(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./static/login.html")

}

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
	db, err := sql.Open("sqlite3", "./foo.db")
	//stmt, err := db.Prepare("INSERT INTO userinfo(username) values(?)")

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	status := false
	for rows.Next() {

		var usernamedb string
		var passworddb string

		err = rows.Scan(&usernamedb, &passworddb)

		checkErr(err)

		if usernamedb == username && password == passworddb {
			st := Status{Value: "Success"}
			status = true
			//http.Redirect(w, r, "/views/User.html", http.StatusMovedPermanently)
			if err := json.NewEncoder(w).Encode(st); err != nil {
				panic(err)
			}

		}

	}
	if status == false {
		st := Status{Value: "Fail"}
		if err := json.NewEncoder(w).Encode(st); err != nil {
			panic(err)
		}
	}
	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Area struct {
	AreaId   int    `json:"id"`
	AreaName string `json:"name"`
	ParentId int    `json:parentid`
}

type Scene struct {
	AreaId    int    `json:"areaid"`
	SceneId   int    `json:"sceneid"`
	SceneName string `json:"Scenename"`
}

type Zone struct {
	ZoneId    int    `json:"zoneid"`
	SceneId   int    `json:"sceneid"`
	ZoneName  string `json:"zonename"`
	ZoneLevel int    `json:zonelevel`
}

func GetAreas(w http.ResponseWriter, r *http.Request) {
	var AreaIdVal, ParentIdVal int
	var AreaNameVal string

	arCollection := make([]Area, 2)

	db, err := sql.Open("sqlite3", "./foo.db")
	rows, err := db.Query("SELECT * FROM Area")
	checkErr(err)

	for rows.Next() {

		err = rows.Scan(&AreaIdVal, &AreaNameVal, &ParentIdVal)
		if AreaIdVal != 0 {
			ar := Area{AreaId: AreaIdVal, AreaName: AreaNameVal, ParentId: ParentIdVal}

			arCollection = append(arCollection, ar)

			checkErr(err)
		}

	}
	if err := json.NewEncoder(w).Encode(arCollection); err != nil {
		panic(err)
	}

}

func GetScene(w http.ResponseWriter, r *http.Request) {

	var AreaIdVal, SceneIdVal int
	var SceneNameVal string
	vars := mux.Vars(r)
	areaId := vars["areaId"]

	arCollection := make([]Scene, 0)

	db, err := sql.Open("sqlite3", "./foo.db")
	rows, err := db.Query("SELECT * FROM Scene where areaid=" + areaId)
	checkErr(err)

	for rows.Next() {

		err = rows.Scan(&SceneIdVal, &AreaIdVal, &SceneNameVal)
		if AreaIdVal != 0 {
			ar := Scene{SceneId: SceneIdVal, AreaId: AreaIdVal, SceneName: SceneNameVal}

			arCollection = append(arCollection, ar)

			checkErr(err)
		}

	}
	if err := json.NewEncoder(w).Encode(arCollection); err != nil {
		panic(err)
	}

}
func GetZone(w http.ResponseWriter, r *http.Request) {

	var ZoneIdVal, SceneIdVal, ZoneLevelVal int
	var ZoneNameVal string
	vars := mux.Vars(r)
	sceneId := vars["sceneId"]
	fmt.Println(sceneId)
	arCollection := make([]Zone, 0)

	db, err := sql.Open("sqlite3", "./foo.db")
	rows, err := db.Query("SELECT * FROM zone where sceneid=" + sceneId)
	checkErr(err)

	for rows.Next() {

		err = rows.Scan(&ZoneIdVal, &SceneIdVal, &ZoneNameVal, &ZoneLevelVal)
		if ZoneIdVal != 0 {
			ar := Zone{ZoneId: ZoneIdVal, SceneId: SceneIdVal, ZoneName: ZoneNameVal, ZoneLevel: ZoneLevelVal}

			arCollection = append(arCollection, ar)

			checkErr(err)
		}

	}
	if err := json.NewEncoder(w).Encode(arCollection); err != nil {
		panic(err)
	}

}

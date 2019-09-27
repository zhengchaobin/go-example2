package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
	"runtime"
)

// http://localhost:8088/user/124?name=zcb GET
func getuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	name := params.Get("name")
	fmt.Fprintf(w, "you are get user %s, name %s", uid, name)

	go say(name, w)
	say("world", w)
}

//http://localhost:8088/user/124 POST
func modifyuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are modify user %s", uid)
}

//http://localhost:8088/user/124 DELETE
func deleteuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are delete user %s", uid)
}

//http://localhost:8088/user/124 PUT
func adduser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are add user %s", uid)
	//animal := zcb.Cat{uid}
	//return animal
}

func say(str string, w http.ResponseWriter) {
	for i := 1; i <= 5; i++ {
		runtime.Gosched()
		fmt.Fprintln(w, "hello", str)
	}
}


func main() {
	mux := routes.New()
	mux.Get("/user/:uid", getuser)
	mux.Post("/user/:uid", modifyuser)
	mux.Del("/user/:uid", deleteuser)
	mux.Put("/user/:uid", adduser)
	http.Handle("/", mux)
	http.ListenAndServe(":8088", nil)
}
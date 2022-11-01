package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

func main() {
	server := http.Server{
		Addr: "go:8000",
	}
	http.HandleFunc("/users", handleGetList)
	http.HandleFunc("/users/", handleRequest)

	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleGetList(w http.ResponseWriter, r *http.Request) {
	var err error
	fmt.Println("start handleGetList")
	users, err := getUsers(100)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	output, err := json.MarshalIndent(&users, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Write(output)
	fmt.Println("end handleGetList", users)
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// Base は，path の最後の要素を返します。 末尾のスラッシュは，最後の要素を抽出する前に削除されます。
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	user, err := retrieve(id)
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	contentLength := r.ContentLength
	contentBody := make([]byte, contentLength)
	r.Body.Read(contentBody)

	var user User
	err = json.Unmarshal(contentBody, &user)
	if err != nil {
		return err
	}

	err = user.create()
	if err != nil {
		return err
	}

	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		return err
	}

	w.WriteHeader(200)
	w.Write(output)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	user, err := retrieve(id)
	if err != nil {
		return
	}

	contentLength := r.ContentLength
	contentBody := make([]byte, contentLength)
	r.Body.Read(contentBody)

	err = json.Unmarshal(contentBody, &user)
	if err != nil {
		return
	}

	err = user.update()
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(output)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	user, err := retrieve(id)
	if err != nil {
		return
	}

	err = user.delete()
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(output)
	return
}

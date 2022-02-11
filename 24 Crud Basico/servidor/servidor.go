package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// User Create
func UserCreate(w http.ResponseWriter, r *http.Request) {
	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Request fail!"))
		return
	}

	var user user

	if erro = json.Unmarshal(corpoDaRequisicao, &user); erro != nil {
		w.Write([]byte("Convert user to struct fail!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("DB connection fail!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into users (name, email) value (?, ?)")
	if erro != nil {
		w.Write([]byte("Create statement fail!"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(user.Name, user.Email)
	if erro != nil {
		w.Write([]byte("Insert statement fail!"))
		return
	}

	idInserted, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("ID return fail!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted success! ID: %d", idInserted)))

}

// Users Search
func UsersSearch(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("DB request fail!"))
	}
	defer db.Close()

	linhas, erro := db.Query("select * from users")
	if erro != nil {
		w.Write([]byte("Search fail!"))
	}
	defer linhas.Close()

	var users []user
	for linhas.Next() {
		var user user

		if erro := linhas.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("Users scan fail!"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Users convert json fail!"))
		return
	}
}

// User Search
func UserSearch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Param convert fail!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("DB request fail!"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from users where id = ?", ID)
	if erro != nil {
		w.Write([]byte("User search fail!"))
		return
	}

	var user user
	if linha.Next() {
		if erro := linha.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("User scan fail!"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte("User convert json fail!"))
		return
	}

}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Param convert fail! Update"))
		return
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Body read fail! Update"))
		return
	}

	var user user
	if erro := json.Unmarshal(corpoDaRequisicao, &user); erro != nil {
		w.Write([]byte("Read user fail! Update"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("DB request fail! Update"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update users set name = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Statement fail! Update"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(user.Name, user.Email, ID); erro != nil {
		w.Write([]byte("Update fail! Update"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// User delete
func UserDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Param convert fail! Delete"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("DB request fail! Delete"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from users where id = ?")
	if erro != nil {
		w.Write([]byte("Statement fail! Delete"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Delete user Fail! Delete"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

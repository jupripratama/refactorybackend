package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"

	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

func dbConns() (db *sql.DB) {
	er := godotenv.Load(".env")

	if er != nil {
		log.Fatalf("Error loading .env file")
	}
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":3306)/"+dbName)
	// db, err := sql.Open("mysql", "root:password@tcp(host:3306)/refactory")
	if err != nil {
		panic(err.Error())
	}
	return db

}

func main() {
	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days
	isProd := false             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	er := godotenv.Load(".env")

	if er != nil {
		log.Fatalf("Error loading .env file")
	}

	goth.UseProviders(

		facebook.New(os.Getenv("FB_KEY"), os.Getenv("FB_SECRET"), os.Getenv("FB_CALLBACKURL"), "email"),

		google.New(os.Getenv("GLE_KEY"), os.Getenv("GLE_SECRET"), os.Getenv("GLE_CALLBACKURL"), "email", "profile"),
	) //use src of providers:google ,facebook &..

	p := pat.New()

	p.Get("/auth/{provider}/callback", completeauth)

	p.Get("/auth/{provider}", beginauth)

	p.HandleFunc("/", home)

	log.Println("listening on http://localhost:8000")

	p.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8000", p)
}

func home(res http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles("templates/social.html")
	t.Execute(res, nil)

}

func beginauth(res http.ResponseWriter, req *http.Request) {

	gothic.BeginAuthHandler(res, req) //authentication with provider

}

func completeauth(res http.ResponseWriter, req *http.Request) {

	user, err := gothic.CompleteUserAuth(res, req) //get autherised data's (name,id,profile)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	db := dbConns()

	gmail := user.Email
	firstname := user.FirstName
	lastname := user.LastName
	userid := user.UserID
	provider := user.Provider
	name := user.Name
	p, err := db.Prepare("INSERT INTO oauths(gmail,firstname,lastname,userid,provider) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	p.Exec(gmail, firstname, lastname, userid, provider)

	fmt.Println("customer email:", gmail)

	fmt.Println("customer Firstname:", firstname)

	fmt.Println("customer Lastname:", lastname)

	fmt.Println("customer user id:", userid)

	fmt.Println("customer data provider:", provider)

	fmt.Println("customer raw data provider:", name)

	t, _ := template.ParseFiles("templates/success.html")

	t.Execute(res, user)

}

package main

import (
	"fmt"
	"net/http"
	"yaroslavl-parkings/api/views"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	template, err := template.ParseFiles("index.html")
// 	if err != nil {
// 		panic(err)
// 	}

// 	template.ExecuteTemplate(w, "index.html", nil)
// }

// func privateHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello"))
// }

func main() {
	viewsRouter := views.NewViewsRouter("web/html")
	http.Handle("/", viewsRouter)

	// sessionStorage := session.NewSessionStorage()
	// authRouter := auth.NewAuthRouter(sessionStorage)
	// http.Handle("/auth/", http.StripPrefix("/auth", authRouter))

	// authMiddleware := middlewares.NewAuthMiddleware(sessionStorage)
	// http.HandleFunc("/private", authMiddleware(privateHandler))

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db := database.NewDatabase("host=localhost user=postgres password=password dbname=parkings port=5432 sslmode=disable")
	// db.InitTables()

	// http.HandleFunc("/", index)
	files := http.FileServer(http.Dir("./web"))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	fmt.Println("Server started on" + "http://localhost:8880")
	http.ListenAndServe(":8880", nil)
}

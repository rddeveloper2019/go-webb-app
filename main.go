package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	user := "Alex"
	responseMsg := fmt.Sprintf("<h1>Hello %s! Lenslocked</h1>", user)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(responseMsg))
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>Contact me by <a href="mailto: samsobaka@kot.ru">samsobaka@kot.ru</a></h1>`)
}
func NoFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Sorry, page not found", http.StatusNotFound)
}
func FaqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	responseStr := `<ul>
	<li>Q: Is there a free version?<b>A: Yes! We offer a free trial for 30 days on any paid plans.</b>
	</li><li>Q: What are your support hours?<b>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</b></li>
	<li>Q: How do I contact support?<b>A: Email us - support@lenslocked.com</b></li>
	</ul>`
	fmt.Fprint(w, responseStr)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "userID")
	userID := user[len("userID")+1:]
	//// fetch `"key"` from the request context
	//ctx := r.Context()
	//key := ctx.Value("key").(string)
	//fmt.Println(user, key)
	//// respond to the client
	w.Write([]byte(fmt.Sprintf("Hi %v", userID)))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", MainHandler)
	r.Get("/users/{userID}", UserHandler)
	r.Get("/contact", ContactHandler)
	r.Get("/faq", FaqHandler)
	r.NotFound(NoFoundHandler)
	fmt.Println("Server started at http://localhost:3000")
	http.ListenAndServe(":3000", r)
}

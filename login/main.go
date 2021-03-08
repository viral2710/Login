package main
import (
	"html/template"
	"net/http"
	_"sql/mysql"
	"database/sql"
	"fmt"
	_"golang.org/x/crypto/bcrypt"
	
)
var db *sql.DB
var err error
var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
type ds struct{
	Username string
	Password string
}
func main(){
	http.HandleFunc("/login", login)
	http.HandleFunc("/login1", login1)
	http.HandleFunc("/process", process)
	http.HandleFunc("/signup",signup)
	http.HandleFunc("/user",view)
	http.HandleFunc("/sp",sp)
	http.ListenAndServe(":8080", nil)
}
func view(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("DB selected successfully..")
	}
	
	rows, err := db.Query("SELECT * FROM users")
	
	var d []ds
	for rows.Next() {
		var username string
		var password string
		err = rows.Scan(&username, &password)
		fmt.Println(username)
		fmt.Println(password)
		d = append(d,ds{Username:username,Password:password})
		
	}


	tpl.ExecuteTemplate(w, "view.gohtml",d)
	
}
func login(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.gohtml",nil)
}
func login1(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login1.gohtml",nil)
}
func process(w http.ResponseWriter,r *http.Request){
	ui := r.FormValue("userid")
	passwd := r.FormValue("passwd")
	var pass string
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("Database created successfully")
	}
	_,err = db.Exec("USE test")
	if err != nil {
	fmt.Println(err.Error())
	} else {
	fmt.Println("DB selected successfully..")
	}
	err = db.QueryRow("select password from users WHERE username =?",ui).Scan(&pass)

	if err != nil {
		fmt.Println("fail1")
		tpl.ExecuteTemplate(w, "login1.gohtml",nil)
		return
	}
	fmt.Println(pass)
	fmt.Println(passwd)
	

	if pass == passwd {
		d := struct {
			First string
			Last  string
		}{
			First: ui,
			Last:  passwd,
		}
		tpl.ExecuteTemplate(w, "processor.gohtml", d)
		return
		
	}
	fmt.Println("fail")
	tpl.ExecuteTemplate(w, "login1.gohtml",nil)	
		
}
func signup(w http.ResponseWriter,r *http.Request){
	tpl.ExecuteTemplate(w, "signup.gohtml",nil)
}
func sp(w http.ResponseWriter,r *http.Request){
	ui := r.FormValue("userid")
	passwd := r.FormValue("passwd")
	passwd1 := r.FormValue("passwd1")
	if passwd == passwd1 {
		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("Database created successfully")
		}
		_,err = db.Exec("USE test")
		if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println("DB selected successfully..")
		}
		INSERT,err := db.Prepare("INSERT  INTO `users` VALUES (?,?);")
		_,err = INSERT.Exec(ui,passwd)
		if err == nil {
			tpl.ExecuteTemplate(w, "sp.gohtml",nil)
			return
		}
		tpl.ExecuteTemplate(w, "signup1.gohtml",nil)	
	} else {
		tpl.ExecuteTemplate(w, "signup1.gohtml",nil)	
	}
}
func check(Uids string, r *http.Request,w http.ResponseWriter) bool{
	fmt.Println(Uids)
		
		c, err := r.Cookie(Uids)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return	true
		}
		if c.Name != Uids{
			fmt.Println("error")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		return true
}


		

	

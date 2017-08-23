package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
        "html/template"
        "io/ioutil"
)


type Page struct {
    Title string
    Body  []byte
}

type Widgets struct {
	gorm.Model
	WidgetName  string
	WidgetCount int
}

//func (p *Page) save() error {
//    filename := p.Title + ".txt"
//    return ioutil.WriteFile(filename, p.Body, 0644)
//}

func loadPage(title string) (*Page, error) {
        filename := title + ".txt"
        body, err := ioutil.ReadFile(filename)
        if err != nil {
                return nil, err
        }
        return &Page{Title: title, Body: body}, nil
  }


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "webSphere %s!", r.URL.Path[1:])
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    fmt.Printf("%+v\n", "6")
    t, _ := template.ParseFiles(tmpl + ".html")
    fmt.Printf("%+v\n", t)
    t.Execute(w, p)
}

//func viewHandler(w http.ResponseWriter, r *http.Request) {
//    title := r.URL.Path[len("/view/"):]
//    p, _ := loadPage(title)
//    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
//}


func viewHandler(w http.ResponseWriter, r *http.Request) {
        title := r.URL.Path[len("/view/"):]
        p, _ := loadPage(title)
        renderTemplate(w, "view", p)
  }




func editHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v\n", "1")
        title := r.URL.Path[len("/edit/"):]
	fmt.Printf("%+v\n", "2")
        p, err := loadPage(title)
	fmt.Printf("%+v\n", "3")
        if err != nil {
                p = &Page{Title: title}
        }
	fmt.Printf("%+v\n", "4")
        defer renderTemplate(w, "edit", p)
	fmt.Printf("%+v\n", "5")
  }


func main() {
//        var match string
	db, err := gorm.Open("mysql", "webSphere:ContainerBleed@/Widgets?charset=utf8&parseTime=True&loc=Local")
	_ = err
        var widget Widgets
        db.AutoMigrate(&Widgets{})
        //db.Create(&Widgets{WidgetName: "Sphere Widget", WidgetCount: 1})
        db.First(&widget, "widget_name = ?", "Sphere Widget")
        //fmt.Printf(widget.WidgetName)
	fmt.Printf("%+v\n", widget.WidgetName)
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.ListenAndServe(":8080", nil)
        defer db.Close()
}

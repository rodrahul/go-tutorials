package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("./template/02_date-formatting.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2016")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "02_date-formatting.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}

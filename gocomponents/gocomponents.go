package gocomponents

import (
        "log"
        "html/template"
        "net/http"
)

func Header(w http.ResponseWriter, r *http.Request) {        
        t, err := template.ParseFiles("gocomponents/templates/header.html")
        
        if err != nil  {
                log.Fatal(err)
        }
        
        templateData := map[string]string{
                "Name": "Eggs Benedict",
        }
        t.Execute(w, templateData)
}
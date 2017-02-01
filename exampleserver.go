package main

import (
        "fmt"
        "net/http"
        "os"
        "bufio"
        "mime"
        "path/filepath"
        "log"
        "html/template"
        "./gocomponents"
)

func static(w http.ResponseWriter, r *http.Request) {
        // Get path and add gocomponents to it to get the correct file
        path := "gocomponents/" + r.URL.Path

        // Open file
        data, err := os.Open(path)

        // If an error occurred, render an error page
        if err != nil {
                fmt.Fprintf(w, err.Error())
                return
        }

        // Close data on next return
        defer data.Close()

        // Create a scanner
        scanner := bufio.NewScanner(data)

        var fileContent string

        // Scan file and save result in fileContent
        for scanner.Scan() {
                fileContent += scanner.Text() + "\n"
        }

        // If the scanner ran into an error, display it
        if err := scanner.Err(); err != nil {
                fmt.Fprintf(w, err.Error())
                return
        }

        // Get mime type by extension
        mimeType := mime.TypeByExtension(filepath.Ext(path))

        // Set content type 
        w.Header().Set("Content-Type", mimeType)

        // Show file
        fmt.Fprintf(w, fileContent)
        return;
}

func index(w http.ResponseWriter, r *http.Request) {
        // Parse required files
        t, err := template.ParseFiles("./templates/index.html", "./gocomponents/templates/header.html", "./gocomponents/templates/sidebar.html", "./gocomponents/templates/colors.html")
        
        // Make sure there are no errors
        if err != nil  {
                log.Fatal(err)
        }
        
        // Data for the template
        data := gocomponents.TemplateData{
                SiteTitle: "Kees", 
                Colors: map[string]string{
                        "primary": "#2196F3",
                        "primaryDark": "#1976D2",
                        "primaryText": "#FFFFFF",
                        "accent": "#FF6E40",
                        "accentText": "rgba(0, 0, 0, .87)",
                },
                HasSidebar: true,
                SidebarItems: map[string]string{
                        "kaas": "Kaas",
                        "kees-btn": "Koel",
                },
                Components: map[string]template.HTML{
                        "MyCard": gocomponents.Card("henk jan", "<h1>kees</h1>"),
                        "MyButton": gocomponents.Button("kees iscool", "klikkie"),
                        "MyCheckbox": gocomponents.CheckBox("myCheckbox", "kees", "Gratis YT monnie?"),
                        "MyRadio": gocomponents.Radio("myRadio", "jan henk", "Radio button"),
                        "MyInput": gocomponents.Input("myInput", "kees henk", "Your name", -1),
                        "CompleteCard": gocomponents.Card("complete-card", `
                                <h3>kees</h3>
                                ` + string(gocomponents.Button("kees iscool", "klikkie") + `
                                <p>Cool, ain't it?</p>
                        `)),
                },
        }

        // Execute templates
        t.ExecuteTemplate(w, "layout", data)
}

func main() {
        http.HandleFunc("/static/", static)
        http.HandleFunc("/", index)
        http.ListenAndServe(":666", nil)
}
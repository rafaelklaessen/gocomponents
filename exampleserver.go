package main

import (
        "net/http"
        "log"
        "html/template"
        "./gocomponents"
)

// The index function serves at localhost:9090/
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
                        "accentText": "#000000",
                },
                HasSidebar: true,
                SidebarItems: map[string]string{
                        "kaas": "Kaas",
                        "kees-btn": "Koel",
                },
                Components: map[string]template.HTML{
                        "MyCard": gocomponents.Card("henk jan", "<h1>kees</h1>"),
                        "MyButton": gocomponents.Button("primary-btn", "klikkie"),
                        "MyButtonTwo": gocomponents.Button("accent-btn", "klikkie"),
                        "MyButtonThree": gocomponents.Button("flat-btn", "klikkie"),
                        "MyCheckbox": gocomponents.CheckBox("myCheckbox", "kees", "Free YT money?"),
                        "MyRadio": gocomponents.Radio("myRadio", "jan henk", "Radio button"),
                        "MyInput": gocomponents.Input("myInput", "kees henk", "Your name", -1),
                        "MyList": gocomponents.List("kees-list", map[string]string{
                                "Pizza": "omnom",
                                "More pizza": "extra omnom",
                                "Infinite pizza": "ultramegaomnom",
                        }),
                        "MyProgressBar": gocomponents.ProgressBar("kees-bar", "Progress", false),
                        "MyTabGroup": gocomponents.TabGroup("kees is-cool", map[string]string{
                                        "tabName": "kees",
                                        "tabContent": "<p>cool</p>",
                                }, 
                                map[string]string{
                                        "tabName": "Henk",
                                        "tabContent": "Oh boi",
                        }),
                        "MyMenu": gocomponents.Menu("kees-menu", "MyMenu", map[string]string{
                                "item-one": "Item ONEEE",
                                "item-two": "WE ARE NUMBER TWO",
                        }),
                        "MyActionButton": gocomponents.FloatingActionButton("my-action-btn", "add"),
                        "CompleteCard": gocomponents.Card("complete-card", `
                                <h3>kees</h3>
                                ` + string(gocomponents.Button("kees iscool", "klikkie")) + `
                                <p>Cool, ain't it?</p>
                        `),
                },
        }

        // Execute templates
        t.ExecuteTemplate(w, "layout", data)
}

func main() {
        http.HandleFunc("/static/", gocomponents.ServeStatic)
        http.HandleFunc("/", index)
        http.ListenAndServe(":9090", nil)
}
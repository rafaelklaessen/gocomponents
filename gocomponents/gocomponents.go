package gocomponents

import (
        "html/template"
        "strconv"
)

type TemplateData struct {
        SiteTitle    string
        Colors       map[string]string
        HasSidebar   bool
        SidebarItems map[string]string
        Components   map[string]template.HTML
}

func Card(classes string, content string) template.HTML {
        return template.HTML(`
                <article class="` + classes + ` card">
                        ` + content + `
                </article>
        `)
}

func Button(classes string, content string) template.HTML {
        return template.HTML(`
                <button class="` + classes + `btn">` + content + `</button>
        `)
}

func CheckBox(id string, classes string, label string) template.HTML {
        return template.HTML(`
                <div class="` + classes + ` checkbox-container">
                        <input id="` + id + `" class="checkbox" type="checkbox">
                        <label for="` + id + `" class="label">` + label + `</label>
                </div>
        `)
}

func Radio(id string, classes string, label string) template.HTML {
        return template.HTML(`
                <div class="` + classes + ` radio-container">
                        <input id="` + id + `" class="radio-btn" type="radio">
                        <label for="` + id + `" class="label">` + label + `</label>
                </div>
        `)
}

func Input(id string, classes string, label string, maxlength int) template.HTML {        
        var maxlengthAttr string
        
        if maxlength == -1 {
                maxlengthAttr = " "
        } else {
                maxlengthAttr = `maxlength="` + strconv.Itoa(maxlength) + `"` 
        }
        
        return template.HTML(`
                <div class="` + classes + ` input-container">
                        <input id="` + id + `" ` + maxlengthAttr + `class="input" type="text">
                        <label for="` + id + `" class="input-label">` + label + `</label>
                </div>
        `)
}

func List(classes string, listItems map[string]string) template.HTML {
        var listItemsHtml string
        
        for itemTitle, itemDescription := range listItems {
                listItemsHtml += `
                        <li class="list-item">
                                <h4 class="list-item-title">` + itemTitle + `
                                <p class="list-item-description">` + itemDescription + `
                        </li>
                `
        }
        
        return template.HTML(`
                <ul class="` + classes + ` list">
                        ` + listItemsHtml + `
                </ul>
        `)
}

func ProgressBar(classes string, label string, determinate bool) template.HTML {
        var determinateAttr string

        if determinate {
                determinateAttr = `determinate="determinate"`
        } else {
                determinateAttr = `determinate="indeterminate"`
        }
        
        return template.HTML(`
                <div class="` + classes + ` progressbar-container">
                        <label class="progressbar-label">` + label + `</label>
                        <div class="progressbar" ` + determinateAttr + `>
                                <div class="progress" style="width: 10%";></div>
                        </div>
                </div>
        `)
}

func TabGroup(classes string, tabs ...map[string]string) template.HTML {
        tabNav := `<nav class="tab-nav">
                        <ul class="nav-items">`
        
        tabsContainer := `<section class="tab-container">`

        for _, tab := range tabs {
                tabNav += `<li class="nav-item">
                                <h4 class="tab-name">` + tab["tabName"] + `</h4>
                           </li>`
                
                tabsContainer += `<section class="tab">
                                        ` + tab["tabContent"] + `
                                  </section>`
        }

        tabNav += `</ul></nav>`
        tabsContainer += `</section>`

        return template.HTML(`
                <section class="` + classes + ` tabgroup">
                        ` + tabNav + `
                        ` + tabsContainer + `
                </section>
        `)
}
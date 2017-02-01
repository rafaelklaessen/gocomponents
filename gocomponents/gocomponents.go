package gocomponents

import (
        "html/template"
        "strconv"
)

type TemplateData struct {
        SiteTitle string
        Colors map[string]string
        HasSidebar bool
        SidebarItems map[string]string
        Components map[string]template.HTML
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
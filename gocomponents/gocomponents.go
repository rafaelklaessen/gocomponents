package gocomponents

import (
        "html/template"
        "strconv"
)

// The templatedata struct is the struct that should be used for passing to the template
type TemplateData struct {
        SiteTitle    string
        Colors       map[string]string
        HasSidebar   bool
        SidebarItems map[string]string
        Components   map[string]template.HTML
}

// The card function gives a card
func Card(classes string, content string) template.HTML {
        return template.HTML(`
                <article class="` + classes + ` card">
                        ` + content + `
                </article>
        `)
}

// The button function gives a button
func Button(classes string, content string) template.HTML {
        return template.HTML(`
                <button class="` + classes + `btn">` + content + `</button>
        `)
}

// The checkbox function gives a checkbox with label
func CheckBox(id string, classes string, label string) template.HTML {
        return template.HTML(`
                <div class="` + classes + ` checkbox-container">
                        <input id="` + id + `" class="checkbox" type="checkbox">
                        <label for="` + id + `" class="label">` + label + `</label>
                </div>
        `)
}

// The radio function gives a radio button with label
func Radio(id string, classes string, label string) template.HTML {
        return template.HTML(`
                <div class="` + classes + ` radio-container">
                        <input id="` + id + `" class="radio-btn" type="radio">
                        <label for="` + id + `" class="label">` + label + `</label>
                </div>
        `)
}

// The input function gives an input. It also takes a maxlength parameter.
//  -1 means that there is no maxlength
func Input(id string, classes string, label string, maxlength int) template.HTML {        
        var maxlengthAttr string
        
        // If maxlength is -1, just don't generate a maxlength attribute
        if maxlength != -1 {
                maxlengthAttr = `maxlength="` + strconv.Itoa(maxlength) + `"` 
        }
        
        return template.HTML(`
                <div class="` + classes + ` input-container">
                        <input id="` + id + `" ` + maxlengthAttr + `class="input" type="text">
                        <label for="` + id + `" class="input-label">` + label + `</label>
                </div>
        `)
}

// The list function gives a list. Each list item consists of a title and a 
// description. Those are given in the listItems map
func List(classes string, listItems map[string]string) template.HTML {
        // listItemsHtml will contain the HTML of all the list items. It's inserted in a ul element.
        var listItemsHtml string
        
        // Loop through all list items and add HTML for them to listItemsHtml.
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

// The progress bar function gives a progress bar. This progress bar can be both determinate and 
// inderminate, as set by the determinate parameter
func ProgressBar(classes string, label string, determinate bool) template.HTML {
        var determinateAttr string

        // Generate determinate attribute (will be put into the HTML)
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

// The tabGroup function gives a tab group. This means a navbar with the tab titles and two tabs. 
// When you click a tab title, the corresponding tab will be opened
func TabGroup(classes string, tabs ...map[string]string) template.HTML {
        // Add opening tags for the tab nav
        tabNav := `<nav class="tab-nav">
                        <ul class="nav-items">`
        
        // Add opening tag for the tab container
        tabsContainer := `<section class="tab-container">`

        // Loop through the given tabs
        for _, tab := range tabs {
                // Add the tab title to the tabNav
                tabNav += `<li class="nav-item">
                                <h4 class="tab-name">` + tab["tabName"] + `</h4>
                           </li>`
                
                // Add the tab content (the tab itself) to the tab container
                tabsContainer += `<section class="tab">
                                        ` + tab["tabContent"] + `
                                  </section>`
        }

        // Add closing tags for the tab nav and the tab container
        tabNav += `</ul></nav>`
        tabsContainer += `</section>`

        // Put everything together in a tabgroup element
        return template.HTML(`
                <section class="` + classes + ` tab-group">
                        ` + tabNav + `
                        ` + tabsContainer + `
                </section>
        `)
}

// The menu function gives a button that opens a menu when you click it
func Menu(classes string, btnText string, menuItems map[string]string) template.HTML {
        // The menu will be saved as a JSON object so that the JavaScript can read it. 
        // The menu JSON object will be put into the data-menu attribute of the button. 

        // Open object
        menu := `{`

        // Loop through all menu items and them to the menu object
        for itemClass, itemContent := range menuItems {
                menu += `{"class": "` + itemClass + `", "content": "` + itemContent + `"},`
        }

        // Close the object
        menu += `}`

        return template.HTML(`
                <div class="` + classes + ` menu-container">
                        <button class="menu-btn btn" data-menu='` + menu + `'>` + btnText + `</button>
                </div>
        `)
}

// The floating action button function gives a material design floating action button
func FloatingActionButton(classes string, btnIcon string) template.HTML {
        return template.HTML(`
                <button class="` + classes + ` floating-action-button material-icons">` + btnIcon + `</button>
        `)
}
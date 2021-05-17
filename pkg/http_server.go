package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/ttudrej/pokertrainer/debugging"
)

var tpl *template.Template

func init() {
	// tpl = template.Must(template.ParseGlob("web/*"))
	// tplIndex = template.Must(template.ParseFiles("web/index.html"))
	// tplHeader = template.Must(template.ParseFiles("web/include-header.html"))
	// tplFooter = template.Must(template.ParseFiles("web/include-footer.html"))
	// indexTemplate := template.New("index_template")
	// indexTemplate = template.Must(template.ParseFiles("web/include-header.html", "web/index.html", "web/include-footer.html"))
	// Info.Println("index err: ", errIndex)
	// Info.Println("header err: ", errHeader)
	// Info.Println("footer err: ", errFooter)
}

// #################################################################
func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	Info.Println(debugging.ThisFunc())
	// With this fucntion I can check if my filepath is working for serving static files such as CSS or Templates etc
	// IMPORTANT:I failed to add static files because Go will use the current Directory you are in as the App's ROOT.
	// If I run main.go from GolangBlog, the root is/Users/jorn/Documents/Golang/src/github.com/jschalkwijk/GolangBlog
	// If I run it from jschalkwijk (my github folder)
	_, err := os.Stat(filepath.Join(".", "web/styles", "main.css"))
	Info.Println("css file read error: ", err)

	templates, err := template.ParseFiles("web/include-header.html", "web/index.html", "web/include-footer.html")

	if err != nil {
		Info.Println(err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}

	// templates.ExecuteTemplate(w, "index_template", person1)
	//Info.Println("table id: ", prPtr.tableList[0].tableID)
	templates.ExecuteTemplate(w, "index_template", prPtr.tableList[0].tableStatePtr)
}

// #################################################################
func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	Info.Println(debugging.ThisFunc())

	r.ParseForm()
	r.PostForm.Get("deal_new")

	_ = executeHand(prPtr.tableList[0])

	// Re-draw the page.
	_ = updateTableStateStruct(prPtr.tableList[0])
	http.Redirect(w, r, "/", 302)
}

// #################################################################
func dealNewHandPostHandler(w http.ResponseWriter, r *http.Request) {
	Info.Println(debugging.ThisFunc())

	r.ParseForm()
	r.PostForm.Get("id_deal_new_hand")

	_ = executeHand(prPtr.tableList[0])

	// Re-draw the page.
	_ = updateTableStateStruct(prPtr.tableList[0])
	http.Redirect(w, r, "/", 302)
}

// <input type="hidden" id="id_deal_new_hand" name="name_deal_new_hand" value="1">

// #################################################################
func executeNextStepPostHandler(w http.ResponseWriter, r *http.Request) {
	Info.Println(debugging.ThisFunc())
	r.ParseForm()
	r.PostForm.Get("id_next_step")

	_ = executeNextStep(prPtr.tableList[0])

	// Re-draw the page.
	_ = updateTableStateStruct(prPtr.tableList[0])
	http.Redirect(w, r, "/", 302)
}

/*
<form method="POST" action="dealNextStep">
<div>
  <input type="hidden" id="id_deal_next_step" name="name_deal_next_step" value="1">
  <input class="button_clickable_table" type="submit" value="Deal Next Step">
</div>
</form>
*/

// #################################################################
func postBetPostHandler(w http.ResponseWriter, r *http.Request) {
	Info.Println(debugging.ThisFunc())
	r.ParseForm()
	// postBet_value := r.PostForm.Get("id_postBet")
	fractionOfPot_value := r.PostForm.Get("name_fractionOfPot")
	tableID_value := r.PostForm.Get("name_tableID")
	playerID_value := r.PostForm.Get("name_playerID")

	Info.Println()
	Info.Println("describe: postBetPostHandler; tIDv / pIDv / fractionV: ", tableID_value, playerID_value, fractionOfPot_value)
	Info.Println("postBetPostHandler; tIDv / pIDv / fractionV: ", tableID_value, playerID_value, fractionOfPot_value)
	Info.Println()

	Info.Printf("postform: [%s] %T \n", tableID_value, r.PostForm.Get("id_tableID"))
	Info.Printf("postform: [%s] %T \n", playerID_value, r.PostForm.Get("id_playerID"))
	Info.Printf("postform: [%s] %T \n", fractionOfPot_value, r.PostForm.Get("id_fractionOfPot"))

	// _ = postBet(prPtr.tableList[0])
	_ = postBet(tableID_value, playerID_value, fractionOfPot_value)

	// Re-draw the page.
	_ = updateTableStateStruct(prPtr.tableList[0])
	http.Redirect(w, r, "/", 302)
}

// #################################################################
func fullResetPostHandler(w http.ResponseWriter, r *http.Request) {
	Info.Println(debugging.ThisFunc())
	r.ParseForm()
	// postBet_value := r.PostForm.Get("id_postBet")

	tableID_value := r.PostForm.Get("name_tableID")

	Info.Printf("postform: [%s] %T \n", tableID_value, r.PostForm.Get("id_tableID"))
	tPtr, _ := getTablePtrFromTIDStr(tableID_value)
	_ = resetTable(tPtr)

	// Re-draw the page.
	_ = updateTableStateStruct(prPtr.tableList[0])
	http.Redirect(w, r, "/", 302)
}

// #################################################################

// #################################################################
func startWebServer() {
	// http.HandleFunc("/", indexHandler)
	// http.Handle("/favicon.ico", http.NotFoundHandler())
	// http.ListenAndServe(":8080", nil)
	Info.Println(debugging.ThisFunc())
	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/dealNewHand", dealNewHandPostHandler).Methods("POST")
	r.HandleFunc("/NextStep", executeNextStepPostHandler).Methods("POST")
	r.HandleFunc("/postBet", postBetPostHandler).Methods("POST")
	r.HandleFunc("/fullReset", fullResetPostHandler).Methods("POST")

	r.Handle("/favicon.ico", http.NotFoundHandler())

	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/",
		http.FileServer(http.Dir("web/styles/"))))

	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/",
		http.FileServer(http.Dir("web/images/"))))

	http.Handle("/", r)
	log.Fatalln(http.ListenAndServe(":8080", nil))

}

// #################################################################

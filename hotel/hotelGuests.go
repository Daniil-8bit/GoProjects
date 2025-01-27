package hotel

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func StartServer() {
	http.HandleFunc("/guests", pageHandler)
	http.HandleFunc("/guests/new", addGuestHanndler)
	http.HandleFunc("/guests/create", creatGuestHandler)
	err := http.ListenAndServe(":2323", nil)
	log.Fatal(err)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {

	clientsList := getStrings("C://Go//go1.23.4//src//github.com//Daniil-8bit//GoProjects//hotel//clients.txt")
	res, err := template.ParseFiles("C://Go//go1.23.4//src//github.com//Daniil-8bit//GoProjects//hotel//main.html")

	guests := GuestBook{
		Amount: len(clientsList),
		Person: clientsList,
	}

	err = res.Execute(w, guests)
	checkError(err)

	//respText := []byte("List of clients") // byte text for output
	//_, err := w.Write(respText)
	//fmt.Printf("%#v\n", clientsList)
	/*checkError(err)
	err = res.Execute(w, nil) //output template data
	checkError(err)*/
	//guestsTmpl := "<div>\n<p>Кол-во гостей:</p>\n<p>{{.Amount}}</p>\n</div>\n<div>\n<p>Гости:</p>\n{{range .Person}}\n<p>{{.}}</p>\n{{end}}\n</div>" // template
	//executeTemplate(guestsTmpl, guests, w) // func for templates
	//fmt.Println(guests)
	//executeTemplate("Client: {{range .}}{{.}}<br>{{end}}", clientsList, w)
}

func addGuestHanndler(w http.ResponseWriter, r *http.Request) {
	page, err := template.ParseFiles("C://Go//go1.23.4//src//github.com//Daniil-8bit//GoProjects//hotel//addGuest.html")
	checkError(err)

	err = page.Execute(w, nil)
	checkError(err)
}

func creatGuestHandler(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("guestName")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("C://Go//go1.23.4//src//github.com//Daniil-8bit//GoProjects//hotel//clients.txt", options, os.FileMode(0600))
	checkError(err)

	_, err = fmt.Fprintln(file, val)
	checkError(err)

	err = file.Close()
	checkError(err)

	//_, err = w.Write([]byte("Гость " + val + " был добавлен!"))
	//checkError(err)
	http.Redirect(w, r, "/guests", http.StatusFound)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	checkError(err)
	defer file.Close()

	var line []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	checkError(scanner.Err())
	return line
}

func executeTemplate(text string, data interface{}, w http.ResponseWriter) {
	tmpl, err := template.New("test").Parse(text)
	checkError(err)
	err = tmpl.Execute(w, data)
	checkError(err)
}

type GuestBook struct {
	Amount int
	Person []string
}

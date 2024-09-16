package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Sum Calculator</title>
</head>
<body>
    <h1>Sum Calculator</h1>
    <form method="POST">
        <label for="num1">Number 1:</label>
        <input type="text" id="num1" name="num1"><br><br>
        <label for="num2">Number 2:</label>
        <input type="text" id="num2" name="num2"><br><br>
        <input type="submit" value="Calculate Sum">
    </form>
    {{if .}}
    <h2>Sum: {{.}}</h2>
    {{end}}
</body>
</html>
`))

func sumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		num1, err1 := strconv.Atoi(r.FormValue("num1"))
		num2, err2 := strconv.Atoi(r.FormValue("num2"))
		if err1 == nil && err2 == nil {
			sum := num1 + num2
			tmpl.Execute(w, sum)
			return
		}
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", sumHandler)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

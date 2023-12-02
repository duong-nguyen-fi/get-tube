package main

import (
	"fmt"
	"get-tube/pkg"
	"html/template"
	"log"
	//"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

var secretKey = "mysecret" // Replace with your secret key
var scriptPath = "./download-video.sh"

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/download", downloadHandler)

	fmt.Println("Server is running on :8091")
	http.ListenAndServe(":8091", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "index.html", nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		url := r.FormValue("url")
		secret := r.FormValue("secret")

		if secret != secretKey {
			http.Error(w, "Invalid secret key", http.StatusUnauthorized)
			return
		}
		fmt.Println("secret=" + secret)
		fmt.Println("url=" + url)
		id, err := pkg.CheckParameters(url)
		if err != nil {
			http.Error(w, "Error your url is not in correct format", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		fmt.Println("id=" + id)
		// Trigger the Bash script with the URL as a parameter
		cmd := exec.Command(scriptPath, id)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		err = cmd.Run()
		if err != nil {
			http.Error(w, "Error running the script", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		// After the script finishes, redirect to the download page
		http.Redirect(w, r, "/download", http.StatusSeeOther)
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// Assuming the script generates a file named "output.txt"
	filename := "audio.wav"

	filePath := "/files/" + filename
	// Set the content type header
	w.Header().Set("Content-Type", "application/octet-stream")
	// Set the content disposition header to trigger a download
	w.Header().Set("Content-Disposition", "attachment; filename="+filePath)

	// Serve the file directly
	http.ServeFile(w, r, filePath)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := "templates/" + tmpl
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func info(text string) {
	log.Printf("INFO: " + text)
}

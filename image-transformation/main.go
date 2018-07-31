package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/MohitArora1/gophercises/image-transformation/transform"
)

// it handle the all request at "/" show the html page from which user able to
// upload the image
func index(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
		<head>
		<title>Image Transform Service</title>
			<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.2/css/bootstrap.min.css" integrity="sha384-Smlep5jCw/wG7hdkwQ/Z5nLIefveQRIY9nfy6xoR1uRYBtpZgI6339F5dgvm/e9B" crossorigin="anonymous">
			
			<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js" integrity="sha384-ZMP7rVo3mIykV+2+9J3UJ46jBk0WLaUAdn689aCwoqbBJiSnjAK/l8WvCWPIPm49" crossorigin="anonymous"></script>
			<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.2/js/bootstrap.min.js" integrity="sha384-o+RDsa0aLu++PJvFqy8fFScvbHFLtbvScb8AjopnFD+iEQ7wo/CG0xlczd+2O/em" crossorigin="anonymous"></script>
		</head>
		<body>
			<div class="col-md-6 offset-md-3 border">
			<div class="col-md-10 offset-md-1">
				<form action="/upload" method="post" enctype="multipart/form-data">
					<div class="form-group">
						<label for="image">Choose Image</label>
						<input class="form-control" type="file" name="image" id="image">
					</div>
					<button type="submit" class="btn btn-default">Upload</button>
				</form>
			</div>
			</div>
		</body>
	</html`
	fmt.Fprint(w, html)

}

// uploadHandle is handle the request at "/upload" it basically save the file for persistance
// and redirect to modify url
func uploadHandle(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()
	ext := filepath.Ext(header.Filename)[1:]
	saveFile, err := tempfile("", ext)
	io.Copy(saveFile, file)
	http.Redirect(w, r, "modify/"+filepath.Base(saveFile.Name()), http.StatusFound)

}

// modifyHandle is used for the show all the transformed image based on the query params
// if mode is not defined in the url it will show four different image with different images
// if mode is defined and number of shapes is not defined it will show the 3 images with the same mode
// but different number of shapes
// if mode and number of shapes is in url the it will show one image
func modifyHandle(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./img/" + filepath.Base(r.URL.Path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	ext := filepath.Ext(file.Name())[1:]
	mode := r.FormValue("mode")
	if mode == "" {
		renderAllMode(w, r, file, ext)
		return
	}
	number := r.FormValue("number")
	if number == "" {
		renderSingleMode(w, r, file, ext, mode)
		return
	}
	http.Redirect(w, r, "/img/"+filepath.Base(file.Name()), http.StatusFound)
}

// this function will generate three image with the same mode
// but different number of shapes
func renderSingleMode(w http.ResponseWriter, r *http.Request, file io.ReadSeeker, ext, mode string) {
	a, err := genrateImage(file, ext, mode, "50")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file.Seek(0, 0)
	b, err := genrateImage(file, ext, mode, "100")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file.Seek(0, 0)
	c, err := genrateImage(file, ext, mode, "150")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	html := `<html><body>
		{{range .}}
			<a href="/modify/{{.Name}}?mode={{.Mode}}&number={{.Number}}">
			<img style="width: 20%;" src="/{{.Name}}">
			</a>
		{{end}}
		</body></html>`
	tpl := template.Must(template.New("").Parse(html))
	type Images struct {
		Name   string
		Mode   int
		Number int
	}
	images := []Images{
		{a, 2, 50}, {b, 2, 100}, {c, 2, 150},
	}

	tpl.Execute(w, images)

}

// this funcion generate the four different images with differnt modes
func renderAllMode(w http.ResponseWriter, r *http.Request, file io.ReadSeeker, ext string) {
	a, err := genrateImage(file, ext, "2", "30")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file.Seek(0, 0)
	b, err := genrateImage(file, ext, "3", "30")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file.Seek(0, 0)
	c, err := genrateImage(file, ext, "4", "30")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file.Seek(0, 0)
	d, err := genrateImage(file, ext, "5", "30")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file.Seek(0, 0)
	html := `<html><body>
			{{range .}}
				<a href="/modify/{{.Name}}?mode={{.Mode}}">
				<img style="width: 20%;" src="/{{.Name}}">
				</a>
			{{end}}
			</body></html>`
	tpl := template.Must(template.New("").Parse(html))
	type Images struct {
		Name string
		Mode int
	}
	images := []Images{
		{a, 2}, {b, 3}, {c, 4}, {d, 5},
	}

	tpl.Execute(w, images)
}

// this function call the transform function from transform package
// to get trasform image
func genrateImage(file io.Reader, ext, mode, number string) (string, error) {
	out, err := transform.Transform(file, ext, mode, number)
	outFile, err := tempfile("", ext)
	if err != nil {
		return "", err
	}
	io.Copy(outFile, out)
	return outFile.Name(), nil
}

// this function generate the temp file for us
func tempfile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("./img/", prefix)
	if err != nil {
		return nil, errors.New("primitive: failed to create temporary file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}

// getHandlers will return the router mux with handlers
func getHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/upload", uploadHandle)
	fs := http.FileServer(http.Dir("./img/"))
	mux.Handle("/img/", http.StripPrefix("/img", fs))
	mux.HandleFunc("/modify/", modifyHandle)
	return mux
}

// this is main function we defined all the router here
func main() {
	http.ListenAndServe(":8000", getHandlers())
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime/debug"
	"strconv"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func debugHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	lineStr := r.FormValue("line")
	line, err := strconv.Atoi(lineStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := bytes.NewBuffer(nil)
	io.Copy(b, file)
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}
	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, b.String())
	style := styles.Get("monokai")
	formatter := html.New(html.TabWidth(2), html.WithLineNumbers(), html.HighlightLines(lines))
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>")
	formatter.Format(w, style, iterator)
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	funcThatPanic()
}

func funcThatPanic() {
	panic("ho no!")
}

func recoveryMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, makeLinks(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

func makeLinks(stack string) string {
	re := regexp.MustCompile("(\t.*:[0-9]*)")
	lines := re.FindAllString(stack, -1)
	for _, line := range lines {
		regexSplit := regexp.MustCompile(":")
		splits := regexSplit.Split(line, -1)
		link := "<a href='/debug?path=" + splits[0] + "&line=" + splits[1] + "'>" + line + "</a>"
		reg := regexp.MustCompile(line)
		stack = reg.ReplaceAllString(stack, link)
	}
	return stack
}

func gethandles() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug", debugHandler)
	mux.HandleFunc("/panic", panicHandler)
	return mux
}

func main() {
	http.ListenAndServe(":8000", recoveryMw(gethandles()))
}

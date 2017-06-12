package httpServer

import (
	"context"
	"fmt"
	"github.com/acmacalister/skittles"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/synw/microb-dashboard/types"
	"github.com/synw/microb/libmicrob/events"
	"github.com/synw/terr"
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var dir = getDir()

func getDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	d := fmt.Sprintf("%s", path.Dir(filename))
	d = strings.Replace(d, "/httpServer", "", -1)
	return d
}

func getTemplate(name string) string {
	t := dir + "/templates/" + name + ".html"
	return t
}

var View = template.Must(template.New("index.html").ParseFiles(getTemplate("index"), getTemplate("head"), getTemplate("header"), getTemplate("navbar"), getTemplate("footer")))
var V404 = template.Must(template.New("404.html").ParseFiles(getTemplate("404"), getTemplate("head"), getTemplate("header"), getTemplate("navbar"), getTemplate("footer")))
var V500 = template.Must(template.New("500.html").ParseFiles(getTemplate("500"), getTemplate("head"), getTemplate("header"), getTemplate("navbar"), getTemplate("footer")))

type httpResponseWriter struct {
	http.ResponseWriter
	status *int
}

func InitHttpServer(server *types.DashboardServer, serve bool) {
	// routing
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	// main route
	r.Route("/", func(r chi.Router) {
		r.Get("/*", serveRequest)
	})
	fmt.Println("ADDR", server.Addr)

	// init
	httpServer := &http.Server{
		Addr:         server.Addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
	}
	server.Instance = httpServer
	// static
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	r.FileServer("/static", http.Dir(filesDir))
	// run
	if serve == true {
		Run(server)
	}
}

func Run(server *types.DashboardServer) {
	events.New("state", "http", "httpServer.Run", startMsg(server), nil)
	server.Running = true
	server.Instance.ListenAndServe()
}

func Stop(server *types.DashboardServer) *terr.Trace {
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	srv := server.Instance
	err := srv.Shutdown(ctx)
	if err != nil {
		tr := terr.New("httpServer.Stop", err)
		events.New("error", "http", "httpServer.Stop", stopMsg(), nil)
		return tr
	}
	server.Running = false
	events.New("state", "http", "httpServer.Stop", stopMsg(), nil)
	return nil
}

func renderTemplate(response http.ResponseWriter, page *types.Page) {
	err := View.Execute(response, page)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func render404(response http.ResponseWriter, page *types.Page) {
	err := V404.Execute(response, page)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func render500(response http.ResponseWriter, page *types.Page) {
	err := V500.Execute(response, page)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func serveRequest(response http.ResponseWriter, request *http.Request) {
	url := request.URL.Path
	status := http.StatusOK
	page := &types.Page{Url: url, Title: "Microb", Content: ""}
	response = httpResponseWriter{response, &status}
	renderTemplate(response, page)
}

func stopMsg() string {
	msg := "Dashboard server stopped"
	return msg
}

func startMsg(server *types.DashboardServer) string {
	var msg string
	msg = "Dashboard server started at " + server.Addr + " for domain " + skittles.BoldWhite(server.Domain)
	return msg
}

package notes

import (
	"embed"
	"html/template"
	"net/http"
	"strconv"

	"github.com/simonkana/go-notes/internal/db"
)

//go:embed templates/*.html
var templatesFS embed.FS

type handler struct {
	q    *db.Queries
	tmpl *template.Template
}

func NewHandler(q *db.Queries) http.Handler {
	tmpl := template.Must(template.ParseFS(templatesFS, "templates/*.html"))
	h := &handler{q: q, tmpl: tmpl}

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.index)
	mux.HandleFunc("/create", h.create)
	mux.HandleFunc("/delete", h.delete)
	return mux
}

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	notes, _ := h.q.ListNotes(r.Context())
	h.tmpl.ExecuteTemplate(w, "index.html", notes)
}

func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	h.q.CreateNote(r.Context(), db.CreateNoteParams{Title: title, Content: content})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	h.q.DeleteNote(r.Context(), int64(id))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}


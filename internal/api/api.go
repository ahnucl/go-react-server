package api

import (
	"net/http"

	"github.com/ahnucl/go-react-server/internal/store/pgstore"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
	}
	r := chi.NewRouter() // mesma coisa que NewMux()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOken"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/subscribe/{room_id}", a.handleSubscribe)

	r.Route("/api", func(r chi.Router) {
		r.Route("/rooms", func(r chi.Router) {
			r.Post("/", a.handleCreateRoom)
			r.Get("/", a.handleGetRooms)

			r.Route("/{room_id}/messages", func(r chi.Router) { // Parâmetro de "path" ou "caminho" (eu chamo de "rota")
				r.Post("/", a.handleCreateRoomMessage)
				r.Get("/", a.handleGetRoomMessages)

				r.Route("/{message_id}", func(r chi.Router) {
					r.Get("/", a.handleGetRoomMessage)
					r.Patch("/react", a.handleReactionToMessage)
					r.Delete("/react", a.handleRemoveReactionToMessage)
					r.Patch("/answer", a.handleMarkMessageAsAnswered)
				})
			})
		})
	})

	a.r = r
	return a
}

func (h apiHandler) handleSubscribe(w http.ResponseWriter, r *http.Request)               {}
func (h apiHandler) handleCreateRoom(w http.ResponseWriter, r *http.Request)              {}
func (h apiHandler) handleGetRooms(w http.ResponseWriter, r *http.Request)                {}
func (h apiHandler) handleCreateRoomMessage(w http.ResponseWriter, r *http.Request)       {}
func (h apiHandler) handleGetRoomMessages(w http.ResponseWriter, r *http.Request)         {}
func (h apiHandler) handleGetRoomMessage(w http.ResponseWriter, r *http.Request)          {}
func (h apiHandler) handleReactionToMessage(w http.ResponseWriter, r *http.Request)       {}
func (h apiHandler) handleRemoveReactionToMessage(w http.ResponseWriter, r *http.Request) {}
func (h apiHandler) handleMarkMessageAsAnswered(w http.ResponseWriter, r *http.Request)   {}

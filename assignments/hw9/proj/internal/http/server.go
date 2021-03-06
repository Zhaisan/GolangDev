package http

import 	(
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"lectures-6/internal/cache"
	"lectures-6/internal/models"
	"lectures-6/internal/store"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}
	store       store.Store
	cache 		cache.Cache

	Address string
}

func NewServer(ctx context.Context, opts ...ServerOption) *Server {
	srv := &Server{
		ctx:         ctx,
		idleConnsCh: make(chan struct{}),
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	laptopsResource := NewLaptopResource(s.store, s.cache)
	r.Mount("/electronics/laptops", laptopsResource.Routes())
	// -------

	r.Post("/sport-hobby/snowboards", func(w http.ResponseWriter, r *http.Request) {
		snowboard := new(models.Snowboard)
		if err := json.NewDecoder(r.Body).Decode(snowboard); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		fmt.Println("Status OK")

		if err := s.store.Snowboards().Create(r.Context(), snowboard); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
	r.Get("/sport-hobby/snowboards", func(w http.ResponseWriter, r *http.Request) {
		snowboards, err := s.store.Snowboards().All(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, snowboards)
	})
	r.Get("/sport-hobby/snowboards/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		snowboard, err := s.store.Snowboards().ByID(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, snowboard)
	})
	r.Put("/sport-hobby/snowboards", func(w http.ResponseWriter, r *http.Request) {
		snowboard := new(models.Snowboard)
		if err := json.NewDecoder(r.Body).Decode(snowboard); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		if err := s.store.Snowboards().Update(r.Context(), snowboard); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
			return
		}
	})
	r.Delete("/sport-hobby/snowboards/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		if err := s.store.Snowboards().Delete(r.Context(), id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
			return
		}

	})

	// -------

	r.Post("/fashion-style/shirts", func(w http.ResponseWriter, r *http.Request) {
		shirt := new(models.Shirt)
		if err := json.NewDecoder(r.Body).Decode(shirt); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		log.Println("Status OK")

		if err := s.store.Shirts().Create(r.Context(), shirt); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
	r.Get("/fashion-style/shirts", func(w http.ResponseWriter, r *http.Request) {
		shirts, err := s.store.Shirts().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, shirts)
	})
	r.Get("/fashion-style/shirts/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		shirt, err := s.store.Shirts().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, shirt)
	})
	r.Put("/fashion-style/shirts", func(w http.ResponseWriter, r *http.Request) {
		shirt := new(models.Shirt)
		if err := json.NewDecoder(r.Body).Decode(shirt); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		if err := s.store.Shirts().Update(r.Context(), shirt); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
			return
		}
	})
	r.Delete("/fashion-style/shirts/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		if err := s.store.Shirts().Delete(r.Context(), id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
			return
		}
	})

	// -------

	r.Post("/for-kids/toys", func(w http.ResponseWriter, r *http.Request) {
		toy := new(models.Toy)
		if err := json.NewDecoder(r.Body).Decode(toy); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		log.Println("Status OK")

		if err := s.store.Toys().Create(r.Context(), toy); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
	r.Get("/for-kids/toys", func(w http.ResponseWriter, r *http.Request) {

		toys, err := s.store.Shirts().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, toys)
	})
	r.Get("/for-kids/toys/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		toy, err := s.store.Toys().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, toy)
	})
	r.Put("/for-kids/toys", func(w http.ResponseWriter, r *http.Request) {
		toy := new(models.Toy)
		if err := json.NewDecoder(r.Body).Decode(toy); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		if err := s.store.Toys().Update(r.Context(), toy); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
			return
		}
	})
	r.Delete("/for-kids/toys/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		if err := s.store.Toys().Delete(r.Context(), id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DB err: %v", err)
			return
		}
	})

	// ----

	r.Post("/registration", func(w http.ResponseWriter, r *http.Request) {
		user := new(models.User)
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		fmt.Println("Status OK")



		if err := s.store.Users().Create(r.Context(), user); err != nil{
			fmt.Fprintf(w, "Unknown err: %v", err)
		}
	})

	r.Get("/registration", func(w http.ResponseWriter, r *http.Request) {
		users, err := s.store.Users().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, users)
	})
	

	return r
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         s.Address,
		Handler:      s.basicHandler(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 30,
	}
	go s.ListenCtxForGT(srv)

	log.Println("[HTTP] Server running on", s.Address)
	return srv.ListenAndServe()
}

func (s *Server) ListenCtxForGT(srv *http.Server) {
	<-s.ctx.Done() // ??????????????????????, ???????? ???????????????? ???????????????????? ???? ??????????????

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down^ %v", err)
	}

	log.Println("[HTTP] Proccessed all idle connections")
	close(s.idleConnsCh)
}

func (s *Server) WaitForGracefulTermination() {
	// ???????? ???? ???????????? ?????? ???????????????? ????????????
	<-s.idleConnsCh
}

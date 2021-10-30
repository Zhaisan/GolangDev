package http

import (
	"context"
	"encoding/json"
	"fmt"
	"lectures-6/internal/models"
	"lectures-6/internal/store"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}
	store       store.Store

	Address string
}

func NewServer(ctx context.Context, address string, store store.Store) *Server {
	return &Server{
		ctx:         ctx,
		idleConnsCh: make(chan struct{}),
		store:       store,

		Address: address,
	}
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	
	r.Post("/electronics/laptops", func(w http.ResponseWriter, r *http.Request) {
		laptop := new(models.Laptop)
		if err := json.NewDecoder(r.Body).Decode(laptop); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		fmt.Println("Status OK")

		s.store.Laptops().Create(r.Context(), laptop)
	})
	r.Get("/electronics/laptops", func(w http.ResponseWriter, r *http.Request) {
		laptops, err := s.store.Laptops().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, laptops)
	})
	r.Get("/electronics/laptops/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		laptop, err := s.store.Laptops().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, laptop)
	})
	r.Put("/electronics/laptops", func(w http.ResponseWriter, r *http.Request) {
		laptop := new(models.Laptop)
		if err := json.NewDecoder(r.Body).Decode(laptop); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Laptops().Update(r.Context(), laptop)
	})
	r.Delete("/electronics/laptops/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Laptops().Delete(r.Context(), id)
	})


	// -------

	r.Post("/sport-hobby/snowboards", func(w http.ResponseWriter, r *http.Request) {
		snowboard := new(models.Snowboard)
		if err := json.NewDecoder(r.Body).Decode(snowboard); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		fmt.Println("Status OK")

		s.store.Snowboards().Create(r.Context(), snowboard)
	})
	r.Get("/sport-hobby/snowboards", func(w http.ResponseWriter, r *http.Request) {
		snowboards, err := s.store.Snowboards().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, snowboards)
	})
	r.Get("/sport-hobby/snowboards/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		snowboard, err := s.store.Snowboards().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, snowboard)
	})
	r.Put("/sport-hobby/snowboards", func(w http.ResponseWriter, r *http.Request) {
		snowboard := new(models.Snowboard)
		if err := json.NewDecoder(r.Body).Decode(snowboard); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Snowboards().Update(r.Context(), snowboard)
	})
	r.Delete("/sport-hobby/snowboards/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Snowboards().Delete(r.Context(), id)
	})

	// -------

	r.Post("/fashion-style/shirts", func(w http.ResponseWriter, r *http.Request) {
		shirt := new(models.Shirt)
		if err := json.NewDecoder(r.Body).Decode(shirt); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		log.Println("Status OK")

		s.store.Shirts().Create(r.Context(), shirt)
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

		s.store.Shirts().Update(r.Context(), shirt)
	})
	r.Delete("/fashion-style/shirts/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Shirts().Delete(r.Context(), id)
	})

	// -------

	r.Post("/for-kids/toys", func(w http.ResponseWriter, r *http.Request) {
		shirt := new(models.Shirt)
		if err := json.NewDecoder(r.Body).Decode(shirt); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		log.Println("Status OK")

		s.store.Shirts().Create(r.Context(), shirt)
	})
	r.Get("/for-kids/toys", func(w http.ResponseWriter, r *http.Request) {
		shirts, err := s.store.Shirts().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, shirts)
	})
	r.Get("/for-kids/toys/{id}", func(w http.ResponseWriter, r *http.Request) {
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
	r.Put("/for-kids/toys", func(w http.ResponseWriter, r *http.Request) {
		shirt := new(models.Shirt)
		if err := json.NewDecoder(r.Body).Decode(shirt); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Shirts().Update(r.Context(), shirt)
	})
	r.Delete("/for-kids/toys/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Shirts().Delete(r.Context(), id)
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
	<-s.ctx.Done() // блокируемся, пока контекст приложения не отменен

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down^ %v", err)
	}

	log.Println("[HTTP] Proccessed all idle connections")
	close(s.idleConnsCh)
}

func (s *Server) WaitForGracefulTermination() {
	// блок до записи или закрытия канала
	<-s.idleConnsCh
}

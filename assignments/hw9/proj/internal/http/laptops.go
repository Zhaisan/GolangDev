package http

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"lectures-6/internal/cache"
	"lectures-6/internal/models"
	"lectures-6/internal/store"
	"net/http"
	"strconv"
)

type LaptopResource struct {
	store store.Store
	cache cache.Cache
}

func NewLaptopResource(store store.Store, cache cache.Cache) *LaptopResource {
	return &LaptopResource{
		store: store,
		cache: cache,
	}
}

func (lr *LaptopResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", lr.CreateLaptop)
	r.Get("/", lr.AllLaptops)
	r.Get("/{id}", lr.ByID)
	r.Put("/", lr.UpdateLaptop)
	r.Delete("/{id}", lr.DeleteLaptop)

	return r
}

func (lr *LaptopResource) CreateLaptop(w http.ResponseWriter, r *http.Request) {
	laptop := new(models.Laptop)
	if err := json.NewDecoder(r.Body).Decode(laptop); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	fmt.Println("Status OK")

	if err := lr.store.Laptops().Create(r.Context(), laptop); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}


func (lr *LaptopResource) AllLaptops(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	filter := &models.LaptopsFilter{}

	searchQuery := queryValues.Get("query")
	if searchQuery != "" {
		laptopsFromCache, err := lr.cache.Laptops().Get(r.Context(), searchQuery)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Cache err: %v", err)
			return
		}

		if laptopsFromCache != nil {
			render.JSON(w, r, laptopsFromCache)
			return
		}

		filter.Query = &searchQuery
	}

	laptops, err := lr.store.Laptops().All(r.Context(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if searchQuery != "" && len(laptops) > 0{
		err = lr.cache.Laptops().Set(r.Context(), searchQuery, laptops)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Cache err: %v", err)
			return
		}
	}

	render.JSON(w, r, laptops)
}

func (lr *LaptopResource) ByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}


	laptop, err := lr.store.Laptops().ByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	render.JSON(w, r, laptop)
}

func (lr *LaptopResource) UpdateLaptop(w http.ResponseWriter, r *http.Request) {
	laptop := new(models.Laptop)
	if err := json.NewDecoder(r.Body).Decode(laptop); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := lr.store.Laptops().Update(r.Context(), laptop); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

}

func (lr *LaptopResource) DeleteLaptop(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := lr.store.Laptops().Delete(r.Context(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

}
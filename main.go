package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/blevesearch/bleve/v2"
	"github.com/go-chi/chi/v5"
)


type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}


var index bleve.Index
var productMap = make(map[string]Product)


var categories = []string{
	"Electronics", "Footwear", "Home & Kitchen", "Fitness", "Books", "Clothing",
}


func generateProducts(n int) []Product {
	rand.Seed(time.Now().UnixNano())
	products := make([]Product, n)

	
	products = append(products, Product{
		ID:       1,
		Name:     "Wireless Mouse",
		Category: "Electronics",
	})
	products = append(products, Product{
		ID:       2,
		Name:     "Gaming Keyboard",
		Category: "Electronics",
	})

	for i := 3; i <= n; i++ {
		products = append(products, Product{
			ID:       i,
			Name:     fmt.Sprintf("Product %d", i),
			Category: categories[rand.Intn(len(categories))],
		})
	}

	return products
}

func createIndex(products []Product) {
	mapping := bleve.NewIndexMapping()
	var err error
	index, err = bleve.NewMemOnly(mapping)
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	for _, product := range products {
		docID := strconv.Itoa(product.ID)
		err := index.Index(docID, product)
		if err != nil {
			log.Printf("Error indexing product %d: %v", product.ID, err)
		}
		productMap[docID] = product
	}
}

func searchIndex(query string) []Product {
	q := bleve.NewMatchQuery(query)
	searchRequest := bleve.NewSearchRequestOptions(q, 50, 0, false)
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		log.Printf("Search error: %v", err)
		return nil
	}

	var results []Product
	for _, hit := range searchResult.Hits {
		if product, exists := productMap[hit.ID]; exists {
			results = append(results, product)
		}
	}
	return results
}


func main() {
	log.Println("Generating 1 million product records...")
	products := generateProducts(1000000) 
	log.Println("Indexing products...")
	createIndex(products)
	log.Println("Indexing complete!")

	r := chi.NewRouter()

	
	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			http.Error(w, "Missing query parameter 'q'", http.StatusBadRequest)
			return
		}
		results := searchIndex(query)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	
	go func() {
		log.Println("Server is running on http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown signal received...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %v", err)
	}
	log.Println("Server stopped gracefully")
}

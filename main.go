package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(writer, request)
	})
}

func (cfg *apiConfig) serverMetrics(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.WriteHeader(200)
	writer.Write([]byte(fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`, cfg.fileserverHits.Load())))
}

func (cfg *apiConfig) resetMetrics(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	cfg.fileserverHits.Store(0)
}

func healthz(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(200)
	writer.Write([]byte("OK"))
}

func profanityFilter(input string) (cleaned string) {
	profanity := []string{"kerfuffle", "sharbert", "fornax"}

	words := strings.Fields(input)
	cleanedWords := make([]string, len(words))
	copy(cleanedWords, words)
	for i, word := range words {
		for _, badWord := range profanity {
			if strings.ToLower(word) == badWord {
				cleanedWords[i] = strings.Repeat("*", 4)
			}
		}
	}

	return strings.Join(cleanedWords, " ")
}

func validateChirp(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	type validateRequest struct {
		Message string `json:"body"`
	}

	type validateResponse struct {
		Error       string `json:"error"`
		Valid       bool   `json:"valid"`
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(request.Body)
	req := validateRequest{}
	err := decoder.Decode(&req)
	if err != nil {
		errString := fmt.Sprintf("Failed to decode json request body: %s", err)
		log.Print(errString)
		writer.WriteHeader(400)
		response := validateResponse{Error: errString, Valid: false}
		dat, err := json.Marshal(response)
		if err != nil {
			log.Printf("Failed to marshal json response body")
		} else {
			writer.Write(dat)
		}
		return
	}

	messageLength := len(req.Message)
	if messageLength > 140 {
		errString := fmt.Sprintf("Chirp was too long, message length %d greater than 140 characters", messageLength)
		log.Print(errString)
		writer.WriteHeader(400)
		response := validateResponse{Error: errString, Valid: false}
		dat, err := json.Marshal(response)
		if err != nil {
			log.Printf("Failed to marshal json response body")
		} else {
			writer.Write(dat)
		}
		return
	}

	writer.WriteHeader(200)
	cleanedMessage := profanityFilter(req.Message)
	response := validateResponse{Error: "", Valid: true, CleanedBody: cleanedMessage}
	dat, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to marshal json response body")
	} else {
		writer.Write(dat)
	}
}

func main() {
	serveMux := http.NewServeMux()
	apiCfg := apiConfig{fileserverHits: atomic.Int32{}}
	serveMux.Handle("/app/", http.StripPrefix("/app", apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(".")))))
	serveMux.HandleFunc("GET /api/healthz", healthz)
	serveMux.HandleFunc("POST /api/validate_chirp", validateChirp)
	serveMux.HandleFunc("GET /admin/metrics", apiCfg.serverMetrics)
	serveMux.HandleFunc("POST /admin/reset", apiCfg.resetMetrics)

	server := http.Server{}
	server.Addr = ":8080"
	server.Handler = serveMux

	server.ListenAndServe()
}

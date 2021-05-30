package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"marcos979/go-totenizator/model"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type TokenizeParam struct {
	Url string
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("api-key")
	if apiKey == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenizeParamBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenizeParam := TokenizeParam{}
	err = json.Unmarshal(tokenizeParamBytes, &tokenizeParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := model.GenerateTokenizedUrl(tokenizeParam.Url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(token))
}

func middlewareTimeHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("\nHandler elapsed time: %s", time.Since(start))
	})
}

func main() {
	godotenv.Load()

	tokenizeHandlerFunc := http.HandlerFunc(tokenizeHandler)
	http.Handle("/", middlewareTimeHandler(tokenizeHandlerFunc))

	fmt.Printf("Process running: %d\n", os.Getpid())
	log.Fatal(http.ListenAndServe(":3000", nil))
}

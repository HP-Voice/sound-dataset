package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func initApi() error {
	http.HandleFunc("/labels", cors(getLabelsHandler))
	http.HandleFunc("/random-spell", cors(getRandomSpellHandler))
	http.HandleFunc("/sample", cors(postSampleHandler))
	http.HandleFunc("/sentence", cors(getSentenceHandler))
	return http.ListenAndServe(config.Api.Address, nil)
}

func getLabelsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	labels, err := getLabels()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(labels)
}

func getRandomSpellHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	labelId, err := getRandomSpell()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(labelId)
}

func postSampleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	labelId, err := strconv.Atoi(r.FormValue("Label-Id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	f, _, err := r.FormFile("Sample")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer f.Close()

	_, err = writeSample(labelId, f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
}

func getSentenceHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:5000/sentence")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(bytes)
}

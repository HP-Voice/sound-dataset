package main

import (
	"encoding/json"
	"github.com/jackc/pgx/pgtype"
	"net/http"
	"strconv"
)

func initApi() error {
	http.HandleFunc("/labels", cors(getLabelsHandler))
	http.HandleFunc("/random-spell", cors(getRandomSpellHandler))
	http.HandleFunc("/sample", cors(postSampleHandler))
	http.HandleFunc("/sentence", cors(getSentenceHandler))
	http.HandleFunc("/admin/stats", cors(admin(getStatsHandler)))
	http.HandleFunc("/admin/sample-for-approval", cors(admin(getSampleForApprovalHandler)))
	http.HandleFunc("/admin/verdict", cors(admin(postVerdictHandler)))
	if config.Tls == nil {
		return http.ListenAndServe(config.Api.Address, nil)
	} else {
		return http.ListenAndServeTLS(config.Api.Address, config.Tls.Cert, config.Tls.Key, nil)
	}
}

func getLabelsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	labels, err := readLabels()
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

	labelId, err := getRandomSpellId()
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
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	bytes, err := getSentence()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write(bytes)
}

func getStatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	stats, err := readStats()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(stats)
}

func getSampleForApprovalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sample, err := readSampleForApproval()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_ = json.NewEncoder(w).Encode(sample)
}

func postVerdictHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := &struct {
		SampleId pgtype.UUID `json:"sampleId"`
		Verdict  int         `json:"verdict"`
	}{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = writeVerdict(request.SampleId, request.Verdict)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}

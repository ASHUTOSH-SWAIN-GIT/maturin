package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ASHUTOSH-SWAIN-GIT/maturin/database"
	"github.com/ASHUTOSH-SWAIN-GIT/maturin/security"

	"github.com/go-chi/chi/v5"
)

type BucketHandler struct {
	queries *database.Queries
}

// NewBucketHandler now accepts *database.Queries
func NewBucketHandler(queries *database.Queries) *BucketHandler {
	return &BucketHandler{queries: queries}
}

type CreateBucketRequest struct {
	Name      string `json:"name"`
	AccountID string `json:"account_id"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

func (h *BucketHandler) RegisterBucket(w http.ResponseWriter, r *http.Request) {
	var req CreateBucketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 1. Encrypt the keys before saving
	encAccess, err := security.Encrypt(req.AccessKey)
	if err != nil {
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		return
	}
	encSecret, err := security.Encrypt(req.SecretKey)
	if err != nil {
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		return
	}

	// Use the generated CreateBucket method
	params := database.CreateBucketParams{
		Name:      req.Name,
		AccountID: req.AccountID,
		AccessKey: encAccess,
		SecretKey: encSecret,
	}

	_, err = h.queries.CreateBucket(context.Background(), params)
	if err != nil {
		http.Error(w, "Failed to register bucket: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "created"})
}

func (h *BucketHandler) ListBuckets(w http.ResponseWriter, r *http.Request) {
	// Use generated ListBuckets method
	buckets, err := h.queries.ListBuckets(context.Background())
	if err != nil {
		http.Error(w, "Failed to fetch buckets", http.StatusInternalServerError)
		return
	}

	type BucketResponse struct {
		ID        int32  `json:"id"` // sqlc uses int32 for SERIAL/INTEGER
		Name      string `json:"name"`
		AccountID string `json:"account_id"`
	}

	var response []BucketResponse
	for _, b := range buckets {
		response = append(response, BucketResponse{
			ID:        b.ID,
			Name:      b.Name,
			AccountID: b.AccountID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *BucketHandler) TriggerScan(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr) // sqlc ID is int32, usually fits in int
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Verify bucket exists
	_, err = h.queries.GetBucketByID(context.Background(), int32(id))
	if err != nil {
		http.Error(w, "Bucket not found", http.StatusNotFound)
		return
	}

	// TODO: Trigger actual scan logic here

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "scan_started"})
}

func (h *BucketHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	snapshots, err := h.queries.ListSnapshots(context.Background(), int32(id))
	if err != nil {
		http.Error(w, "Failed to fetch stats", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snapshots)
}

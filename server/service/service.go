package service

import (
	"context"
	"fmt"

	"github.com/ASHUTOSH-SWAIN-GIT/maturin/database"
	"github.com/ASHUTOSH-SWAIN-GIT/maturin/security"
)

// handles the interaction with R2/S3
type Scanner struct {
	queries *database.Queries
}

func NewScanner(q *database.Queries) *Scanner {
	return &Scanner{queries: q}
}

func (s *Scanner) RegisterNewBucket(ctx context.Context, name, accountID, accessKey, secretKey string) (database.Bucket, error) {
	//encrypt the sensitive keys
	encAccess, err := security.Encrypt(accessKey)
	if err != nil {
		return database.Bucket{}, fmt.Errorf("encryption of access key failed : %w", err)
	}

	encSecret, err := security.Encrypt(accessKey)
	if err != nil {
		return database.Bucket{}, fmt.Errorf("encryption of secret key failed : %w", err)
	}

}

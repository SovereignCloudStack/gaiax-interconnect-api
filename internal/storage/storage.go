package storage

import (
	"context"

	"github.com/akafazov/gaiax-interconnect-api/internal/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initializes the postgres driver
)

type StorageInterface interface {
	AddVPN(ctx context.Context, vpn model.AddVPNRequest) (int, error)
	GetVPN(ctx context.Context, id int) (model.VPN, error)
	GetVPNs(ctx context.Context) ([]model.VPN, error)
	UpdateVPN(ctx context.Context, vpn model.UpdateVPNRequest) (int, error)
	DeleteVPN(ctx context.Context, id int) error
	VerifyVPNExists(ctx context.Context, id int) (bool, error)
}

// Storage contains an SQL db. Storage implements the StorageInterface.
type Storage struct {
	db *sqlx.DB
}

func (s *Storage) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetDB() *sqlx.DB {
	return s.db
}

package storage

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/akafazov/gaiax-interconnect-api/internal/model"
)

func (s *Storage) AddVPN(ctx context.Context, vpn model.AddVPNRequest) (int, error) {
	var id int
	err := s.db.Get(&id, `INSERT INTO vpn(name, type, local_as_number, remote_as_number, vni, created_at, updated_at)
			VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`, vpn.Name, vpn.Type, vpn.LocalAsNumber, vpn.RemoteAsNumber, vpn.VNI, time.Now().UTC(), time.Now().UTC())

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) GetVPN(ctx context.Context, id int) (model.VPN, error) {
	var vpn model.VPN

	err := s.db.Get(&vpn, `Select * from vpn where id=$1`, id)
	if err != nil {
		return vpn, err
	}

	return vpn, nil
}

func (s *Storage) GetVPNs(ctx context.Context) ([]model.VPN, error) {
	var vpns []model.VPN
	err := s.db.Select(&vpns, `SELECT * from vpn`)
	if err != nil {
		return nil, err
	}

	return vpns, nil
}

func (s *Storage) UpdateVPN(ctx context.Context, vpn model.UpdateVPNRequest) (int, error) {
	var columns []string
	var argCount = 1
	var args []interface{}

	if vpn.Name != "" {
		columns = append(columns, fmt.Sprintf("name=$%d", argCount))
		args = append(args, vpn.Name)
		argCount++
	}

	if vpn.Type != "" {
		columns = append(columns, fmt.Sprintf("type=$%d", argCount))
		args = append(args, vpn.Type)
		argCount++
	}

	if vpn.LocalAsNumber != 0 {
		columns = append(columns, fmt.Sprintf("local_as_number=$%d", argCount))
		args = append(args, vpn.LocalAsNumber)
		argCount++
	}

	if vpn.RemoteAsNumber != 0 {
		columns = append(columns, fmt.Sprintf("remote_as_number=$%d", argCount))
		args = append(args, vpn.RemoteAsNumber)
		argCount++
	}

	if vpn.VNI != 0 {
		columns = append(columns, fmt.Sprintf("vni=$%d", argCount))
		args = append(args, vpn.VNI)
		argCount++
	}

	columns = append(columns, fmt.Sprintf("updated_at=$%d", argCount))
	args = append(args, time.Now().UTC())
	argCount++

	if len(columns) == 0 {
		return 0, errors.New("No fields to update")
	}

	args = append(args, vpn.ID)

	query := fmt.Sprintf(`UPDATE vpn SET %s WHERE id=$%d RETURNING id`, strings.Join(columns, ", "), argCount)

	var id int
	err := s.db.Get(&id, query, args...)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) DeleteVPN(ctx context.Context, id int) error {
	_, err := s.db.Exec(`DELETE FROM vpn WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) VerifyVPNExists(ctx context.Context, id int) (bool, error) {
	var exists bool
	err := s.db.Get(&exists, `SELECT EXISTS(SELECT 1 from vpn where id=$1)`, id)
	if err != nil {
		return false, err
	}

	return exists, nil
}

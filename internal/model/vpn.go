package model

// VPN is the db schema for the VPNs table
type VPN struct {
	ID             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Type           string `json:"type" db:"type"`
	LocalAsNumber  int    `json:"local_as_number" db:"local_as_number"`
	RemoteAsNumber int    `json:"remote_as_number" db:"remote_as_number"`
	VNI            int    `json:"vni" db:"vni"`
	Base
}

// Below are the structures of the request/response structs in the VPNs handler.

type AddVPNRequest struct {
	Name           string `json:"name" validate:"required"`
	Type           string `json:"type" validate:"required"`
	LocalAsNumber  int    `json:"local_as_number" validate:"required"`
	RemoteAsNumber int    `json:"remote_as_number" validate:"required"`
	VNI            int    `json:"vni" validate:"required"`
}

type GetVPNResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	LocalAsNumber  int    `json:"local_as_number"`
	RemoteAsNumber int    `json:"remote_as_number"`
	VNI            int    `json:"vni"`
}

type UpdateVPNRequest struct {
	ID             int    `json:"id" validate:"required"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	LocalAsNumber  int    `json:"local_as_number"`
	RemoteAsNumber int    `json:"remote_as_number"`
	VNI            int    `json:"vni"`
}

type IDResponse struct {
	ID int `json:"id"`
}

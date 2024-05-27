package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/akafazov/gaiax-interconnect-api/internal/model"
	"github.com/akafazov/gaiax-interconnect-api/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// GetVPNs godoc
//
//	@Summary		Get all VPNs
//	@Tags			VPNs
//	@Produce		json
//	@Success		200	{object} []model.GetVPNResponse
//
// @Router			/api/vpns [get]
func (h *Handlers) GetVPNsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vpns, err := h.Storage.GetVPNs(ctx)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	vpnResponse := []model.GetVPNResponse{}
	for _, vpn := range vpns {
		vpnResponse = append(vpnResponse, model.GetVPNResponse{
			ID:             vpn.ID,
			Name:           vpn.Name,
			Type:           vpn.Type,
			LocalAsNumber:  vpn.LocalAsNumber,
			RemoteAsNumber: vpn.RemoteAsNumber,
			VNI:            vpn.VNI,
		})

	}

	err = h.Sender.JSON(w, http.StatusOK, vpnResponse)
	if err != nil {
		logger.OutputLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatal("Error when requesting /vpns")

		panic(err)
	}
}

// GetVPN godoc
//
//	@Summary		Get a specific VPN
//	@Tags			VPNs
//	@Produce		json
//	@Param			id path int	true "VPN ID"
//	@Success		200	{object} model.GetVPNResponse
//
// @Router			/api/vpn/{id} [get]
func (h *Handlers) GetVPNHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	var vpn model.VPN
	vpn, err = h.Storage.GetVPN(ctx, id)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if (model.VPN{}) == vpn {
		h.Sender.JSON(w, http.StatusBadRequest, "VPN with id="+fmt.Sprint(vpn.ID)+" not found")
		if err != nil {
			panic(err)
		}
		return
	}

	vpnResponse := model.GetVPNResponse{
		ID:             vpn.ID,
		Name:           vpn.Name,
		Type:           vpn.Type,
		LocalAsNumber:  vpn.LocalAsNumber,
		RemoteAsNumber: vpn.RemoteAsNumber,
		VNI:            vpn.VNI,
	}

	err = h.Sender.JSON(w, http.StatusOK, vpnResponse)
	if err != nil {
		logger.OutputLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatal(fmt.Sprint("Error when requesting /vpn/", vpn.ID))

		panic(err)
	}
}

// AddVPN godoc
//
//		@Summary		Add a specific VPN
//		@Tags			VPNs
//		@Produce		json
//	 	@Accept			json
//		@Param			name body string true "VPN name"
//		@Param			type body string true "VPN type"
//		@Param			localAsNumber body int true "VPN localAsNumber"
//		@Param			remoteAsNumber body int true "VPN remoteAsNumber"
//		@Param			vni body int true "VPN vni"
//		@Success		200	{object} model.IDResponse
//
// @Router			/api/vpn/add [post]
func (h *Handlers) AddVPNHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var vpn model.AddVPNRequest

	err := json.NewDecoder(r.Body).Decode(&vpn)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = Validate.Struct(vpn)
	if err != nil {
		var errs []string
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, err.Field()+" "+err.Tag())
		}
		h.Sender.JSON(w, http.StatusBadRequest, strings.Join(errs, ", "))
		return
	}

	id, err := h.Storage.AddVPN(ctx, vpn)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := model.IDResponse{ID: id}
	err = h.Sender.JSON(w, http.StatusOK, response)
	if err != nil {
		panic(err)
	}
}

// UpdateVPN godoc
//
//		@Summary		Update a specific VPN
//		@Tags			VPNs
//		@Produce		json
//	 	@Accept			json
//		@Param			name body string false "VPN name"
//		@Param			type body string false "VPN type"
//		@Param			localAsNumber body int false "VPN localAsNumber"
//		@Param			remoteAsNumber body int false "VPN remoteAsNumber"
//		@Param			vni body int false "VPN vni"
//		@Success		200	{object} model.IDResponse
//
// @Router			/api/vpn/update [patch]
func (h *Handlers) UpdateVPNHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var vpn model.UpdateVPNRequest

	err := json.NewDecoder(r.Body).Decode(&vpn)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = Validate.Struct(vpn)
	if err != nil {
		var errs []string
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, err.Field()+" "+err.Tag())
		}
		h.Sender.JSON(w, http.StatusBadRequest, strings.Join(errs, ", "))
		return
	}

	exists, err := h.Storage.VerifyVPNExists(ctx, vpn.ID)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		h.Sender.JSON(w, http.StatusBadRequest, "VPN with id="+fmt.Sprint(vpn.ID)+" not found")
		return
	}

	id, err := h.Storage.UpdateVPN(ctx, vpn)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := model.IDResponse{ID: id}
	err = h.Sender.JSON(w, http.StatusOK, response)
	if err != nil {
		panic(err)
	}
}

// DeleteVPN godoc
//
//	@Summary		Delete a specific VPN
//	@Tags			VPNs
//	@Produce		json
//	@Param			id path int	true "VPN ID"
//	@Success		200	{object} model.GetVPNResponse
//
// @Router			/api/vpn/delete/{id} [delete]
func (h *Handlers) DeleteVPNHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	exists, err := h.Storage.VerifyVPNExists(ctx, id)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		err = h.Sender.JSON(w, http.StatusBadRequest, "VPN with id="+fmt.Sprint(id)+" not found")
		if err != nil {
			panic(err)
		}
		return
	}

	err = h.Storage.DeleteVPN(ctx, id)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.Sender.JSON(w, http.StatusOK, map[string]bool{"success": true})
	if err != nil {
		panic(err)
	}
}

// cSpell:ignore godoc logrus

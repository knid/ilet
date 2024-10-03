package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/knid/ilet/internal/models"
	"github.com/knid/ilet/internal/utils"
)

func (h *HTTPHandler) RouteToLongURL(w http.ResponseWriter, r *http.Request) {
	short := chi.URLParam(r, "shortLink")
    log.Println(short)

	link, err := h.DB.GetLinkByShortURL(short)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusNotFound, err)
		return
	}

	http.Redirect(w, r, link.Long, http.StatusMovedPermanently)
}

func (h *HTTPHandler) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUserFromRequest(r, h.DB)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	links, err := h.DB.GetLinksByUser(user)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusNotFound, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, links)
}

func (h *HTTPHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUserFromRequest(r, h.DB)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, errors.New("parsing error: id"))
		return
	}

	link, err := h.DB.GetLinkById(intId)

	if link.User.ID != user.ID {
		utils.JsonErrorResponse(w, http.StatusForbidden, errors.New("you do not have permission to access this resource"))
		return
	}

	utils.JsonResponse(w, http.StatusOK, link)
}

func (h *HTTPHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUserFromRequest(r, h.DB)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	var link models.Link
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&link); err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	link.User.ID = user.ID

	createdLink, err := h.DB.CreateLink(link)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, createdLink)
}

func (h *HTTPHandler) UpdateLink(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUserFromRequest(r, h.DB)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, errors.New("parsing error: id"))
		return
	}

	link, err := h.DB.GetLinkById(intId)

	if link.User.ID != user.ID {
		utils.JsonErrorResponse(w, http.StatusForbidden, errors.New("you do not have permission to access this resource"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&link); err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	link.ID = intId

	updatedLink, err := h.DB.UpdateLink(link)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, updatedLink)
}

func (h *HTTPHandler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUserFromRequest(r, h.DB)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, errors.New("parsing error: id"))
		return
	}

	link, err := h.DB.GetLinkById(intId)

	if link.User.ID != user.ID {
		utils.JsonErrorResponse(w, http.StatusForbidden, errors.New("you do not have permission to access this resource"))
		return
	}

	if err := h.DB.DeleteLink(link); err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	utils.JsonResponse(w, http.StatusNoContent, "")
}

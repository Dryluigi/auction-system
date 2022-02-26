package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Dryluigi/auction-system/apiHelpers/response"
	request "github.com/Dryluigi/auction-system/request/v1"
	serviceError "github.com/Dryluigi/auction-system/service/api/errors"
	service "github.com/Dryluigi/auction-system/service/api/v1"
	"github.com/gorilla/mux"
)

type BidSessionControllerImpl struct {
	BidSessionService service.BidSessionService
	BiddingService    service.BiddingService
}

func (controller *BidSessionControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	bidSession := &request.BidSessionSave{}

	err := json.NewDecoder(r.Body).Decode(bidSession)
	if err != nil {
		response.BuildErrorResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	model, err := controller.BidSessionService.Save(bidSession)
	if err != nil {
		response.BuildErrorResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	response.BuildSuccessResponse(w, http.StatusCreated, "Bid Session created", true, model)
}

func (controller *BidSessionControllerImpl) Occuring(w http.ResponseWriter, r *http.Request) {
	occuring, err := controller.BidSessionService.GetOccuring()

	if err != nil {
		response.BuildErrorResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	if occuring == nil {
		response.BuildSuccessResponse(w, http.StatusOK, "Fetched", true, nil)
		return
	}

	response.BuildSuccessResponse(w, http.StatusOK, "Fetched", true, occuring)
}

func (controller *BidSessionControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.BuildErrorResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	err = controller.BidSessionService.Delete(uint(id))

	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, serviceError.ErrEntityNotFound) {
			status = http.StatusNotFound
		}

		response.BuildErrorResponse(w, status, err, nil)
		return
	}

	response.BuildSuccessResponse(w, http.StatusOK, "Bid Session deleted successfully", true, nil)
}

func (controller *BidSessionControllerImpl) BidItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bidSessionId, err := strconv.ParseUint(vars["bidSessionId"], 10, 32)
	if err != nil {
		response.BuildErrorResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	productId, err := strconv.ParseUint(vars["productId"], 10, 32)
	if err != nil {
		response.BuildErrorResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	bidProduct, err := controller.BiddingService.BidItem(uint(bidSessionId), uint(productId))
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, serviceError.ErrEntityNotFound) {
			status = http.StatusBadRequest
		}

		response.BuildErrorResponse(w, status, err, nil)
		return
	}

	response.BuildSuccessResponse(w, http.StatusOK, "Bid successfully", true, bidProduct)
}

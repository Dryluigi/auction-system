package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Dryluigi/auction-system/apiHelpers/response"
	request "github.com/Dryluigi/auction-system/request/v1"
	"github.com/Dryluigi/auction-system/service/api/errors"

	service "github.com/Dryluigi/auction-system/service/api/v1"
	"github.com/Dryluigi/auction-system/validator"
	"github.com/gorilla/mux"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func (controller *ProductControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	saveRequest := &request.ProductSave{}
	err := json.NewDecoder(r.Body).Decode(saveRequest)

	if err != nil {
		response.BuildErrorResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	invalids, ok := validator.Validator.Validate(saveRequest)

	if !ok {
		response.BuildErrorResponse(w, http.StatusBadRequest, err, invalids)
		return
	}

	product, err := controller.ProductService.Save(saveRequest)

	if err != nil {
		response.BuildErrorResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	response.BuildSuccessResponse(w, http.StatusCreated, "Product created successfully", true, product)
}

func (controller *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	productId, err := controller.extractUrlProductId(r)

	if err != nil {
		response.BuildErrorResponse(w, http.StatusBadRequest, err, "Product ID should be integer")
		return
	}

	updateRequest := &request.ProductUpdate{}
	err = json.NewDecoder(r.Body).Decode(updateRequest)

	if err != nil {
		response.BuildErrorResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	product, err := controller.ProductService.Update(uint(productId), updateRequest)

	if err != nil {
		if err == errors.ErrEntityNotFound {
			response.BuildErrorResponse(w, http.StatusNotFound, err, nil)
		} else {
			response.BuildErrorResponse(w, http.StatusInternalServerError, err, nil)
		}
		return
	}

	response.BuildSuccessResponse(w, http.StatusOK, "Product updated successfully", true, product)
}

func (controller *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	productId, err := controller.extractUrlProductId(r)

	if err != nil {
		response.BuildErrorResponse(w, http.StatusBadRequest, err, "Product ID should be integer")
		return
	}

	err = controller.ProductService.Delete(uint(productId))

	if err != nil {
		response.BuildErrorResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	response.BuildSuccessResponse(w, http.StatusOK, "Product deleted successfully", true, nil)
}

func (controller *ProductControllerImpl) extractUrlProductId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	productId, err := strconv.Atoi(idStr)

	if err != nil {
		return 0, err
	}

	return productId, err
}

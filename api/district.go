package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/irossa/educate/db/sqlc"
)

type createDistrictRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createDistrict(ctx *gin.Context) {
	var req createDistrictRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	name := req.Name

	district, err := server.store.CreateDistrict(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, district)
}

type getDistrictRequest struct {
	ID int64 `form:"id" binding:"required"`
}

func (server *Server) getDistrict(ctx *gin.Context) {
	var req getDistrictRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	district, err := server.store.GetDistrict(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, district)
}

type getAllDistrictsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getAllDistricts(ctx *gin.Context) {
	var req getAllDistrictsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllDistrictsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	districts, err := server.store.GetAllDistricts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, districts)
}

type updateDistrictRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (server *Server) updateDistrict(ctx *gin.Context) {
	var req updateDistrictRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateDistrictParams{
		Name: req.Name,
		ID:   req.ID,
	}

	district, err := server.store.UpdateDistrict(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, district)
}

type deleteDistrictRequest struct {
	ID int64 `form:"id" binding:"required"`
}

func (server *Server) deleteDistrict(ctx *gin.Context) {
	var req deleteDistrictRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteDistrict(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

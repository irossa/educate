package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/irossa/educate/db/sqlc"
)

type createSchoolRequest struct {
	Name       string        `json:"name" binding:"required"`
	DistrictID sql.NullInt64 `json:"districtid" binding:"required"`
}

func (server *Server) createSchool(ctx *gin.Context) {
	var req createSchoolRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSchoolParams{
		Name:       req.Name,
		DistrictID: req.DistrictID,
	}

	school, err := server.store.CreateSchool(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, school)
}

type getSchoolRequest struct {
	ID int64 `form:"id" binding:"required"`
}

func (server *Server) getSchool(ctx *gin.Context) {
	var req getSchoolRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	school, err := server.store.GetSchool(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, school)
}

type getAllSchoolsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getAllSchools(ctx *gin.Context) {
	var req getAllSchoolsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllSchoolsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	schools, err := server.store.GetAllSchools(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schools)
}

type updateSchoolRequest struct {
	ID         int64         `json:"id" binding:"required"`
	Name       string        `json:"name" binding:"required"`
	DistrictID sql.NullInt64 `json:"districtid"`
}

func (server *Server) updateSchool(ctx *gin.Context) {
	var req updateSchoolRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateSchoolParams{
		Name:       req.Name,
		ID:         req.ID,
		DistrictID: req.DistrictID,
	}

	school, err := server.store.UpdateSchool(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, school)
}

type deleteSchoolRequest struct {
	ID int64 `form:"id" binding:"required"`
}

func (server *Server) deleteSchool(ctx *gin.Context) {
	var req deleteSchoolRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteSchool(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

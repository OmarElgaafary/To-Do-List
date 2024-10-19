package handlers

import (
	"log"
	"net/http"
	"todo/internal/models"
	"todo/internal/service"

	"github.com/labstack/echo/v4"
)

type ToDoHandler struct { // constructor struct
	tdsvc *service.ToDoService
}

func NewtoDoHandler(svc *service.ToDoService) *ToDoHandler { // constructor struct
	return &ToDoHandler{tdsvc: svc}
}

func (tdH *ToDoHandler) CreateToDoHandler(c echo.Context) error {
	// PARSE REQUEST BODY (CHECK IF REQUEST IS VALID)
	var Body models.ToDo
	if err := c.Bind(&Body); err != nil {
		//return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
		//return echo.ErrBadRequest
		return echo.NewHTTPError(http.StatusBadRequest, "bad response")
	}

	// PREPARE RESPONSE (CREATE TODO)
	err := tdH.tdsvc.CreateToDo(&Body)
	if err != nil {
		// return c.JSON(http.StatusInternalServerError, map[string]string{"error": "server error"})
		return echo.NewHTTPError(http.StatusInternalServerError, "server error")
	}

	// RESPONSE RIGHT
	// return c.JSON(http.StatusOK, struct{}{})
	return c.NoContent(http.StatusCreated)
}

func (tdH *ToDoHandler) GetToDoByIDHandler(c echo.Context) error {

	// PARSE REQUEST BODY (GET ID)
	paramId := c.Param("id")
	if paramId == "" {
		return echo.ErrBadRequest
	}

	// Get to do by param ID
	log.Printf("type of paramID %T", paramId)
	toDo, err := tdH.tdsvc.GetToDoByID(paramId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	//
	return c.JSON(http.StatusOK, toDo)

}

// func (tdH *ToDoHandler) UpdateToDoHandler(c echo.Context) error {

// }

// func (tdJ *ToDoHandler) DeleteToDoByIDHandler(c echo.Context) error {

// }

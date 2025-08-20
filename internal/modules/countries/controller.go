package countries

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct{ svc Service }

func NewController(s Service) *Controller { return &Controller{svc: s} }

func (h *Controller) Create(c echo.Context) error {
	var in CreateCountryDTO
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid json"})
	}
	out, err := h.svc.Create(c.Request().Context(), in)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, out)
}

func (h *Controller) List(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	size, _ := strconv.Atoi(c.QueryParam("page_size"))
	items, err := h.svc.List(c.Request().Context(), page, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "list error"})
	}
	return c.JSON(http.StatusOK, items)
}

func (h *Controller) Get(c echo.Context) error {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid id"})
	}
	item, err := h.svc.Get(c.Request().Context(), uint(id64))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "not found"})
	}
	return c.JSON(http.StatusOK, item)
}

func (h *Controller) Update(c echo.Context) error {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid id"})
	}
	var in UpdateCountryDTO
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid json"})
	}
	item, err := h.svc.Update(c.Request().Context(), uint(id64), in)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, item)
}

func (h *Controller) Delete(c echo.Context) error {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid id"})
	}
	if err := h.svc.Delete(c.Request().Context(), uint(id64)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "delete error"})
	}
	return c.NoContent(http.StatusNoContent)
}

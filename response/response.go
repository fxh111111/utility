package response

import (
	"net/http"
	"utility/wmwerrors"
)

type GinContext interface {
	AbortWithStatusJSON(code int, jsonObj any)
}

func ErrorExit(c GinContext, err wmwerrors.WMWError) {
	c.AbortWithStatusJSON(http.StatusOK, map[string]any{"code": err.Code(), "message": err.Error()})
}

func DataExit(c GinContext, data any) {
	err := wmwerrors.Nil()
	c.AbortWithStatusJSON(http.StatusOK, map[string]any{"code": err.Code(), "message": err.Error(), "data": data})
}

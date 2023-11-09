package middleware

import (
	"github.com/coby-amar/diet_tracker/main/utils"
)

type configHandler func(*utils.ConfigWithRequestAndResponse)
type jsonHandler[T interface{}] func(*utils.ConfigWithRequestAndResponse, T)
type jsonParams[T interface{}] struct {
	params T
}

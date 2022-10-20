package render

import (
	"net/http"
)

type codeErr struct {
	msg  string
	code int
}

func (e codeErr) Error() string {
	return e.msg
}

func (e codeErr) Code() int {
	return e.code
}

func NewErrBadRequest(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusBadRequest,
	}
}

func NewErrUnauthorized(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusUnauthorized,
	}
}

func NewErrPaymentRequired(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusPaymentRequired,
	}
}

func NewErrForbidden(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusForbidden,
	}
}

func NewErrNotFound(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusNotFound,
	}
}

func NewErrMethodNotAllowed(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusMethodNotAllowed,
	}
}

func NewErrNotAcceptable(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusNotAcceptable,
	}
}

func NewErrProxyAuthRequired(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusProxyAuthRequired,
	}
}

func NewErrRequestTimeout(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusRequestTimeout,
	}
}

func NewErrConflict(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusConflict,
	}
}

func NewErrGone(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusGone,
	}
}

func NewErrLengthRequired(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusLengthRequired,
	}
}

func NewErrPreconditionFailed(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusPreconditionFailed,
	}
}

func NewErrRequestEntityTooLarge(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusRequestEntityTooLarge,
	}
}

func NewErrRequestURITooLong(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusRequestURITooLong,
	}
}

func NewErrUnsupportedMediaType(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusUnsupportedMediaType,
	}
}

func NewErrRequestedRangeNotSatisfiable(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusRequestedRangeNotSatisfiable,
	}
}

func NewErrExpectationFailed(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusExpectationFailed,
	}
}

func NewErrTeapot(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusTeapot,
	}
}

func NewErrMisdirectedRequest(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusMisdirectedRequest,
	}
}

func NewErrUnprocessableEntity(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusUnprocessableEntity,
	}
}

func NewErrLocked(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusLocked,
	}
}

func NewErrFailedDependency(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusFailedDependency,
	}
}

func NewErrTooEarly(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusTooEarly,
	}
}

func NewErrUpgradeRequired(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusUpgradeRequired,
	}
}

func NewErrPreconditionRequired(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusPreconditionRequired,
	}
}

func NewErrTooManyRequests(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusTooManyRequests,
	}
}

func NewErrRequestHeaderFieldsTooLarge(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusRequestHeaderFieldsTooLarge,
	}
}

func NewErrUnavailableForLegalReasons(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusUnavailableForLegalReasons,
	}
}

func NewErrInternalServerError(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusInternalServerError,
	}
}

func NewErrNotImplemented(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusNotImplemented,
	}
}

func NewErrBadGateway(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusBadGateway,
	}
}

func NewErrServiceUnavailable(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusServiceUnavailable,
	}
}

func NewErrGatewayTimeout(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusGatewayTimeout,
	}
}

func NewErrHTTPVersionNotSupported(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusHTTPVersionNotSupported,
	}
}

func NewErrVariantAlsoNegotiates(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusVariantAlsoNegotiates,
	}
}

func NewErrInsufficientStorage(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusInsufficientStorage,
	}
}

func NewErrLoopDetected(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusLoopDetected,
	}
}

func NewErrNotExtended(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusNotExtended,
	}
}

func NewErrNetworkAuthenticationRequired(msg string) error {
	return codeErr{
		msg:  msg,
		code: http.StatusNetworkAuthenticationRequired,
	}
}

var (
	ErrBadRequest                   = NewErrBadRequest(http.StatusText(http.StatusBadRequest))
	ErrUnauthorized                 = NewErrUnauthorized(http.StatusText(http.StatusUnauthorized))
	ErrPaymentRequired              = NewErrPaymentRequired(http.StatusText(http.StatusPaymentRequired))
	ErrForbidden                    = NewErrForbidden(http.StatusText(http.StatusForbidden))
	ErrNotFound                     = NewErrNotFound(http.StatusText(http.StatusNotFound))
	ErrMethodNotAllowed             = NewErrMethodNotAllowed(http.StatusText(http.StatusMethodNotAllowed))
	ErrNotAcceptable                = NewErrNotAcceptable(http.StatusText(http.StatusNotAcceptable))
	ErrProxyAuthRequired            = NewErrProxyAuthRequired(http.StatusText(http.StatusProxyAuthRequired))
	ErrRequestTimeout               = NewErrRequestTimeout(http.StatusText(http.StatusRequestTimeout))
	ErrConflict                     = NewErrConflict(http.StatusText(http.StatusConflict))
	ErrGone                         = NewErrGone(http.StatusText(http.StatusGone))
	ErrLengthRequired               = NewErrLengthRequired(http.StatusText(http.StatusLengthRequired))
	ErrPreconditionFailed           = NewErrPreconditionFailed(http.StatusText(http.StatusPreconditionFailed))
	ErrRequestEntityTooLarge        = NewErrRequestEntityTooLarge(http.StatusText(http.StatusRequestEntityTooLarge))
	ErrRequestURITooLong            = NewErrRequestURITooLong(http.StatusText(http.StatusRequestURITooLong))
	ErrUnsupportedMediaType         = NewErrUnsupportedMediaType(http.StatusText(http.StatusUnsupportedMediaType))
	ErrRequestedRangeNotSatisfiable = NewErrRequestedRangeNotSatisfiable(http.StatusText(http.StatusRequestedRangeNotSatisfiable))
	ErrExpectationFailed            = NewErrExpectationFailed(http.StatusText(http.StatusExpectationFailed))
	ErrTeapot                       = NewErrTeapot(http.StatusText(http.StatusTeapot))
	ErrMisdirectedRequest           = NewErrMisdirectedRequest(http.StatusText(http.StatusMisdirectedRequest))
	ErrUnprocessableEntity          = NewErrUnprocessableEntity(http.StatusText(http.StatusUnprocessableEntity))
	ErrLocked                       = NewErrLocked(http.StatusText(http.StatusLocked))
	ErrFailedDependency             = NewErrFailedDependency(http.StatusText(http.StatusFailedDependency))
	ErrTooEarly                     = NewErrTooEarly(http.StatusText(http.StatusTooEarly))
	ErrUpgradeRequired              = NewErrUpgradeRequired(http.StatusText(http.StatusUpgradeRequired))
	ErrPreconditionRequired         = NewErrPreconditionRequired(http.StatusText(http.StatusPreconditionRequired))
	ErrTooManyRequests              = NewErrTooManyRequests(http.StatusText(http.StatusTooManyRequests))
	ErrRequestHeaderFieldsTooLarge  = NewErrRequestHeaderFieldsTooLarge(http.StatusText(http.StatusRequestHeaderFieldsTooLarge))
	ErrUnavailableForLegalReasons   = NewErrUnavailableForLegalReasons(http.StatusText(http.StatusUnavailableForLegalReasons))

	ErrInternalServerError           = NewErrInternalServerError(http.StatusText(http.StatusInternalServerError))
	ErrNotImplemented                = NewErrNotImplemented(http.StatusText(http.StatusNotImplemented))
	ErrBadGateway                    = NewErrBadGateway(http.StatusText(http.StatusBadGateway))
	ErrServiceUnavailable            = NewErrServiceUnavailable(http.StatusText(http.StatusServiceUnavailable))
	ErrGatewayTimeout                = NewErrGatewayTimeout(http.StatusText(http.StatusGatewayTimeout))
	ErrHTTPVersionNotSupported       = NewErrHTTPVersionNotSupported(http.StatusText(http.StatusHTTPVersionNotSupported))
	ErrVariantAlsoNegotiates         = NewErrVariantAlsoNegotiates(http.StatusText(http.StatusVariantAlsoNegotiates))
	ErrInsufficientStorage           = NewErrInsufficientStorage(http.StatusText(http.StatusInsufficientStorage))
	ErrLoopDetected                  = NewErrLoopDetected(http.StatusText(http.StatusLoopDetected))
	ErrNotExtended                   = NewErrNotExtended(http.StatusText(http.StatusNotExtended))
	ErrNetworkAuthenticationRequired = NewErrNetworkAuthenticationRequired(http.StatusText(http.StatusNetworkAuthenticationRequired))
)

type tagErr struct {
	err  error
	code int
}

func (e tagErr) Error() string {
	return e.err.Error()
}

func (e tagErr) Code() int {
	return e.code
}

func (e tagErr) Unwrap() error {
	return e.err
}

func TagBadRequest(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusBadRequest,
	}
}

func TagUnauthorized(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusUnauthorized,
	}
}

func TagPaymentRequired(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusPaymentRequired,
	}
}

func TagForbidden(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusForbidden,
	}
}

func TagNotFound(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusNotFound,
	}
}

func TagMethodNotAllowed(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusMethodNotAllowed,
	}
}

func TagNotAcceptable(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusNotAcceptable,
	}
}

func TagProxyAuthRequired(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusProxyAuthRequired,
	}
}

func TagRequestTimeout(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusRequestTimeout,
	}
}

func TagConflict(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusConflict,
	}
}

func TagGone(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusGone,
	}
}

func TagLengthRequired(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusLengthRequired,
	}
}

func TagPreconditionFailed(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusPreconditionFailed,
	}
}

func TagRequestEntityTooLarge(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusRequestEntityTooLarge,
	}
}

func TagRequestURITooLong(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusRequestURITooLong,
	}
}

func TagUnsupportedMediaType(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusUnsupportedMediaType,
	}
}

func TagRequestedRangeNotSatisfiable(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusRequestedRangeNotSatisfiable,
	}
}

func TagExpectationFailed(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusExpectationFailed,
	}
}

func TagTeapot(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusTeapot,
	}
}

func TagMisdirectedRequest(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusMisdirectedRequest,
	}
}

func TagUnprocessableEntity(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusUnprocessableEntity,
	}
}

func TagLocked(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusLocked,
	}
}

func TagFailedDependency(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusFailedDependency,
	}
}

func TagTooEarly(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusTooEarly,
	}
}

func TagUpgradeRequired(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusUpgradeRequired,
	}
}

func TagPreconditionRequired(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusPreconditionRequired,
	}
}

func TagTooManyRequests(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusTooManyRequests,
	}
}

func TagRequestHeaderFieldsTooLarge(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusRequestHeaderFieldsTooLarge,
	}
}

func TagUnavailableForLegalReasons(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusUnavailableForLegalReasons,
	}
}

func TagInternalServerError(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusInternalServerError,
	}
}

func TagNotImplemented(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusNotImplemented,
	}
}

func TagBadGateway(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusBadGateway,
	}
}

func TagServiceUnavailable(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusServiceUnavailable,
	}
}

func TagGatewayTimeout(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusGatewayTimeout,
	}
}

func TagHTTPVersionNotSupported(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusHTTPVersionNotSupported,
	}
}

func TagVariantAlsoNegotiates(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusVariantAlsoNegotiates,
	}
}

func TagInsufficientStorage(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusInsufficientStorage,
	}
}

func TagLoopDetected(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusLoopDetected,
	}
}

func TagNotExtended(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusNotExtended,
	}
}

func TagNetworkAuthenticationRequired(err error) error {
	return tagErr{
		err:  err,
		code: http.StatusNetworkAuthenticationRequired,
	}
}

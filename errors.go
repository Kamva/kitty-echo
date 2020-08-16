package hecho

import (
	"errors"
	"github.com/kamva/hexa"
	"net/http"
)

var (
	// Error code description:
	// hec = hexa echo (package or project name)
	// 1 = errors about user section (identify some part in application)
	// E = Error (type of code : error|response)
	// 00 = error number zero (id of code in that part and type)

	//--------------------------------
	// User and authentication Errors
	//--------------------------------
	errUserNotFound = hexa.NewError(
		http.StatusInternalServerError,
		"lib.user.not_found_error",
		errors.New("user not found"),
	)

	errContextUserNotImplementedHexaUser = hexa.NewError(
		http.StatusInternalServerError,
		"lib.user.interface_not_implemented_error",
		errors.New("user in the hexa context does not implemented User interface"),
	)

	errJwtMissing = hexa.NewError(
		http.StatusBadRequest,
		"lib.user.missing_jwt_token_error",
		errors.New("missing or malformed jwt"),
	)

	errInvalidOrExpiredJwt = hexa.NewError(
		http.StatusUnauthorized,
		"lib.user.invalid_expired_jwt_error",
		errors.New("invalid or expired jwt"),
	)

	errUserMustBeGuest = hexa.NewError(
		http.StatusUnauthorized,
		"lib.user.must_be_guest_error",
		errors.New("the user must be guest to access this API"),
	)

	errUserNeedToAuthenticate = hexa.NewError(
		http.StatusUnauthorized,
		"lib.user.must_authenticate_error",
		errors.New("the user need to login to access this API"),
	)

	errRefreshTokenCanNotBeEmpty = hexa.NewError(
		http.StatusBadRequest,
		"lib.user.refresh_token_is_empty_error",
		errors.New("refresh token can not be empty"),
	)

	errInvalidRefreshToken = hexa.NewError(http.StatusBadRequest, "lib.user.invalid_refresh_token_error", nil)

	//--------------------------------
	// Request errors
	//--------------------------------
	errRequestIdNotFound = hexa.NewError(
		http.StatusInternalServerError,
		"lib.request.request_id_not_found_error",
		errors.New("request id not found in the request"),
	)

	errCorrelationIDNotFound = hexa.NewError(
		http.StatusInternalServerError,
		"lib.request.correlation_id_not_found_error",
		errors.New("correlation id not found in the request"),
	)

	//--------------------------------
	// DEBUG
	//--------------------------------
	errRouteAvailableInDebugMode = hexa.NewError(
		http.StatusUnauthorized,
		"lib.route.available_in_debug_mode",
		errors.New("route is available just in debug mode"),
	)

	//--------------------------------
	// Other errors
	//--------------------------------
	errHTTPNotFoundError = hexa.NewError(http.StatusNotFound, "lib.route.not_found", nil)

	// Set this error status manually on return relative to echo error code.
	errEchoHTTPError = hexa.NewError(http.StatusNotFound, "lib.error.occurred_an_error", nil)

	errUnknownError = hexa.NewError(http.StatusInternalServerError, "lib.unknown.unknown_error", nil)
)

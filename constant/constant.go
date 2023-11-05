package constant

import "time"

const (
	// code
	RespCodeSuccess               = 200
	RespCodeResourceNotFound      = 404
	RespCodeClientParamError      = 400
	RespCodeUnauthorized          = 401
	RespCodeInvalidResourceAccess = 403
	RespCodeServerError           = 500

	// message
	RespMsgSuccess               = "success"
	RespMsgInvalidResourceAccess = "error: no permission to access the request resource"
	RespMsgUnauthenticated       = "error: Unauthenticated, please login first or check token data"
	RespMsgResourceNotFound      = "error: request resource is not found"
	RespMsgServerError           = "error: server error"

	// auth
	TokenIssuer           = "app"
	TokenValidityDuration = time.Minute * 10

	// user
	UserRoleAdmin   = "admin"
	UserRoleStudent = "student"

	// header
	AppUserIdHeader   = "app-user-id"
	AppUserRoleHeader = "app-user-role"

	// redisKey
	MockUserListRedisKey = "app:mock:user:list"
)

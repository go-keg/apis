// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package common

import (
	response "github.com/go-keg/keg/third_party/response"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the package it is being compiled against.
const _ = response.SupportPackageIsVersion1

func IsSuccess(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_Success.String() && e.GetCode() == 0
}

func Success(args ...interface{}) *response.Response {
	return response.Newf(200, 0, CommonError_Success.String(), "Success", args...)
}

func IsErrUnknown(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrUnknown.String() && e.GetCode() == 100001
}

func ErrUnknown(args ...interface{}) *response.Response {
	return response.Newf(500, 100001, CommonError_ErrUnknown.String(), "Internal Server Error", args...)
}

func IsErrCommon(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrCommon.String() && e.GetCode() == 100002
}

func ErrCommon(args ...interface{}) *response.Response {
	return response.Newf(500, 100002, CommonError_ErrCommon.String(), "Internal Server Error", args...)
}

func IsErrValidate(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrValidate.String() && e.GetCode() == 100003
}

func ErrValidate(args ...interface{}) *response.Response {
	return response.Newf(200, 100003, CommonError_ErrValidate.String(), "%s", args...)
}

func IsErrUnauthorized(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrUnauthorized.String() && e.GetCode() == 100100
}

func ErrUnauthorized(args ...interface{}) *response.Response {
	return response.Newf(401, 100100, CommonError_ErrUnauthorized.String(), "Unauthorized", args...)
}

func IsErrTokenInvalid(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrTokenInvalid.String() && e.GetCode() == 100101
}

func ErrTokenInvalid(args ...interface{}) *response.Response {
	return response.Newf(401, 100101, CommonError_ErrTokenInvalid.String(), "Token Invalid", args...)
}

func IsErrTokenExpiration(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrTokenExpiration.String() && e.GetCode() == 100102
}

func ErrTokenExpiration(args ...interface{}) *response.Response {
	return response.Newf(200, 100102, CommonError_ErrTokenExpiration.String(), "Token Expiration", args...)
}

func IsErrNoPartnerIdentity(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrNoPartnerIdentity.String() && e.GetCode() == 100110
}

func ErrNoPartnerIdentity(args ...interface{}) *response.Response {
	return response.Newf(200, 100110, CommonError_ErrNoPartnerIdentity.String(), "The account does not have partner identity", args...)
}

func IsErrNoMerchantIdentity(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrNoMerchantIdentity.String() && e.GetCode() == 100111
}

func ErrNoMerchantIdentity(args ...interface{}) *response.Response {
	return response.Newf(200, 100111, CommonError_ErrNoMerchantIdentity.String(), "The account does not have merchant identity", args...)
}

func IsErrNoBankerIdentity(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrNoBankerIdentity.String() && e.GetCode() == 100112
}

func ErrNoBankerIdentity(args ...interface{}) *response.Response {
	return response.Newf(200, 100112, CommonError_ErrNoBankerIdentity.String(), "The account does not have banker identity", args...)
}

func IsErrNoOrganizationIdentity(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrNoOrganizationIdentity.String() && e.GetCode() == 100113
}

func ErrNoOrganizationIdentity(args ...interface{}) *response.Response {
	return response.Newf(200, 100113, CommonError_ErrNoOrganizationIdentity.String(), "The account does not have organization identity", args...)
}

func IsErrPermissionDenied(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrPermissionDenied.String() && e.GetCode() == 100120
}

func ErrPermissionDenied(args ...interface{}) *response.Response {
	return response.Newf(200, 100120, CommonError_ErrPermissionDenied.String(), "Permission Denied", args...)
}

func IsErrResourceNotFound(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrResourceNotFound.String() && e.GetCode() == 100201
}

func ErrResourceNotFound(args ...interface{}) *response.Response {
	return response.Newf(404, 100201, CommonError_ErrResourceNotFound.String(), "Resource Not Found", args...)
}

func IsErrQueryConditionIsEmpty(err error) bool {
	e := response.FromError(err)
	return e.GetReason() == CommonError_ErrQueryConditionIsEmpty.String() && e.GetCode() == 100202
}

func ErrQueryConditionIsEmpty(args ...interface{}) *response.Response {
	return response.Newf(200, 100202, CommonError_ErrQueryConditionIsEmpty.String(), "Query Condition Is Empty", args...)
}

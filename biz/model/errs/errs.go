package errs

import "fmt"

type Error interface {
	Error() string
	Code() int32
	Msg() string
	SetErr(err error) Error
	SetMsg(msg string) Error
}

type bizError struct {
	code int32
	msg  string
}

func (bizErr *bizError) Error() string {
	return fmt.Sprintf("%d:%s", bizErr.code, bizErr.msg)
}

func (bizErr *bizError) Code() int32 {
	return bizErr.code
}

func (bizErr *bizError) Msg() string {
	return bizErr.msg
}

func (bizErr *bizError) SetErr(err error) Error {
	return New(bizErr.Code(), err.Error())
}

func (bizErr *bizError) SetMsg(msg string) Error {
	return New(bizErr.Code(), msg)
}

func New(code int32, msg string) Error {
	return &bizError{
		code: code,
		msg:  msg,
	}
}

func ErrorEqual(err1, err2 Error) bool {
	// 都为空
	if err1 == nil && err2 == nil {
		return true
	}

	// 只有一个不为空
	if err1 == nil || err2 == nil {
		return false
	}

	// 都不为空
	return err1.Code() == err2.Code()
}

var (
	Success         = New(0, "success")
	ServerError     = New(1_0001, "internal server error")
	ParamError      = New(1_0002, "param error")
	Unauthorized    = New(1_0003, "user unauthorized")
	TooManyRequest  = New(1_0004, "too many request")
	LoginReachLimit = New(1_0005, "login reach limit")
	RequestBlocked  = New(1_0006, "request is blocked")

	AccountNotExist      = New(2_0001, "account not exist or password incorrect")
	PasswordIncorrect    = AccountNotExist
	AccountStatusInvalid = New(2_0002, "account is invalid")

	DomainNameDuplicatedErr = New(3_0001, "domain name duplicated")
	DomainNotExistErr       = New(3_0002, "domain not exist")

	RoleNameDuplicatedErr        = New(4_0001, "role name duplicated")
	RoleNotExistErr              = New(4_0002, "role not exist")
	ParentRoleNotExistErr        = New(4_0003, "parent role not exist")
	ParentRoleNotInSameDomainErr = New(4_0004, "parent role not in same domain")

	ResourceNameDuplicatedErr = New(5_0001, "resource name duplicated")
	ResourceNotExistErr       = New(5_0002, "resource not exist")
)

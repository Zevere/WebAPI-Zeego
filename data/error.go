package data

type ErrorCode int

const (
	Default ErrorCode = iota
	NotFound
	InvalidValue
)

type Error interface {
	error
	Code() ErrorCode
}

type RepoError struct {
	Msg     string
	ErrCode ErrorCode
}

func NewRepoError(err error, code ErrorCode) *RepoError {
	return &RepoError{
		Msg:     err.Error(),
		ErrCode: code,
	}
}

func (e *RepoError) Error() string {
	return e.Msg
}

func (e *RepoError) Code() ErrorCode {
	return e.ErrCode
}

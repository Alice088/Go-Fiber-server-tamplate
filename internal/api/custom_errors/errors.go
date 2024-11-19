package custom_errors

import "errors"

var ErrKeyExpired = errors.New("key expired")
var ErrRepeatedRequest = errors.New("repeated request")
var ErrServerSideError = errors.New("server side error")

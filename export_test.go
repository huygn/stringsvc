// Bridge package to expose stringsvc internals to tests in the
// stringsvc_test package.

package stringsvc

func NewTestUppercaseRequest(s string) uppercaseRequest {
	return uppercaseRequest{S: s}
}

func NewTestUppercaseResponse(v, err string) uppercaseResponse {
	return uppercaseResponse{V: v, Err: err}
}

func NewTestCountRequest(s string) countRequest {
	return countRequest{S: s}
}

func NewTestCountResponse(v int) countResponse {
	return countResponse{V: v}
}

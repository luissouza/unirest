package unirest

import "net/http"

func makeRequest(method, urlStr string) *Request {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		panic(err)
	}

	return &Request{req: req}
}

// Get returns a request object for issuing HTTP GET requests.
func Get(urlStr string) *Request {
	return makeRequest("GET", urlStr)
}

// Head returns a request object for issuing HTTP HEAD requests.
func Head(urlStr string) *Request {
	return makeRequest("HEAD", urlStr)
}

// Post returns a request object for issuing HTTP POST requests.
func Post(urlStr string) *Request {
	return makeRequest("POST", urlStr)
}

// Patch returns a request object for issuing HTTP PATCH requests.
func Patch(urlStr string) *Request {
	return makeRequest("PATCH", urlStr)
}

// Delete returns a request object for issuing HTTP DELETE requests.
func Delete(urlStr string) *Request {
	return makeRequest("DELETE", urlStr)
}

// Request is a...
//
// TODO: Document the Request object.
type Request struct {
	req                                  *http.Request
	basicAuthUsername, basicAuthPassword string
	sendAuthImmediately                  bool
}

// Auth can be used for setting HTTP Basic authentication for a request.
//
// sendImmediately, if true, will add the HTTP Basic authentication to the
// request when the End function tries to send the request. However, if
// sendImmediately is false, then the End function will attempt the request
// without the HTTP Basic authentication, and if the server responds with an
// HTTP 401 Unauthorized status, then the request will be tried again, but with
// the HTTP Basic authentication set. Please note that in the latter case, the
// 401 response from the server must have a "WWW-Authenticate" header set,
// indicating the required authentication method.
func (r *Request) Auth(user, pass string, sendImmediately bool) *Request {
	r.basicAuthUsername = user
	r.basicAuthPassword = pass
	r.sendAuthImmediately = sendImmediately

	return r
}

// Header accepts a map[string]string of headers to set, for the request.
//
// If you would like to use HTTP Basic authentication for the request, please
// use the Auth method, instead of setting it manually.
func (r *Request) Header(hdr map[string]string) *Request {
	for h, v := range hdr {
		if _, ok := r.Header[h]; !ok {
			r.Header[h] = make([]string, 0)
		}

		r.Header[h] = append(r.Header[h], v)
	}

	return r
}

// Set is analogous to the Header method.
//
// It does nothing different than Header, simply because it just calls the
// Header method with the provided map[string]string.
//
// This method is defined only for consistency with the Unirest for Node.js
// documentation.
func (r *Request) Set(hdr map[string]string) *Request {
	return r.Header(hdr)
}

// Query appends the provided parameters to the URL.
func (r *Request) Query(params map[string]string) *Request {
	for k, v := range params {
		r.req.URL.Query().Add(k, v)
	}

	return r
}

// Send adds data to the request body.
//
// TODO: finish
func (r *Request) Send(v interface{}) *Request {
	return r
}

func (r *Request) Type(ctyp string) *Request {
	typ := "application/x-www-form-urlencoded"
	switch ctyp {
	case "json":
		typ = "application/json"
	case "xml":
		typ = "application/xml"
	default:

	}

	r.req.Header.Set("Content-Type", typ)

	return r
}

func (r *Request) Field(param, val string) *Request {}

func (r *Request) Attach(name, pth string) *Request {}

func (r *Request) End(fn func(*http.Response) error) error {}

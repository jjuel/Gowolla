package dwolla

type ErrorResp struct {
	Code     string   `json:"code"`
	Message  string   `json:"message"`
	Embedded Embedded `json:"_embedded"`
}

type Embedded struct {
	Errors []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Path    string `json:"path"`
}

type Client struct {
}

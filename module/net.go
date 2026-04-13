package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/vintlang/vintlang/object"
)

var NetFunctions = map[string]object.ModuleFunction{}

func init() {
	NetFunctions["get"] = getRequest
	NetFunctions["post"] = postRequest
	NetFunctions["put"] = putRequest
	NetFunctions["delete"] = deleteRequest
	NetFunctions["patch"] = patchRequest
	NetFunctions["fetch"] = fetchRequest
	// NetFunctions["http"] = httpServer
}

func deleteRequest(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return handleRequest("DELETE", args, defs)
}

func patchRequest(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return handleRequest("PATCH", args, defs)
}

func handleRequest(method string, args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	var url *object.String
	var headers, params *object.Dict

	for k, v := range defs {
		switch k {
		case "url":
			strUrl, ok := v.(*object.String)
			if !ok {
				return ErrorMessage(
					"net", method,
					"string value for 'url' parameter",
					string(v.Type()),
					fmt.Sprintf(`net.%s(url="https://example.com")`, strings.ToLower(method)),
				)
			}
			url = strUrl
		case "headers":
			dictHead, ok := v.(*object.Dict)
			if !ok {
				return ErrorMessage(
					"net", method,
					"dictionary value for 'headers' parameter",
					string(v.Type()),
					fmt.Sprintf(`net.%s(headers={"Content-Type": "application/json"})`, strings.ToLower(method)),
				)
			}
			headers = dictHead
		case "body":
			dictBody, ok := v.(*object.Dict)
			if !ok {
				return ErrorMessage(
					"net", method,
					"dictionary value for 'body' parameter",
					string(v.Type()),
					fmt.Sprintf(`net.%s(body={"key": "value"})`, strings.ToLower(method)),
				)
			}
			params = dictBody
		default:
			return &object.Error{
				Message: fmt.Sprintf("\033[1;31m -> net.%s()\033[0m:\n"+
					"  Invalid parameter '%s'.\n"+
					"  Valid parameters are: 'url', 'headers', 'body'.\n"+
					"  Usage: net.%s(url=\"https://example.com\", headers={...}, body={...})\n",
					method, k, strings.ToLower(method)),
			}
		}
	}

	if url == nil || url.Value == "" {
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31m -> net.%s()\033[0m:\n"+
				"  Missing required 'url' parameter.\n"+
				"  Please provide a valid URL for the HTTP request.\n"+
				"  Usage: net.%s(url=\"https://example.com\") or net.%s(\"https://example.com\")\n",
				method, strings.ToLower(method), strings.ToLower(method)),
		}
	}

	var requestBody *bytes.Buffer
	if params != nil {
		bodyContent := convertObjectToWhatever(params)
		jsonBody, err := json.Marshal(bodyContent)
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("net.%s() failed to serialize request body to JSON: %v. Please ensure your body contains valid JSON-serializable data.", method, err)}
		}
		requestBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url.Value, requestBody)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("net.%s() failed to create HTTP request to '%s': %v. Please check if the URL is valid and properly formatted.", method, url.Value, err)}
	}

	if headers != nil {
		for _, val := range headers.Pairs {
			req.Header.Set(val.Key.Inspect(), val.Value.Inspect())
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("net.%s() failed to execute HTTP request to '%s': %v. Please check your internet connection and ensure the server is accessible.", method, url.Value, err)}
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("net.%s() failed to read HTTP response from '%s': %v. The connection may have been interrupted.", method, url.Value, err)}
	}

	return &object.String{Value: string(respBody)}
}

func getRequest(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		var url *object.String
		var headers, params *object.Dict
		for k, v := range defs {
			switch k {
			case "url":
				strUrl, ok := v.(*object.String)
				if !ok {
					return &object.Error{Message: "URL must be a string"}
				}
				url = strUrl
			case "headers":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Headers must be a dictionary"}
				}
				headers = dictHead
			case "body":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Body must be a dictionary"}
				}
				params = dictHead
			default:
				return &object.Error{Message: "Arguments are incorrect. Use url and headers."}
			}
		}
		if url.Value == "" {
			return &object.Error{Message: "URL is required"}
		}

		var responseBody *bytes.Buffer
		if params != nil {
			booty := convertObjectToWhatever(params)

			jsonBody, err := json.Marshal(booty)

			if err != nil {
				return &object.Error{Message: "Your query is not formatted properly."}
			}

			responseBody = bytes.NewBuffer(jsonBody)
		}

		var req *http.Request
		var err error
		if responseBody != nil {
			req, err = http.NewRequest("GET", url.Value, responseBody)
		} else {
			req, err = http.NewRequest("GET", url.Value, nil)
		}
		if err != nil {
			return &object.Error{Message: "Failed to make the request"}
		}

		if headers != nil {
			for _, val := range headers.Pairs {
				req.Header.Set(val.Key.Inspect(), val.Value.Inspect())
			}
		}
		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Failed to send the request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Failed to read the response."}
		}

		return &object.String{Value: string(respBody)}
	}

	if len(args) == 1 {
		url, ok := args[0].(*object.String)
		if !ok {
			return &object.Error{Message: "URL must be a string"}
		}
		req, err := http.NewRequest("GET", url.Value, nil)
		if err != nil {
			return &object.Error{Message: "Failed to make the request"}
		}

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Failed to send the request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Failed to read the response."}
		}

		return &object.String{Value: string(respBody)}
	}
	return &object.Error{Message: "Arguments are incorrect. Use url and headers."}
}

func postRequest(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		var url *object.String
		var headers, params *object.Dict
		for k, v := range defs {
			switch k {
			case "url":
				strUrl, ok := v.(*object.String)
				if !ok {
					return &object.Error{Message: "URL must be a string"}
				}
				url = strUrl
			case "headers":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Headers must be a dictionary"}
				}
				headers = dictHead
			case "body":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Body must be a dictionary"}
				}
				params = dictHead
			default:
				return &object.Error{Message: "Arguments are incorrect. Use url and headers."}
			}
		}
		if url.Value == "" {
			return &object.Error{Message: "URL is required"}
		}
		var responseBody *bytes.Buffer
		if params != nil {
			booty := convertObjectToWhatever(params)

			jsonBody, err := json.Marshal(booty)

			if err != nil {
				return &object.Error{Message: "Your query is not formatted properly."}
			}

			responseBody = bytes.NewBuffer(jsonBody)
		}
		var req *http.Request
		var err error
		if responseBody != nil {
			req, err = http.NewRequest("POST", url.Value, responseBody)
		} else {
			req, err = http.NewRequest("POST", url.Value, nil)
		}
		if err != nil {
			return &object.Error{Message: "Failed to make the request"}
		}
		if headers != nil {
			for _, val := range headers.Pairs {
				req.Header.Set(val.Key.Inspect(), val.Value.Inspect())
			}
		}
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Failed to send the request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Failed to read the response."}
		}
		return &object.String{Value: string(respBody)}
	}
	return &object.Error{Message: "Arguments are incorrect. Use url and headers."}
}

func putRequest(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		var url *object.String
		var headers, params *object.Dict
		for k, v := range defs {
			switch k {
			case "url":
				strUrl, ok := v.(*object.String)
				if !ok {
					return &object.Error{Message: "URL must be a string"}
				}
				url = strUrl
			case "headers":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Headers must be a dictionary"}
				}
				headers = dictHead
			case "body":
				dictHead, ok := v.(*object.Dict)
				if !ok {
					return &object.Error{Message: "Body must be a dictionary"}
				}
				params = dictHead
			default:
				return &object.Error{Message: "Arguments are incorrect. Use url and headers."}
			}
		}
		if url.Value == "" {
			return &object.Error{Message: "URL is required"}
		}
		var responseBody *bytes.Buffer
		if params != nil {
			booty := convertObjectToWhatever(params)

			jsonBody, err := json.Marshal(booty)

			if err != nil {
				return &object.Error{Message: "Your query is not formatted properly."}
			}

			responseBody = bytes.NewBuffer(jsonBody)
		}
		var req *http.Request
		var err error
		if responseBody != nil {
			req, err = http.NewRequest("PUT", url.Value, responseBody)
		} else {
			req, err = http.NewRequest("PUT", url.Value, nil)
		}
		if err != nil {
			return &object.Error{Message: "Failed to make the request"}
		}
		if headers != nil {
			for _, val := range headers.Pairs {
				req.Header.Set(val.Key.Inspect(), val.Value.Inspect())
			}
		}
		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)

		if err != nil {
			return &object.Error{Message: "Failed to send the request."}
		}
		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &object.Error{Message: "Failed to read the response."}
		}
		return &object.String{Value: string(respBody)}
	}
	return &object.Error{Message: "Arguments are incorrect. Use url and headers."}
}

// fetchRequest performs an HTTP request and returns a Dict with status, headers, and body.
// It supports string body (not just dict), making it suitable for proxying requests.
// Usage: net.fetch(method="GET", url="https://example.com", headers={...}, body="raw string or dict")
func fetchRequest(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
var urlStr string
var method string
var headers *object.Dict
var bodyStr string
var hasBody bool

for k, v := range defs {
switch k {
case "url":
strVal, ok := v.(*object.String)
if !ok {
return &object.Error{Message: "net.fetch(): 'url' must be a string"}
}
urlStr = strVal.Value
case "method":
strVal, ok := v.(*object.String)
if !ok {
return &object.Error{Message: "net.fetch(): 'method' must be a string"}
}
method = strings.ToUpper(strVal.Value)
case "headers":
dictVal, ok := v.(*object.Dict)
if !ok {
return &object.Error{Message: "net.fetch(): 'headers' must be a dictionary"}
}
headers = dictVal
case "body":
hasBody = true
switch bv := v.(type) {
case *object.String:
bodyStr = bv.Value
case *object.Dict:
bodyContent := convertObjectToWhatever(bv)
jsonBody, err := json.Marshal(bodyContent)
if err != nil {
return &object.Error{Message: fmt.Sprintf("net.fetch(): failed to serialize body: %v", err)}
}
bodyStr = string(jsonBody)
default:
bodyStr = v.Inspect()
}
default:
return &object.Error{
Message: fmt.Sprintf("net.fetch(): unknown parameter '%s'. Valid: 'method', 'url', 'headers', 'body'", k),
}
}
}

if urlStr == "" {
return &object.Error{Message: "net.fetch(): 'url' parameter is required"}
}
if method == "" {
method = "GET"
}

var reqBody *bytes.Buffer
if hasBody {
reqBody = bytes.NewBufferString(bodyStr)
}

var req *http.Request
var err error
if reqBody != nil {
req, err = http.NewRequest(method, urlStr, reqBody)
} else {
req, err = http.NewRequest(method, urlStr, nil)
}
if err != nil {
return &object.Error{Message: fmt.Sprintf("net.fetch(): failed to create request: %v", err)}
}

if headers != nil {
for _, val := range headers.Pairs {
req.Header.Set(val.Key.Inspect(), val.Value.Inspect())
}
}

client := &http.Client{Timeout: 30 * time.Second}
resp, err := client.Do(req)
if err != nil {
return &object.Error{Message: fmt.Sprintf("net.fetch(): request failed: %v", err)}
}
defer resp.Body.Close()

respBody, err := ioutil.ReadAll(resp.Body)
if err != nil {
return &object.Error{Message: fmt.Sprintf("net.fetch(): failed to read response: %v", err)}
}

// Build response headers dict
respHeaderPairs := make(map[object.HashKey]object.DictPair)
for key, vals := range resp.Header {
k := &object.String{Value: key}
v := &object.String{Value: strings.Join(vals, ", ")}
respHeaderPairs[k.HashKey()] = object.DictPair{Key: k, Value: v}
}

// Build result dict with status, headers, body
statusKey := &object.String{Value: "status"}
headersKey := &object.String{Value: "headers"}
bodyKey := &object.String{Value: "body"}

resultPairs := make(map[object.HashKey]object.DictPair)
resultPairs[statusKey.HashKey()] = object.DictPair{
Key:   statusKey,
Value: &object.Integer{Value: int64(resp.StatusCode)},
}
resultPairs[headersKey.HashKey()] = object.DictPair{
Key:   headersKey,
Value: &object.Dict{Pairs: respHeaderPairs},
}
resultPairs[bodyKey.HashKey()] = object.DictPair{
Key:   bodyKey,
Value: &object.String{Value: string(respBody)},
}

return &object.Dict{Pairs: resultPairs}
}

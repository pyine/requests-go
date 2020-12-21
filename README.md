requests-go
=======
A simple `HTTP Request` package for golang. `GET` `POST` `DELETE` `PUT` `HEAD` `Upload`



### Installation
go get github.com/pyine/requests-go


### How do we use requests?

#### Create request object use http.DefaultTransport
```go
req := requests.NewRequest()
req := requests.NewRequest().Debug(true).SetTimeout(5)
```

#### Set headers
```go
req.SetHeaders(map[string]string{
    "Content-Type": "application/x-www-form-urlencoded",
    "Connection": "keep-alive",
})

req.SetHeaders(map[string]string{
    "Source":"api",
})
```

#### Set basic auth
```go
req.SetBasicAuth("username","password")
```

#### Set oauth
```go
req.SetOAuth("token")
```

#### Set timeout
```go
req.SetTimeout(5)  //default 30s
```

#### Transport
If you want to customize the Client object and reuse TCP connections, you need to define a global http.RoundTripper or & http.Transport, because the reuse of the http connection pool is based on Transport.
```go
var transport *http.Transport
func init() {   
    transport = &http.Transport{
        MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
    }
}

func demo(){
    // Use http.DefaultTransport
    res, err := requests.Get("http://127.0.0.1:8080")
    // Use custom Transport
    res, err := requests.Transport(transport).Get("http://127.0.0.1:8080")
}
```

#### Ignore Https certificate validation，Only effective for custom Transport
```go
req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
res, err := requests.Transport(transport).SetTLSClient(&tls.Config{InsecureSkipVerify: true}).Get("http://127.0.0.1:8080")
```

#### Object-oriented operation mode
```go
req := requests.NewRequest().
	Debug(true).
	SetHeaders(map[string]string{
	    "Content-Type": "application/x-www-form-urlencoded",
	}).SetTimeout(5)
res, err := requests.NewRequest().Get("http://127.0.0.1")
```

### GET

#### Query parameter
```go
res, err := req.Get("http://127.0.0.1:8000")
res, err := req.Get("http://127.0.0.1:8000?id=10&title=requests")
res, err := req.Get("http://127.0.0.1:8000?id=10&title=requests",nil)
res, err := req.Get("http://127.0.0.1:8000?id=10&title=requests","address=beijing")

res, err := requests.Get("http://127.0.0.1:8000")
res, err := requests.Debug(true).SetHeaders(map[string]string{}).Get("http://127.0.0.1:8000")
```


#### Multi parameter
```go
res, err := req.Get("http://127.0.0.1:8000?id=10&title=requests",map[string]interface{}{
    "name":  "jason",
    "score": 100,
})

body, err := res.Body()
if err != nil {
    return
}

return string(body)
```


### POST

```go
res, err := req.Post("http://127.0.0.1:8000")
res, err := req.Post("http://127.0.0.1:8000", "title=github&type=1")
res, err := req.JSON().Post("http://127.0.0.1:8000", "{\"id\":10,\"title\":\"requests\"}")
res, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "requests",
})
body, err := res.Body()
if err != nil {
    return
}
return string(body)

res, err := requests.Post("http://127.0.0.1:8000")
res, err := requests.JSON().Post("http://127.0.0.1:8000",map[string]interface{}{"title":"github"})
res, err := requests.Debug(true).SetHeaders(map[string]string{}).JSON().Post("http://127.0.0.1:8000","{\"title\":\"github\"}")
```

### Upload
Params: url, filename, fileinput

```go
res, err := req.Upload("http://127.0.0.1:8000/upload", "/root/example.txt","uploadFile")
body, err := res.Body()
if err != nil {
    return
}
return string(body)
```


### Debug
#### Default false

```go
req.Debug(true)
```

#### Print in standard output：
```go
[HttpRequest]
-------------------------------------------------------------------
Request: GET http://127.0.0.1:8000?name=jobs
Headers: map[Content-Type:application/json]
Elapsed time:: 100 ms
Timeout: 30s
ReqBody: map[age:19 score:100]
-------------------------------------------------------------------
```


## Json
Post JSON request

#### Set header
```go
 req.SetHeaders(map[string]string{"Content-Type": "application/json"})
```
Or
```go
req.JSON().Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "github",
})

req.JSON().Post("http://127.0.0.1:8000", "{\"title\":\"github\",\"id\":10}")
```

#### Post request
```go
res, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "requests",
})
```

#### Unmarshal JSON
```go
var u User
err := res.Json(&u)
if err != nil {
   return err
}

var m map[string]interface{}
err := res.Json(&m)
if err != nil {
   return err
}
```

### Response

#### Response() *http.Response
```go
res, err := req.Post("http://127.0.0.1:8000/") //res is a http.Response object
```

#### StatusCode() int
```go
res.StatusCode()
```

#### Body() ([]byte, error)
```go
body, err := res.Body()
log.Println(string(body))
```

#### Close() error
```go
res.Close()
```

#### Time() string
```go
res.Time()  //ms
```

#### Print formatted JSON
```go
str, err := res.Export()
if err != nil {
   return
}
```

#### Unmarshal JSON
```go
var u User
err := res.Json(&u)
if err != nil {
   return err
}

var m map[string]interface{}
err := res.Json(&m)
if err != nil {
   return err
}
```

#### Url() string
```go
res.Url()  //return the requested url
```

#### Headers() http.Header
```go
res.Headers()  //return the response headers
```

### Advanced
#### GET
```go
import "git.zoom.us/op-infra-k8s/gocommon/requests"
   
res,err := requests.Get("http://127.0.0.1:8000/")
res,err := requests.Get("http://127.0.0.1:8000/","title=github")
res,err := requests.Get("http://127.0.0.1:8000/?title=github")
res,err := requests.Debug(true).JSON().Get("http://127.0.0.1:8000/")
```

#### POST
```go
import "git.zoom.us/op-infra-k8s/gocommon/requests"
   
res,err := requests.Post("http://127.0.0.1:8000/")
res,err := requests.SetHeaders(map[string]string{
	"title":"github",
}).Post("http://127.0.0.1:8000/")
res,err := requests.Debug(true).JSON().Post("http://127.0.0.1:8000/")
```


### Example
```go
import "git.zoom.us/op-infra-k8s/gocommon/requests"
   
res,err := requests.Get("http://127.0.0.1:8000/")
res,err := requests.Get("http://127.0.0.1:8000/","title=github")
res,err := requests.Get("http://127.0.0.1:8000/?title=github")
res,err := requests.Get("http://127.0.0.1:8000/",map[string]interface{}{
	"title":"github",
})
res,err := requests.Debug(true).JSON().SetHeaders(map[string]string{
	"source":"api",
}).Post("http://127.0.0.1:8000/")


//Or
req := requests.NewRequest()
req := req.Debug(true).SetHeaders()
res,err := req.Debug(true).JSON().SetHeaders(map[string]string{
    "source":"api",
}).Post("http://127.0.0.1:8000/")
```


# tests
    import "github.com/ardanlabs/kit/tests"

Package tests provides the generic support all tests require.





## Variables
``` go
var (
    Success = "\u2713"
    Failed  = "\u2717"
)
```
Success and failure markers.

``` go
var Context = "Test"
```
Context provides a base context for tests.


## func DisplayLog
``` go
func DisplayLog()
```
DisplayLog writes the logdash data to standand out, if testing in verbose mode
was turned on.


## func Init
``` go
func Init(cfgKey string)
```
Init initializes the log package.


## func InitMongo
``` go
func InitMongo(cfg mongo.Config)
```
InitMongo initializes the mongodb connections for testing.


## func NewRequest
``` go
func NewRequest(method, path string, body io.Reader) *http.Request
```
NewRequest used to setup a request for mocking API calls with httptreemux.


## func ResetLog
``` go
func ResetLog()
```
ResetLog resets the contents of logdash.









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
# Echo

https://echo.labstack.com/


## server
```
mkdir myapp && cd myapp
```{{execute}}

```
pwd;
ls;
```{{execute}}

```
go mod init myapp
```{{execute}}

```
go get github.com/labstack/echo/v4
```{{execute}}

```
cat << 'EOF' > server.go
package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

EOF

```{{execute}}



```
go run server.go
```{{execute}}


https://[[HOST_SUBDOMAIN]]-1323-[[KATACODA_HOST]].environments.katacoda.com/


https://echo.labstack.com/guide/



## Prometheus


https://echo.labstack.com/middleware/prometheus/



```
cat << 'EOF' > server-prom.go
package main
import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo-contrib/prometheus"
)
func main() {
    e := echo.New()
    // Enable metrics middleware
    p := prometheus.NewPrometheus("echo", nil)
    p.Use(e)

    e.Logger.Fatal(e.Start(":1323"))
}
EOF

```{{execute}}



```
go run server-prom.go
```{{execute}}


terminal 2
```
echo t2
```{{execute T2}}

```
curl http://localhost:1323/metrics
```{{execute T2}}

https://[[HOST_SUBDOMAIN]]-1323-[[KATACODA_HOST]].environments.katacoda.com/metrics


## jaegertracing

https://echo.labstack.com/middleware/jaegertracing/


TraceFunction
This is a wrapper function that can be used to seamlessly add a span for the duration of the invoked function. 
There is no need to change function arguments.

```
cat << 'EOF' > server-jaeger.go
package main
import (
    "github.com/labstack/echo-contrib/prometheus"
    "github.com/labstack/echo-contrib/jaegertracing"
    "github.com/labstack/echo/v4"
    "net/http"
    "time"
)
func main() {
    e := echo.New()
    
    // Enable metrics middleware
    p := prometheus.NewPrometheus("echo", nil)
    p.Use(e)

    // Enable tracing middleware
    c := jaegertracing.New(e, nil)
    defer c.Close()
    e.GET("/", func(c echo.Context) error {
        // Wrap slowFunc on a new span to trace it's execution passing the function arguments
		jaegertracing.TraceFunction(c, slowFunc, "Test String")
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}

// A function to be wrapped. No need to change it's arguments due to tracing
func slowFunc(s string) {
	time.Sleep(200 * time.Millisecond)
	return
}
EOF

```{{execute}}



```
go run server-jaeger.go
```{{execute}}



ref:

https://opentelemetry.io/

https://openmetrics.io/

https://opentracing.io/docs/overview/

https://prometheus.io/


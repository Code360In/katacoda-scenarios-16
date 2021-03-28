# Echo

https://echo.labstack.com/


```
mkdir myapp && cd myapp
```{{execute}}

```
go mod init myapp
```{{execute}}

```
go get github.com/labstack/echo/v4
```{{execute}}


server.go

```
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

```{{copy}}


```
go run server.go
```{{execute}}

https://echo.labstack.com/guide/



# Prometheus

https://echo.labstack.com/middleware/prometheus/


server-prom.go
```
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
```{{copy}}

```
go run server-prom.go
```{{execute}}

```
curl http://localhost:1323/metrics
```{{execute}}



# jaegertracing

https://echo.labstack.com/middleware/jaegertracing/



server-jaeger.go
```
package main
import (
    "github.com/labstack/echo-contrib/jaegertracing"
    "github.com/labstack/echo/v4"
)
func main() {
    e := echo.New()
    // Enable tracing middleware
    c := jaegertracing.New(e, nil)
    defer c.Close()

    e.Logger.Fatal(e.Start(":1323"))
}
```{{copy}}



```
go run server-jaeger.go
```{{execute}}

// Обнаружение сетевого тайм-аута с помощью ошибки
package main

import (
    "context"
    "errors"
    "net"
    "net/url"
    "os"
    "strings"
	"net/http"
)

func hasTimedOut(err error) bool {
    if err == nil {
        return false
    }

    if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, os.ErrDeadlineExceeded) {
        return true
    }

    switch e := err.(type) {
    case *url.Error:
        if netErr, ok := e.Err.(net.Error); ok && netErr.Timeout() {
            return true
        }
    case net.Error:
        if e.Timeout() {
            return true
        }
    case *net.OpError:
        if e.Timeout() {
            return true
        }
    }

    type timeouter interface {
        Timeout() bool
    }
    if te, ok := err.(timeouter); ok && te.Timeout() {
        return true
    }

    commonPhrases := []string{
        "use of closed network connection",
        "connection reset by peer",
        "i/o timeout",
        "deadline exceeded",
        "timeout",
        "timed out",
        "temporary failure",
    }
    errStr := err.Error()
    for _, phrase := range commonPhrases {
        if strings.Contains(strings.ToLower(errStr), phrase) {
            return true
        }
    }

    return false
}

func main() {
    res, err := http.Get("http://example.com/test.zip")
    if hasTimedOut(err) {
        panic("request has timed out")
    }
    if err != nil {
        panic("something else has happened: " + err.Error())
    }
    defer res.Body.Close()
}
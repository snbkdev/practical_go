// Функция загрузки данных, предусматривающая повторные попытки
package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Create("file.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	location := "https://example.com/file.zip"
	err = downlaoad(location, file, 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Got it with %v bytes downloaded", fi.Size())
}

func downlaoad(location string, file *os.File, retries int64) error {
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return err
	}

	fi, err := file.Stat()
	if err != nil {
		return err
	}

	current := fi.Size()
	if current > 0 {
		start := strconv.FormatInt(current, 10)
		req.Header.Set("Range", "bytes="+start+"-")
	}

	cc := &http.Client{Timeout: 5 * time.Minute}
	res, err := cc.Do(req)

	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return downlaoad(location, file, retries-1)
		}
		return err
	}

	if res.StatusCode < 200 || res.StatusCode > 300 {
		errFmt := "Unsuccessful HTTP request. Status %s"
		return fmt.Errorf(errFmt, res.Status)
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}

	_, err = io.Copy(file, req.Body)
	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return downlaoad(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}

	return nil
}

func hasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}

	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}

	return false
}
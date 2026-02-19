// Обработка двух разных ошибок
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const MAX_TIMEOUTS = 5

var ErrTimeOut = errors.New("the request timed out")
var ErrRejected = errors.New("the request was rejected")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	response, err := SendRequest("Hello")

	if errors.Is(err, ErrTimeOut) {
		timeouts := 0
		for errors.Is(err, ErrTimeOut) {
			timeouts++
			fmt.Println("Timeout. Retrying")
			if timeouts == MAX_TIMEOUTS {
				panic("too many timeouts")
			}

			response, err = SendRequest("Hello")
		}
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func SendRequest(req string) (string, error) {
	switch rand.Intn(3) % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeOut
	}
}
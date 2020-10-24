package processor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/eapache/go-resiliency/retrier"
	"github.com/pkg/errors"
	"github.com/wesleimp/dc-test/pkg/order"
)

var url = "https://delivery-center-recruitment-ap.herokuapp.com/"

// Process the next request
func Process(order *order.Order) (string, int, error) {
	payload := toPayload(order)

	var body string
	var status int
	var err error

	r := retrier.New(retrier.ConstantBackoff(3, 1*time.Second), nil)
	err = r.Run(func() error {
		body, status, err = doRequest(payload, time.Now())
		if err != nil {
			return err
		}

		return nil
	})

	return body, status, err
}

func doRequest(payload *Payload, now time.Time) (string, int, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return "", 0, errors.Wrap(err, "error parsing payload")
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return "", 0, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-Sent", formatRequestDate(now))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", 0, err
	}
	defer response.Body.Close()

	bts, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", 0, err
	}

	return string(bts), response.StatusCode, nil
}

func formatRequestDate(date time.Time) string {
	return fmt.Sprintf("%02dh%d - %d/%d/%d", date.Hour(), date.Minute(), date.Day(), date.Month(), date.Year())
}

package job_schedule

import (
	"fmt"
	"time"

	"errors"

	"github.com/gunsluo/gorequest"
)

func agent() *gorequest.SuperAgent {
	to := time.Duration(10000)
	return gorequest.New().
		Timeout(to * time.Millisecond)
}

func send(url, data string) (resp string, err error) {
	_, body, errs := agent().Get(url).Send(data).End()
	err = errors_to_error(errs)
	if err != nil {
		return
	}
	resp = body

	return
}

func sendPost(url, data string) (resp string, err error) {

	_, body, errs := agent().Post(url).
		Set("Content-Type", "application/json").
		Send(data).End()
	err = errors_to_error(errs)
	if err != nil {
		return
	}
	resp = body

	return
}

func sendDelete(url, data string) (resp string, err error) {

	_, body, errs := agent().Delete(url).
		Set("Content-Type", "application/json").
		Send(data).End()
	err = errors_to_error(errs)
	if err != nil {
		return
	}
	resp = body

	return
}

func errors_to_error(errs []error) error {
	if errs == nil || len(errs) == 0 {
		return nil
	}

	errString := ""
	for i, e := range errs {
		if i == 0 {
			errString = errString + fmt.Sprintf("error %d: %s ", i, e.Error())
		} else {
			errString = errString + fmt.Sprintf("\nerror %d: %s ", i, e.Error())
		}
	}

	return errors.New(errString)
}

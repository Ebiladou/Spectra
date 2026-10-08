package http

import (
	"context"
	"errors"
	"fmt"
	"io"
	nethttp "net/http"
	"strings"
	"time"

	"github.com/Ebiladou/spectra/internal/model"
)

const defaultTimeout = 30 * time.Second

type HTTPRunner interface {
	Execute(
		ctx context.Context,
		step *model.Step,
		execution *model.ExecutionContext,
	) (*model.Sample, error)
}

type httpRunner struct {
	client *nethttp.Client
}

func NewHTTPRunner() HTTPRunner {
	return &httpRunner{
		client: &nethttp.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (runner *httpRunner) Execute(
	ctx context.Context,
	step *model.Step,
	execution *model.ExecutionContext,
) (*model.Sample, error) {
	if step == nil {
		return nil, errors.New("step cannot be empty")
	}

	if execution == nil {
		return nil, errors.New("execution context cannot be empty")
	}

	sample := &model.Sample{
		ScenarioID:    execution.ScenarioID,
		StepID:        step.ID,
		VirtualUserID: execution.VirtualUserID,
	}

	request, err := nethttp.NewRequestWithContext(
		ctx,
		step.Method,
		step.URL,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("create HTTP request: %w", err)
	}

	if step.Data != "" {
		request.Body = io.NopCloser(
			strings.NewReader(step.Data),
		)
	}

	sample.Timestamp = time.Now()
	start := time.Now()

	response, err := runner.client.Do(request)

	if err != nil {
		sample.Duration = time.Since(start)
		sample.Error = err
		sample.Success = false

		return sample, nil
	}

	defer response.Body.Close()

	_, bodyErr := io.Copy(io.Discard, response.Body)

	sample.Duration = time.Since(start)
	sample.StatusCode = response.StatusCode

	if bodyErr != nil {
		sample.Error = fmt.Errorf("read response body: %w", bodyErr)
		sample.Success = false

		return sample, nil
	}

	sample.Success = response.StatusCode >= 200 &&
		response.StatusCode < 300

	return sample, nil
}

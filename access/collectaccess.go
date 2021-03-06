package access

import (
	"errors"
	"time"

	"github.com/gesiel/gocollector/utils"
)

var MissingClientIdError = errors.New("Access missing field: ClientId")
var MissingPathError = errors.New("Access missing field: Path")

type CollectAccessUseCase struct {
	Gateway Gateway
}

func (this *CollectAccessUseCase) Collect(input CollectAccessInput) (*CollectAccessResponse, error) {
	err := validateInput(input)
	if err != nil {
		return nil, err
	}

	access := createAccessFor(input)
	err = this.Gateway.Save(access)
	if err != nil {
		return nil, err
	}

	return &CollectAccessResponse{
		Access: access,
	}, nil
}

type CollectAccessResponse struct {
	Access *Access
}

type CollectAccessInput interface {
	GetClientId() string
	GetPath() string
	GetDate() time.Time
}

func validateInput(input CollectAccessInput) error {
	if !utils.IsValidValue(input.GetClientId()) {
		return MissingClientIdError
	}

	if !utils.IsValidValue(input.GetPath()) {
		return MissingPathError
	}

	return nil
}

func createAccessFor(input CollectAccessInput) *Access {
	access := &Access{
		ClientId: input.GetClientId(),
		Path:     input.GetPath(),
		Date:     input.GetDate(),
	}
	return access
}

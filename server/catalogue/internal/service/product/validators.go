package product

import (
	"fmt"
	"github.com/omaqase/sato/catalogue/internal/service"
	protobuf "github.com/omaqase/sato/catalogue/pkg/api/v1/catalogue"
	"slices"
)

var ApproveStatus = []string{"pending", "approved", "rejected"}

func ValidateApproveStatus(status string) error {
	if !slices.Contains(ApproveStatus, status) {
		return fmt.Errorf("invalid approve status")
	}

	return nil
}

func ValidateCreateProductRequest(request *protobuf.CreateProductRequest) error {
	if err := service.ValidateUUID(request.CategoryId); err != nil {
		return err
	}

	if request.Stock < 0 {
		return fmt.Errorf("invalid stock")
	}

	if len(request.Title) <= 0 {
		return fmt.Errorf("invalid title")
	}

	return nil
}

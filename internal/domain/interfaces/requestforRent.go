// internal/domain/interfaces/request_repo.go
package interfaces

import (
	"context"
	"rentease/internal/domain/entities"
)

type RequestRepo interface {
	SaveRequest(request entities.Request) error
	FindByTenantUsername(ctx context.Context, tenantUsername string) ([]entities.Request, error)
	FindByLandlordName(ctx context.Context, landlordName string) ([]entities.Request, error)
	UpdateRequest(request entities.Request, status string) error
}

//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/weaviate/weaviate/entities/models"
	pb "github.com/weaviate/weaviate/grpc/generated/protocol/v1"
)

func TestGRPCTenants(t *testing.T) {
	tests := []struct {
		activityStatusGRPC pb.TenantActivityStatus
		activityStatus     string
	}{
		{
			activityStatusGRPC: pb.TenantActivityStatus_ACTIVITY_STATUS_HOT,
			activityStatus:     "HOT",
		},
		{
			activityStatusGRPC: pb.TenantActivityStatus_ACTIVITY_STATUS_COLD,
			activityStatus:     "COLD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.activityStatus, func(t *testing.T) {
			tenantGRPC, err := tenantToGRPC(&models.Tenant{
				Name:           "TestTenant",
				ActivityStatus: tt.activityStatus,
			})
			require.Nil(t, err)
			require.Equal(t, "TestTenant", tenantGRPC.GetName())
			require.Equal(t, tt.activityStatusGRPC, tenantGRPC.GetActivityStatus())
		})
	}
}

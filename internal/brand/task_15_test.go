package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask15(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	parent := IngredientLot{ID: "p", ProducedAt: now.Add(-2 * time.Hour)}
	lot := IngredientLot{ID: "c", SupplierID: "sup", CertificateID: "cert", OriginRegion: "青海", ProducedAt: now.Add(-time.Hour), ExpiresAt: now.Add(time.Hour), ParentLotIDs: []string{"p"}}
	require.NoError(t, s.CheckLotTrace(context.Background(), lot, map[string]IngredientLot{"p": parent}))
}

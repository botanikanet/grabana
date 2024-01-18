package encoder

import (
	"github.com/botanikanet/grabana/encoder/golang"
	"github.com/botanikanet/sdk"
	"go.uber.org/zap"
)

func ToGolang(logger *zap.Logger, dashboard sdk.Board) (string, error) {
	golangEncoder := golang.NewEncoder(logger)

	return golangEncoder.EncodeDashboard(dashboard)
}

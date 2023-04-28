package lo_samples

import (
	"github.com/samber/lo"
	"go.uber.org/zap"
	"testing"
)

func TestHelloUniq(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	names := lo.Uniq[string]([]string{"Samuel", "John", "Samuel"})

	for _, name := range names {
		sugar.Infof("Real Name: %s", name)
	}
}

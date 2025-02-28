package handler_test

import (
	"github.com/Korpenter/tusd/v2/pkg/filestore"
	"github.com/Korpenter/tusd/v2/pkg/handler"
	"github.com/Korpenter/tusd/v2/pkg/memorylocker"
)

func ExampleNewStoreComposer() {
	composer := handler.NewStoreComposer()

	fs := filestore.New("./data")
	fs.UseIn(composer)

	ml := memorylocker.New()
	ml.UseIn(composer)

	config := handler.Config{
		StoreComposer: composer,
	}

	_, _ = handler.NewHandler(config)
}

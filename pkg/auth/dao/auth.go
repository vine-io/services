package dao

import (
	"github.com/vine-io/services/pkg/runtime/inject"
	"github.com/vine-io/vine/util/runtime"
)

func init() {
	_ = inject.Provide(sets)
}

var sets = runtime.NewSchemaSet()

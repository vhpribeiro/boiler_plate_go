package repository

import (
	"github.com/casbin/casbin/v2/persist"
)

type ICasbinRepository interface {
	GetTheAdapter() (persist.BatchAdapter, error)
}

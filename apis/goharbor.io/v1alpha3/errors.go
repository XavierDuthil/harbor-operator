package v1alpha3

import "errors"

var (
	ErrNoStorageConfiguration = errors.New("no storage configuration")
	Err2StorageConfiguration  = errors.New("only 1 storage can be configured")

	ErrNoMigrationConfiguration = errors.New("no migration source configuration")
	Err2MigrationConfiguration  = errors.New("only 1 migration source can be configured")

	ErrWrongURLFormat = errors.New("wrong url format")
)

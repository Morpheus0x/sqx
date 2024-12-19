package sqx

import (
	"context"
)

type directQueryable struct {
	queryable Queryable
	logger    Logger
}

// DatabaseQueryable returns the required struct used by ReadDirect, WriteDirect and TypedWriteDirect.
// This simplifies any queries to a database not configured as defaultQueryable
// without the need to specify WithQueryable for every request
func DatabaseQueryable(queryable Queryable, logger Logger) directQueryable {
	return directQueryable{queryable, logger}
}

// ReadDerict is the entrypoint for creating generic Select builders for the desired directQueryable
func ReadDirect[T any](ctx context.Context, direct directQueryable) typedRunCtx[T] {
	return typedRunCtx[T]{WriteDirect[T](ctx, direct)}
}

// WriteDirect is the entrypoint for creating sql-extra builders that call ExecCtx
// and its variants - it does not have a generic b/c Exec cannot return arbitrary data.
// The argument "direct" is used to specify the database
func WriteDirect[T any](ctx context.Context, direct directQueryable) runCtx {
	return runCtx{
		ctx:       ctx,
		logger:    direct.logger,
		queryable: direct.queryable,
	}
}

func TypedWriteDirect[T any](ctx context.Context, direct directQueryable) typedRunCtx[T] {
	return typedRunCtx[T]{WriteDirect[T](ctx, direct)}
}

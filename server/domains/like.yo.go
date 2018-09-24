// Package domains contains the types.
package domains

// Code generated by yo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"
)

// Like represents a row from 'Likes'.
type Like struct {
	UUID      string    `spanner:"Uuid" json:"Uuid"`           // Uuid
	SessionID int64     `spanner:"SessionId" json:"SessionId"` // SessionId
	UserUUID  string    `spanner:"UserUUID" json:"UserUUID"`   // UserUUID
	CreatedAt time.Time `spanner:"CreatedAt" json:"CreatedAt"` // CreatedAt
}

func LikePrimaryKeys() []string {
	return []string{
		"Uuid",
	}
}

func LikeColumns() []string {
	return []string{
		"Uuid",
		"SessionId",
		"UserUUID",
		"CreatedAt",
	}
}

func (l *Like) columnsToPtrs(cols []string, customPtrs map[string]interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		if val, ok := customPtrs[col]; ok {
			ret = append(ret, val)
			continue
		}

		switch col {
		case "Uuid":
			ret = append(ret, &l.UUID)
		case "SessionId":
			ret = append(ret, &l.SessionID)
		case "UserUUID":
			ret = append(ret, &l.UserUUID)
		case "CreatedAt":
			ret = append(ret, &l.CreatedAt)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (l *Like) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "Uuid":
			ret = append(ret, l.UUID)
		case "SessionId":
			ret = append(ret, l.SessionID)
		case "UserUUID":
			ret = append(ret, l.UserUUID)
		case "CreatedAt":
			ret = append(ret, l.CreatedAt)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newLike_Decoder returns a decoder which reads a row from *spanner.Row
// into Like. The decoder is not goroutine-safe. Don't use it concurrently.
func newLike_Decoder(cols []string) func(*spanner.Row) (*Like, error) {
	customPtrs := map[string]interface{}{}

	return func(row *spanner.Row) (*Like, error) {
		var l Like
		ptrs, err := l.columnsToPtrs(cols, customPtrs)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &l, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (l *Like) Insert(ctx context.Context) *spanner.Mutation {
	return spanner.Insert("Likes", LikeColumns(), []interface{}{
		l.UUID, l.SessionID, l.UserUUID, l.CreatedAt,
	})
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (l *Like) Update(ctx context.Context) *spanner.Mutation {
	return spanner.Update("Likes", LikeColumns(), []interface{}{
		l.UUID, l.SessionID, l.UserUUID, l.CreatedAt,
	})
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (l *Like) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	return spanner.InsertOrUpdate("Likes", LikeColumns(), []interface{}{
		l.UUID, l.SessionID, l.UserUUID, l.CreatedAt,
	})
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (l *Like) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, LikePrimaryKeys()...)

	values, err := l.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "Like.UpdateColumns", "Likes", err)
	}

	return spanner.Update("Likes", colsWithPKeys, values), nil
}

// FindLike gets a Like by primary key
func FindLike(ctx context.Context, db YORODB, uuid string) (*Like, error) {
	key := spanner.Key{uuid}
	row, err := db.ReadRow(ctx, "Likes", key, LikeColumns())
	if err != nil {
		return nil, newError("FindLike", "Likes", err)
	}

	decoder := newLike_Decoder(LikeColumns())
	l, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindLike", "Likes", err)
	}

	return l, nil
}

// Delete deletes the Like from the database.
func (l *Like) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := l.columnsToValues(LikePrimaryKeys())
	return spanner.Delete("Likes", spanner.Key(values))
}

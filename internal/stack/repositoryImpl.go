package stack

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-backend-training/internal/mserror"
	"time"
)

const (
	createStackQuery       = "INSERT INTO stacks (name, description, segment_id, purpose_id, application_name, created_at, updated_at, created_by, updated_by) VALUES(?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?, ?);"
	selectStackByNameQuery = "SELECT s.id, s.name as stack_name, s.description as stack_desc, s.application_name as app_name, seg.name AS segment_name, seg.is_productive as segment_is_productive, p.name AS purpose_name, p.description as purpose_desc, p.is_productive as purpose_is_productive, s.created_by, s.created_at, s.updated_by, s.updated_at FROM stacks s JOIN segments seg ON s.segment_id = seg.id JOIN purposes p ON s.purpose_id = p.id WHERE s.application_name = ? AND s.name = ?;"
)

func (r *stackRepository) InsertStack(ctx context.Context, s Stack) (Stack, error) {
	resp, err := r.db.ExecContext(ctx, createStackQuery, s.Name, s.Description, s.Segment.Id, s.Purpose.Id, s.ApplicationName, s.CreatedBy, s.UpdatedBy)
	if err != nil {
		return Stack{}, mserror.Wrap(err, fmt.Sprintf("failed to insert stack: %s for application: %s: %s", s.Name, s.ApplicationName, err.Error()))
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return Stack{}, mserror.Wrap(err, fmt.Sprintf("failed to retrieve last insert id for stack: %s for application: %s: %s", s.Name, s.ApplicationName, err.Error()))
	}

	s.Id = int(id)
	return s, nil
}

func (r *stackRepository) SelectStackByName(ctx context.Context, applicationName string, stackName string) (Stack, error) {
	var entity stackEntity
	err := r.db.QueryRowContext(ctx, selectStackByNameQuery, applicationName, stackName).
		Scan(&entity.Id, &entity.Name, &entity.Description, &entity.ApplicationName,
			&entity.SegmentName, &entity.SegIsProductive, &entity.PurposeName, &entity.PurposeDescription, &entity.PurpIsProductive,
			&entity.CreatedBy, &entity.CreatedAt, &entity.UpdatedBy, &entity.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Stack{}, mserror.Wrap(ErrStackNotFound, fmt.Sprintf("stack: %s not found for application: %s: %s", stackName, applicationName, err.Error()))
		}
		return Stack{}, mserror.Wrap(ErrSearchStackByName, fmt.Sprintf("failed to search stack: %s, application: %s: %s", stackName, applicationName, err.Error()))
	}
	return toStackDomain(entity), nil
}

func (r *stackRepository) SelectAllStacksByApplicationName(ctx context.Context, applicationName string) ([]Stack, error) {
	return nil, nil
}

func (r *stackRepository) UpdateStack(ctx context.Context, stack Stack) error {
	return nil
}

func (r *stackRepository) DeleteStackByName(ctx context.Context, applicationName string, stackName string) error {
	return nil
}

type stackEntity struct {
	Id                 int       `db:"id"`
	Name               string    `db:"stack_name"`
	Description        string    `db:"stack_desc"`
	ApplicationName    string    `db:"app_name"`
	SegmentName        string    `db:"segment_name"`
	SegIsProductive    int       `db:"segment_is_productive"`
	PurposeName        string    `db:"purpose_name"`
	PurposeDescription string    `db:"purpose_desc"`
	PurpIsProductive   int       `db:"purpose_is_productive"`
	CreatedBy          string    `db:"created_by"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedBy          string    `db:"updated_by"`
	UpdatedAt          time.Time `db:"updated_at"`
}

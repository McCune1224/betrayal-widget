// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: actions.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAction = `-- name: CreateAction :one
INSERT INTO actions (
  game_id, player_id, pending_approval, resolved, target, context, ability_name, role_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8 
)
RETURNING id, game_id, player_id, pending_approval, resolved, target, context, ability_name, role_id
`

type CreateActionParams struct {
	GameID          pgtype.Int4 `json:"game_id"`
	PlayerID        pgtype.Int4 `json:"player_id"`
	PendingApproval bool        `json:"pending_approval"`
	Resolved        bool        `json:"resolved"`
	Target          string      `json:"target"`
	Context         string      `json:"context"`
	AbilityName     string      `json:"ability_name"`
	RoleID          pgtype.Int4 `json:"role_id"`
}

func (q *Queries) CreateAction(ctx context.Context, arg CreateActionParams) (Action, error) {
	row := q.db.QueryRow(ctx, createAction,
		arg.GameID,
		arg.PlayerID,
		arg.PendingApproval,
		arg.Resolved,
		arg.Target,
		arg.Context,
		arg.AbilityName,
		arg.RoleID,
	)
	var i Action
	err := row.Scan(
		&i.ID,
		&i.GameID,
		&i.PlayerID,
		&i.PendingApproval,
		&i.Resolved,
		&i.Target,
		&i.Context,
		&i.AbilityName,
		&i.RoleID,
	)
	return i, err
}

const deleteAction = `-- name: DeleteAction :exec
delete from actions
where id = $1
`

func (q *Queries) DeleteAction(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteAction, id)
	return err
}

const getAction = `-- name: GetAction :one
select id, game_id, player_id, pending_approval, resolved, target, context, ability_name, role_id
from actions
where id = $1
`

func (q *Queries) GetAction(ctx context.Context, id int32) (Action, error) {
	row := q.db.QueryRow(ctx, getAction, id)
	var i Action
	err := row.Scan(
		&i.ID,
		&i.GameID,
		&i.PlayerID,
		&i.PendingApproval,
		&i.Resolved,
		&i.Target,
		&i.Context,
		&i.AbilityName,
		&i.RoleID,
	)
	return i, err
}

const getActionByID = `-- name: GetActionByID :one
select id, game_id, player_id, pending_approval, resolved, target, context, ability_name, role_id
from actions
where id = $1
`

func (q *Queries) GetActionByID(ctx context.Context, id int32) (Action, error) {
	row := q.db.QueryRow(ctx, getActionByID, id)
	var i Action
	err := row.Scan(
		&i.ID,
		&i.GameID,
		&i.PlayerID,
		&i.PendingApproval,
		&i.Resolved,
		&i.Target,
		&i.Context,
		&i.AbilityName,
		&i.RoleID,
	)
	return i, err
}

const listActions = `-- name: ListActions :many
select id, game_id, player_id, pending_approval, resolved, target, context, ability_name, role_id
from actions
`

func (q *Queries) ListActions(ctx context.Context) ([]Action, error) {
	rows, err := q.db.Query(ctx, listActions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Action
	for rows.Next() {
		var i Action
		if err := rows.Scan(
			&i.ID,
			&i.GameID,
			&i.PlayerID,
			&i.PendingApproval,
			&i.Resolved,
			&i.Target,
			&i.Context,
			&i.AbilityName,
			&i.RoleID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listActionsByGame = `-- name: ListActionsByGame :many
select id, game_id, player_id, pending_approval, resolved, target, context, ability_name, role_id
from actions
where game_id = $1
`

func (q *Queries) ListActionsByGame(ctx context.Context, gameID pgtype.Int4) ([]Action, error) {
	rows, err := q.db.Query(ctx, listActionsByGame, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Action
	for rows.Next() {
		var i Action
		if err := rows.Scan(
			&i.ID,
			&i.GameID,
			&i.PlayerID,
			&i.PendingApproval,
			&i.Resolved,
			&i.Target,
			&i.Context,
			&i.AbilityName,
			&i.RoleID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listActionsByPlayer = `-- name: ListActionsByPlayer :many
SELECT a.id, a.game_id, a.player_id, a.pending_approval, a.resolved, a.target, a.context, a.ability_name, a.role_id
FROM actions a
JOIN players p on p.id = a.player_id
WHERE p.id = $1
`

func (q *Queries) ListActionsByPlayer(ctx context.Context, id int32) ([]Action, error) {
	rows, err := q.db.Query(ctx, listActionsByPlayer, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Action
	for rows.Next() {
		var i Action
		if err := rows.Scan(
			&i.ID,
			&i.GameID,
			&i.PlayerID,
			&i.PendingApproval,
			&i.Resolved,
			&i.Target,
			&i.Context,
			&i.AbilityName,
			&i.RoleID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listActionsByRoundForGame = `-- name: ListActionsByRoundForGame :many
SELECT a.id, a.game_id, a.player_id, a.pending_approval, a.resolved, a.target, a.context, a.ability_name, a.role_id
FROM actions a
JOIN games g on $1 = a.game_id
WHERE g.round = $2
`

type ListActionsByRoundForGameParams struct {
	GameID pgtype.Int4 `json:"game_id"`
	Round  int32       `json:"round"`
}

func (q *Queries) ListActionsByRoundForGame(ctx context.Context, arg ListActionsByRoundForGameParams) ([]Action, error) {
	rows, err := q.db.Query(ctx, listActionsByRoundForGame, arg.GameID, arg.Round)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Action
	for rows.Next() {
		var i Action
		if err := rows.Scan(
			&i.ID,
			&i.GameID,
			&i.PlayerID,
			&i.PendingApproval,
			&i.Resolved,
			&i.Target,
			&i.Context,
			&i.AbilityName,
			&i.RoleID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAction = `-- name: UpdateAction :one
UPDATE actions
  SET game_id = $2,
  player_id = $3,
  pending_approval = $4,
  resolved = $5,
  target = $6,
  context = $7,
  ability_name = $8,
  role_id = $9
WHERE id = $1
RETURNING id, game_id, player_id, pending_approval, resolved, target, context, ability_name, role_id
`

type UpdateActionParams struct {
	ID              int32       `json:"id"`
	GameID          pgtype.Int4 `json:"game_id"`
	PlayerID        pgtype.Int4 `json:"player_id"`
	PendingApproval bool        `json:"pending_approval"`
	Resolved        bool        `json:"resolved"`
	Target          string      `json:"target"`
	Context         string      `json:"context"`
	AbilityName     string      `json:"ability_name"`
	RoleID          pgtype.Int4 `json:"role_id"`
}

func (q *Queries) UpdateAction(ctx context.Context, arg UpdateActionParams) (Action, error) {
	row := q.db.QueryRow(ctx, updateAction,
		arg.ID,
		arg.GameID,
		arg.PlayerID,
		arg.PendingApproval,
		arg.Resolved,
		arg.Target,
		arg.Context,
		arg.AbilityName,
		arg.RoleID,
	)
	var i Action
	err := row.Scan(
		&i.ID,
		&i.GameID,
		&i.PlayerID,
		&i.PendingApproval,
		&i.Resolved,
		&i.Target,
		&i.Context,
		&i.AbilityName,
		&i.RoleID,
	)
	return i, err
}

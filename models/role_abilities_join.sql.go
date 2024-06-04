// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: role_abilities_join.sql

package models

import (
	"context"
)

const createRoleAbilityJoin = `-- name: CreateRoleAbilityJoin :one
INSERT INTO role_abilities_join (
  role_id, ability_id
) VALUES (
  $1, $2
) RETURNING role_id, ability_id
`

type CreateRoleAbilityJoinParams struct {
	RoleID    int32 `json:"role_id"`
	AbilityID int32 `json:"ability_id"`
}

func (q *Queries) CreateRoleAbilityJoin(ctx context.Context, arg CreateRoleAbilityJoinParams) (RoleAbilitiesJoin, error) {
	row := q.db.QueryRow(ctx, createRoleAbilityJoin, arg.RoleID, arg.AbilityID)
	var i RoleAbilitiesJoin
	err := row.Scan(&i.RoleID, &i.AbilityID)
	return i, err
}

const getAssociatedRoleAbilities = `-- name: GetAssociatedRoleAbilities :many
SELECT ab.id, ab.name, ab.description, ab.default_charges, ab.category_ids, ab.rarity, ab.priority, ab.any_ability
FROM role_abilities_join raj
JOIN ability_details ab ON raj.ability_id = ab.id
WHERE raj.role_id = $1
`

func (q *Queries) GetAssociatedRoleAbilities(ctx context.Context, roleID int32) ([]AbilityDetail, error) {
	rows, err := q.db.Query(ctx, getAssociatedRoleAbilities, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AbilityDetail
	for rows.Next() {
		var i AbilityDetail
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.DefaultCharges,
			&i.CategoryIds,
			&i.Rarity,
			&i.Priority,
			&i.AnyAbility,
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

const getRoleAbilityJoin = `-- name: GetRoleAbilityJoin :one
SELECT role_abilities_join.role_id, role_abilities_join.ability_id, abilities.id, abilities.ability_details_id, abilities.player_inventory_id
FROM role_abilities_join
JOIN abilities ON role_abilities_join.ability_id = abilities.id
`

type GetRoleAbilityJoinRow struct {
	RoleAbilitiesJoin RoleAbilitiesJoin `json:"role_abilities_join"`
	Ability           Ability           `json:"ability"`
}

func (q *Queries) GetRoleAbilityJoin(ctx context.Context) (GetRoleAbilityJoinRow, error) {
	row := q.db.QueryRow(ctx, getRoleAbilityJoin)
	var i GetRoleAbilityJoinRow
	err := row.Scan(
		&i.RoleAbilitiesJoin.RoleID,
		&i.RoleAbilitiesJoin.AbilityID,
		&i.Ability.ID,
		&i.Ability.AbilityDetailsID,
		&i.Ability.PlayerInventoryID,
	)
	return i, err
}
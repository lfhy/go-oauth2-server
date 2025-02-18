package oauth

import (
	"errors"

	"go-oauth2-server/models"
)

var (
	// ErrRoleNotFound ...
	ErrRoleNotFound = errors.New("Role not found")
)

// FindRoleByID looks up a role by ID and returns it
func (s *Service) FindRoleByID(id string) (*models.OauthRole, error) {
	role := new(models.OauthRole)
	if s.db.Where("id = ?", id).First(role) == nil {
		return nil, ErrRoleNotFound
	}
	return role, nil
}

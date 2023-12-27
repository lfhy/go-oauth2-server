package models

import (
	"database/sql"
	"time"

	"go-oauth2-server/util"

	"github.com/RichardKnop/uuid"
	"gorm.io/gorm"
)

// OauthClient ...
type OauthClient struct {
	MyGormModel
	Key         string         `gorm:"column:Key;unique;not null"`
	Secret      string         `gorm:"column:Secret;not null"`
	RedirectURI sql.NullString `gorm:"column:RedirectURI"`
}

// TableName specifies table name
func (c *OauthClient) TableName() string {
	return "oauth_clients"
}

// OauthScope ...
type OauthScope struct {
	MyGormModel
	Scope       string         `gorm:"column:Scope;unique;not null"`
	Description sql.NullString `gorm:"column:Description"`
	IsDefault   bool           `gorm:"column:IsDefault,default:false"`
}

// TableName specifies table name
func (s *OauthScope) TableName() string {
	return "oauth_scopes"
}

// OauthRole is a one of roles user can have (currently superuser or user)
type OauthRole struct {
	TimestampModel
	ID   string `gorm:"column:ID;primary_key,"`
	Name string `gorm:"column:Name;unique;not null"`
}

// TableName specifies table name
func (r *OauthRole) TableName() string {
	return "oauth_roles"
}

// OauthUser ...
type OauthUser struct {
	MyGormModel
	RoleID   sql.NullString `gorm:"column:RoleID;index;not null"`
	Role     *OauthRole
	Username string         `gorm:"column:Username;unique;not null"`
	Password sql.NullString `gorm:"column:Password"`
}

// TableName specifies table name
func (u *OauthUser) TableName() string {
	return "oauth_users"
}

// OauthRefreshToken ...
type OauthRefreshToken struct {
	MyGormModel
	ClientID  sql.NullString `gorm:"column:ClientID;index;not null"`
	UserID    sql.NullString `gorm:"column:UserID;index"`
	Client    *OauthClient
	User      *OauthUser
	Token     string    `gorm:"column:Token;unique;not null"`
	ExpiresAt time.Time `gorm:"column:ExpiresAt;not null"`
	Scope     string    `gorm:"column:Scope;not null"`
}

// TableName specifies table name
func (rt *OauthRefreshToken) TableName() string {
	return "oauth_refresh_tokens"
}

// OauthAccessToken ...
type OauthAccessToken struct {
	MyGormModel
	ClientID  sql.NullString `gorm:"column:ClientID;index;not null"`
	UserID    sql.NullString `gorm:"column:UserID;index"`
	Client    *OauthClient
	User      *OauthUser
	Token     string    `gorm:"column:Token;unique;not null"`
	ExpiresAt time.Time `gorm:"column:ExpiresAt;not null"`
	Scope     string    `gorm:"column:Scope;not null"`
}

// TableName specifies table name
func (at *OauthAccessToken) TableName() string {
	return "oauth_access_tokens"
}

// OauthAuthorizationCode ...
type OauthAuthorizationCode struct {
	MyGormModel
	ClientID    sql.NullString `gorm:"column:ClientID;index;not null"`
	UserID      sql.NullString `gorm:"column:UserID;index;not null"`
	Client      *OauthClient
	User        *OauthUser
	Code        string         `gorm:"column:Code;unique;not null"`
	RedirectURI sql.NullString `gorm:"column:RedirectURI"`
	ExpiresAt   time.Time      `gorm:"column:ExpiresAt;not null"`
	Scope       string         `gorm:"column:Scope;not null"`
}

// TableName specifies table name
func (ac *OauthAuthorizationCode) TableName() string {
	return "oauth_authorization_codes"
}

// NewOauthRefreshToken creates new OauthRefreshToken instance
func NewOauthRefreshToken(client *OauthClient, user *OauthUser, expiresIn int, scope string) *OauthRefreshToken {
	refreshToken := &OauthRefreshToken{
		MyGormModel: MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		ClientID:  util.StringOrNull(string(client.ID)),
		Token:     uuid.New(),
		ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		Scope:     scope,
	}
	if user != nil {
		refreshToken.UserID = util.StringOrNull(string(user.ID))
	}
	return refreshToken
}

// NewOauthAccessToken creates new OauthAccessToken instance
func NewOauthAccessToken(client *OauthClient, user *OauthUser, expiresIn int, scope string) *OauthAccessToken {
	accessToken := &OauthAccessToken{
		MyGormModel: MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		ClientID:  util.StringOrNull(string(client.ID)),
		Token:     uuid.New(),
		ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		Scope:     scope,
	}
	if user != nil {
		accessToken.UserID = util.StringOrNull(string(user.ID))
	}
	return accessToken
}

// NewOauthAuthorizationCode creates new OauthAuthorizationCode instance
func NewOauthAuthorizationCode(client *OauthClient, user *OauthUser, expiresIn int, redirectURI, scope string) *OauthAuthorizationCode {
	return &OauthAuthorizationCode{
		MyGormModel: MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		ClientID:    util.StringOrNull(string(client.ID)),
		UserID:      util.StringOrNull(string(user.ID)),
		Code:        uuid.New(),
		ExpiresAt:   time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		RedirectURI: util.StringOrNull(redirectURI),
		Scope:       scope,
	}
}

// OauthAuthorizationCodePreload sets up Gorm preloads for an auth code object
func OauthAuthorizationCodePreload(db *gorm.DB) *gorm.DB {
	return OauthAuthorizationCodePreloadWithPrefix(db, "")
}

// OauthAuthorizationCodePreloadWithPrefix sets up Gorm preloads for an auth code object,
// and prefixes with prefix for nested objects
func OauthAuthorizationCodePreloadWithPrefix(db *gorm.DB, prefix string) *gorm.DB {
	return db.
		Preload(prefix + "Client").Preload(prefix + "User")
}

// OauthAccessTokenPreload sets up Gorm preloads for an access token object
func OauthAccessTokenPreload(db *gorm.DB) *gorm.DB {
	return OauthAccessTokenPreloadWithPrefix(db, "")
}

// OauthAccessTokenPreloadWithPrefix sets up Gorm preloads for an access token object,
// and prefixes with prefix for nested objects
func OauthAccessTokenPreloadWithPrefix(db *gorm.DB, prefix string) *gorm.DB {
	return db.
		Preload(prefix + "Client").Preload(prefix + "User")
}

// OauthRefreshTokenPreload sets up Gorm preloads for a refresh token object
func OauthRefreshTokenPreload(db *gorm.DB) *gorm.DB {
	return OauthRefreshTokenPreloadWithPrefix(db, "")
}

// OauthRefreshTokenPreloadWithPrefix sets up Gorm preloads for a refresh token object,
// and prefixes with prefix for nested objects
func OauthRefreshTokenPreloadWithPrefix(db *gorm.DB, prefix string) *gorm.DB {
	return db.
		Preload(prefix + "Client").Preload(prefix + "User")
}

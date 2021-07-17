package store

import "github.com/saturninoabril/dashboard-server/model"

type Store interface {
	Role() RoleStore
	Session() SessionStore
	Token() TokenStore
	User() UserStore
}

type RoleStore interface {
	CreateRole(role *model.Role) (*model.Role, error)
	GetRoleByName(name string) (*model.Role, error)
	UserHasRole(userID, roleID string) (bool, error)
	UserHasRoleByName(userID, roleName string) (bool, error)
	AddUserRole(userID, roleID string) error
	DeleteUserRole(userID, roleID string) error
}

type SessionStore interface {
	CreateSession(session *model.Session) (*model.Session, error)
	GetSession(idOrToken string) (*model.Session, error)
	DeleteSession(id string) error
	DeleteSessionsForUser(userID string) error
}

type TokenStore interface {
	CreateToken(token *model.Token) (*model.Token, error)
	GetToken(tokenValue string) (*model.Token, error)
	GetTokensByEmail(email, tokenType string) ([]*model.Token, error)
	DeleteToken(tokenValue string) error
	DeleteTokensByEmail(email, tokenType string) error
	CleanupTokenStore(expiryTimeMillis int64)
}

type UserStore interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUser(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	VerifyEmail(id, email string) error
	UnverifyEmail(id, email string) error
	UpdatePassword(id, password string) error
	UpdateUser(user *model.User) error
	UpdateUserState(userID string, state string) error
}

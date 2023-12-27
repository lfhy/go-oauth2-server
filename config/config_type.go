package config

// DatabaseConfig stores database connection options
type DatabaseConfig struct {
	Type         string `json:"Type"`
	Host         string `json:"Host"`
	Port         int    `json:"Port"`
	User         string `json:"User"`
	Password     string `json:"Password"`
	DatabaseName string `json:"DatabaseName"`
	MaxIdleConns int    `json:"MaxIdleConns"`
	MaxOpenConns int    `json:"MaxOpenConns"`
}

// OauthConfig stores oauth service configuration options
type OauthConfig struct {
	AccessTokenLifetime  int `json:"AccessTokenLifetime"`
	RefreshTokenLifetime int `json:"RefreshTokenLifetime"`
	AuthCodeLifetime     int `json:"AuthCodeLifetime"`
}

// SessionConfig stores session configuration for the web app
type SessionConfig struct {
	Secret string `json:"Secret"`
	Path   string `json:"Path"`
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge int `json:"MaxAge"`
	// When you tag a cookie with the HttpOnly flag, it tells the browser that
	// this particular cookie should only be accessed by the server.
	// Any attempt to access the cookie from client script is strictly forbidden.
	HTTPOnly bool `json:"HTTPOnly"`
}

// Config stores all configuration options
type Config struct {
	Database      DatabaseConfig `json:"Database"`
	Oauth         OauthConfig    `json:"Oauth"`
	Session       SessionConfig  `json:"Session"`
	IsDevelopment bool           `json:"IsDevelopment"`
}

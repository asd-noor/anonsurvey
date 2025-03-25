package outbound

import "anonsurvey/config"

type DBClient interface {
	ConnectDB(config.DBConfig) error
}

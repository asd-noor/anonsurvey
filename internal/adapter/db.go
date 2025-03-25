package adapter

type DBAdapter struct {
	DSN       string
	IPAddress string
}

func (db *DBAdapter) Save() {
	panic("not implemented") // TODO: Implement
}

func NewDBAdapter() DBAdapter {
	return DBAdapter{}
}

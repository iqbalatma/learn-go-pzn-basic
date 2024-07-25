package database

var connection string

func init() {
	connection = "Mysql"
}

func GetConnection() string {
	return connection
}

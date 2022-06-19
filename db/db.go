package db

import "github.com/yoktobit/gogeneral/general/dataaccess"

func Connect() dataaccess.Connection {
	return dataaccess.NewConnectionWithEnvironment()
}

func CheckMigration(connection dataaccess.Connection) dataaccess.Connection {

	return connection
}

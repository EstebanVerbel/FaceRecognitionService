package dbhelper

import (
	"github.com/Dev/FaceRecognitionService/FaceValidationService/models/constants"
	"github.com/gocql/gocql"
)

// GetSession Gets session from Cassandra's cluster
func GetSession() *gocql.Session {

	cluster := gocql.NewCluster(constants.ClusterAddress)

	cluster.Keyspace = constants.KeyspaceName
	cluster.Consistency = gocql.One

	session, _ := cluster.CreateSession()

	return session
}

package dbhelper

import "github.com/gocql/gocql"

// GetSession Gets session from Cassandra's cluster
func GetSession() *gocql.Session {

	cluster := gocql.NewCluster("127.0.0.1")

	cluster.Keyspace = "faces"
	cluster.Consistency = gocql.Quorum

	session, _ := cluster.CreateSession()
	defer session.Close()

	return session
}

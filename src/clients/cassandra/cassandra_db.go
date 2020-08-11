package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("192.168.0.44")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}

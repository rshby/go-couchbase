package database

import (
	"github.com/couchbase/gocb/v2"
	"github.com/sirupsen/logrus"
	"go-couchbase/config"
	"time"
)

var DatabaseCouchbase *gocb.Bucket

// InitializeDatabaseCouchbase is function to connect database couchbase
func InitializeDatabaseCouchbase() (*gocb.Bucket, func()) {
	options := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: config.CouchbaseUser(),
			Password: config.CouchbasePassword(),
		},
	}

	if err := options.ApplyProfile(gocb.ClusterConfigProfileWanDevelopment); err != nil {
		logrus.Fatal(err)
	}

	// connect
	cluster, err := gocb.Connect(config.CouchbaseHost(), options)
	if err != nil {
		logrus.Fatal(err)
	}

	bucket := cluster.Bucket(config.CouchbaseBucket())
	if err = bucket.WaitUntilReady(5*time.Second, nil); err != nil {
		logrus.Fatal(err)
	}

	DatabaseCouchbase = bucket

	logrus.Infof("success connected to database COUCHBASE [%s]", config.CouchbaseHost())
	return DatabaseCouchbase, func() {
		cluster.Close(nil)
	}
}

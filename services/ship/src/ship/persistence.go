package ship

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var instance *mgo.Session = nil

func isConnectionHealthy() bool {
	return instance != nil && instance.Ping() == nil
}

func GetPersistence() *mgo.Session {
	if !isConnectionHealthy() {
		for {
			connect, err := connect()

			if err != nil {
				log.Printf("[Ship][Persistence] Error connecting to database, waiting for retry... %v", err)
				time.Sleep(2 * time.Second)
			} else {
				instance = connect
				break
			}
		}
	}

	return instance
}

func getDatabase() *mgo.Database {
	return GetPersistence().DB(getConfig().Mongo.DB)
}

func getCollection() *mgo.Collection {
	return getDatabase().C(getConfig().Mongo.DB)
}

func connect() (*mgo.Session, error) {
	conf := getConfig()

	log.Print("[Ship][Persistence] No valid connection to database, attempting to connect...")
	dialInfo := &mgo.DialInfo{
		Username: conf.Mongo.Username,
		Password: conf.Mongo.Password,
		Source: "admin",
		Addrs: []string{fmt.Sprintf("%s:%v", conf.Mongo.Domain, conf.Mongo.Port)},
		Timeout: 60 * time.Second,
	}

	session, err := mgo.DialWithInfo(dialInfo); if err != nil {
		return nil, err
	}

	return session, nil
}

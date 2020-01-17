package persistence

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

type Mongo struct {
	session			*mgo.Session

	MongoDomain     string
	MongoPort       int
	MongoDB         string
	MongoCollection string
	MongoUsername   string
	MongoPassword   string
}

func (m *Mongo) IsConnectionHealthy() bool {
	return m.session != nil && m.session.Ping() == nil
}

func (m *Mongo) GetSession() *mgo.Session {
	if !m.IsConnectionHealthy() {
		for {
			connect, err := m.connect()

			if err != nil {
				log.Printf("[Persistence] Error connecting to database, waiting for retry... %v", err)
				time.Sleep(2 * time.Second)
			} else {
				m.session = connect
				break
			}
		}
	}

	return m.session
}

func (m *Mongo) GetDatabase() *mgo.Database {
	return m.GetSession().DB(m.MongoDB)
}

func (m *Mongo) GetCollection() *mgo.Collection {
	return m.GetDatabase().C(m.MongoCollection)
}

func (m *Mongo) connect() (*mgo.Session, error) {
	log.Print("[Persistence] No valid connection to database, attempting to connect...")
	dialInfo := &mgo.DialInfo{
		Username: m.MongoUsername,
		Password: m.MongoPassword,
		Source: "admin",
		Addrs: []string{fmt.Sprintf("%s:%v", m.MongoDomain, m.MongoPort)},
		Timeout: 60 * time.Second,
	}

	session, err := mgo.DialWithInfo(dialInfo); if err != nil {
		return nil, err
	}

	return session, nil
}

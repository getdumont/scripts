package utilities

import (
	"log"
	"os"
	"net"
	"strings"
	"crypto/tls"
	"github.com/globalsign/mgo"
)

type MongoConfig struct {
	Db 		   string
	Addrs 	   string
	Username   string
	Password   string
	AuthExists bool
}

var LocalConfig = MongoConfig{
	Addrs: "127.0.0.1",
	Db: "dumont_local",
	Username: "",
	Password: "",
	AuthExists: false,
}

var AtlasConfig = MongoConfig{
	Addrs: os.Getenv("MONGO_ADDRS"),
	Db: os.Getenv("MONGO_DB"),
	Username: os.Getenv("MONGO_USER"),
	Password: os.Getenv("MONGO_PASSWORD"),
	AuthExists: true,
}

func ConnectAndGetCollection(config MongoConfig, collection string) (*mgo.Collection, func()) {
	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs: strings.Split(config.Addrs, ","),
	}

	if config.AuthExists {
		dialInfo.Username = config.Username
		dialInfo.Password = config.Password

		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Fatal(err)
	}

	db := session.DB(config.Db)

	return db.C(collection), session.Close
}
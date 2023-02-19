package conf_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/renmcc/go-cmdb/conf"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/config.toml")
	if should.NoError(err) {
		should.Equal("0.0.0.0", conf.C().App.Host)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)

	os.Setenv("D_MYSQL_DATABASE", "demo_db")
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		fmt.Println(conf.C().MySQL.Database)
	}
}

func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/config.toml")
	if should.NoError(err) {
		conf.C().MySQL.GetDB()
	}
}

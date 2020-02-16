package dao

import (
	"context"
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jinzhu/gorm"
	"github.com/starichat/notes/DesignPatter/credit_system/config"
	"log"
)

type Dao struct {
	c *config.DBConfig
	DB *gorm.DB
}

// New init mysql db
func New(c *config.DBConfig) (d *Dao) {
	d = &Dao{
		c : c,
	}
	d.initDB()
	return d
}

func(d *Dao) initDB() (err error) {
	log.Println(d.c.URL)


	d.DB, err = gorm.Open(d.c.Connection, d.c.URL)
	if err != nil {
		log.Fatal("Database connection failed. Database url : "+d.c.URL+" error: ",err)
	}
	log.Println("gorm!!!")
	return err
}

func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}

func (d *Dao) Ping(c context.Context) (err error) {
	if d.DB != nil {
		err = d.DB.DB().PingContext(c)
	}
	return err
}
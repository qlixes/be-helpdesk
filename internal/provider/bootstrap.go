package provider

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
)

var bootstrap *Bootstrap

var mutex sync.Mutex
var once sync.Once

type Bootstrap struct {
	Config *Config
	Database *Database
}

type BootstrapMethod interface {
	UsePgsql()
}

func NewBootstrap() *Bootstrap {

	once.Do(func() {
		bootstrap = &Bootstrap{
			Config: NewConfig(),
		}
	})

	return bootstrap
}

func (b *Bootstrap) UsePgsql() {

	mutex.Lock()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Jakarta",
		b.Config.Db.User,
		b.Config.Db.Pass,
		b.Config.Db.Host,
		b.Config.Db.Port,
		b.Config.Db.Name,
	)

	pgx, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatalf("Unable connect database : %s \n", err.Error())
	}

	connector := postgres.New(postgres.Config{
		Conn: pgx,
	})

	b.Database = NewDatabase(connector)
}

package main

type Config interface {
	Address() string
	Port() int
}

type config struct {
	address string
	port    int
}

func (c *config) Address() string {
	return c.address
}

func (c *config) Port() int {
	return c.port
}

func NewConfig() Config {
	return &config{
		address: "localhost",
		port:    8080,
	}
}

type Database interface {
	NameDB() string
	User() string
}

type db struct {
	name string
	user string
}

func (d *db) NameDB() string {
	return d.name
}

func (d *db) User() string {
	return d.user
}

func NewDB() Database {
	return &db{
		name: "postgres",
		user: "aleksei",
	}
}

type Server struct {
	Db  Database
	Cfg Config
}

func NewServer(db Database, cfg Config) Server {
	return Server{
		Db:  db,
		Cfg: cfg,
	}
}
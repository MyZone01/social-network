package config

import "backend/app/session"

var (
	conf = session.Config{}
	Sess = session.New(&conf)
)
package config

import "testing"

func TestInit(t *testing.T) {
	Init("../../conf/deploy.local.yml")
}

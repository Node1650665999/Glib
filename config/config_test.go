package config_test

import (
	"github.com/Node1650665999/Glib/config"
	"testing"
)

func TestParseYaml(t *testing.T) {
	filename := "../data/config.yaml"
	cfg,err := config.ParseConfig(filename)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cfg.GetStringMap("mysql"))
	t.Log(cfg.GetString("mysql.host"))
	t.Log(cfg.GetStringSlice("nginx.nginx_list"))
	t.Log(cfg.GetString("nginx.nginx_list.0"))
	t.Log(cfg.GetInt("nginx.counter"))
}

func TestParseIni(t *testing.T) {
	filename := "../data/config.ini"
	cfg,err := config.ParseConfig2(filename)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cfg.StringMap("mysql"))
	t.Log(cfg.String("mysql.host"))
	t.Log(cfg.Strings("nginx.nginx_list"))
	t.Log(cfg.String("nginx.nginx_list.0"))
	t.Log(cfg.Int("nginx.counter"))
}

func TestParseJson(t *testing.T) {
	filename := "../data/config.json"
	cfg,err := config.ParseConfig(filename)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cfg.GetStringMap("mysql"))
	t.Log(cfg.GetString("mysql.host"))
	t.Log(cfg.GetStringSlice("nginx.nginx_list"))
	t.Log(cfg.GetString("nginx.nginx_list.2"))
	t.Log(cfg.GetInt("nginx.counter"))
}

func TestParseEnv(t *testing.T) {
	filename := "../data/.env"

	cfg,err := config.ParseConfig(filename)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cfg.GetString("MYSQL_HOST"))
}
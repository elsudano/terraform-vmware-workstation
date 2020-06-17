package vmworkstation

import (
	"fmt"
	"log"

	wsapiclient "github.com/elsudano/vmware-workstation-api-client/wsapiclient"
	"github.com/hashicorp/terraform/helper/schema"
)

type Config struct {
	User         string
	Password     string
	URL          string
	InsecureFlag bool
	Debug        bool
}

func NewConfig(d *schema.ResourceData) (*Config, error) {
	if d.Get("user").(string) == "" || d.Get("password").(string) == "" || d.Get("url").(string) == "" {
		err := fmt.Errorf("User, Password and URL that required parameters")
		return nil, err
	}
	config := &Config{
		User:         d.Get("user").(string),
		Password:     d.Get("password").(string),
		URL:          d.Get("url").(string),
		InsecureFlag: d.Get("https").(bool),
		Debug:        d.Get("debug").(bool),
	}
	log.Printf("[VMWS] Fi: config.go Fu: NewConfig Ob: %#v\n", config)
	return config, nil
}

func (co *Config) Client() (*wsapiclient.Client, error) {
	client, err := wsapiclient.New()
	client.ConfigCli(co.User, co.Password, co.URL, co.Debug)
	if err != nil {
		return nil, err
	}
	log.Printf("[VMWS] Fi: config.go Fu: Client Ob: %#v\n", client)
	// client.SwitchDebug()
	return client, err
}

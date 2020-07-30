package provider

import (

	//"context"

	// "errors"\
	"context"
	"log"
	"time"

	"github.com/bizflycloud/gobizfly"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BIZFLY_API_URL", "https://manage.bizflycloud.vn"),
				Description: "The URL to use for the BIZFLY API.",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("USERNAME", nil),
				Description: "username for the BIZFLY API`.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PASSWORD", nil),
				Description: "password for the BIZFLY API`.",
			},
		},
		ConfigureFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"bizfly_instance": resouceInstance(),
		},
	}
}
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	host := d.Get("host").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	// client, err := gobizfly.NewClient(
	// 	gobizfly.WithAPIUrl(host),
	// 	gobizfly.WithTenantName(username),
	// )
	client, err := gobizfly.NewClient(
		gobizfly.WithAPIUrl(host),
		gobizfly.WithTenantName(username),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	token, err := client.Token.Create(ctx, &gobizfly.TokenCreateRequest{AuthMethod: "password", Username: username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	client.SetKeystoneToken(token.KeystoneToken)
	return client, nil
}

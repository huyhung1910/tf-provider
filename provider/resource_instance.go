package provider

import (
	"fmt"
		"log"
	"context"
	"time"
	"github.com/bizflycloud/gobizfly"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resouceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resouceInstanceCreate,
		// List:   resouceInstanceList,
		Delete: resourceInstanceDelete,
		Read:   resourceInstanceRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "Name Instance`.",
			},
			"flavorname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "EXP 2c_2g.",
			},
			"sshkey": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "public key`.",
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  true,
				ForceNew: true,

				Description: "Password login to server`.",
			},
			"rootdisk": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"size": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     true,
							ForceNew:    true,
							Description: "Size of RootDisk`.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     true,
							ForceNew:    true,
							Description: "Type of RootDisk`.",
						},
					},
				},
			},
			"datadisk": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"size": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     true,
							Description: "Size of DataDisk`.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     true,
							Description: "Type of DataDisk`.",
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "Type server`.",
			},
			"availabilityzone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "zone HN or HCM`.",
			},
			"os": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     true,
							Description: "Size of RootDisk`.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     true,
							Description: "Type of RootDisk`.",
						},
					},
				},
			},
		},
	}
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resouceInstanceCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*gobizfly.Client)
	fmt.Println(apiClient)
	rootdisk := d.Get("rootdisk").(*schema.Set).List()
	fmt.Println(rootdisk)
	rdisk := make([]string, len(rootdisk))
	log.Println(rdisk)
	for _, raws := range rootdisk {
		fmt.Println(raws)
		// if i == size{
		// 	rdisk[i] = raws.(int)
		// }
		log.Print("--------------------------")
		log.Printf("%T", raws)
		// rdisk[i] = raws.(string)
	}
	// datadisk := d.Get("datadisk").([]interface{})
	// ddisk := make([]string, len(datadisk))
	// for i, raw := range datadisk {
	// 	ddisk[i] = raw.(string)
	// }
	// os := d.Get("example").([]interface{})
	// oss := make([]string, len(os))
	// for i, rawos := range os {
	// 	oss[i] = rawos.(string)
	// }
	// a :=  rdisk.
	item := gobizfly.ServerCreateRequest{
		Name:       d.Get("name").(string),
		FlavorName: d.Get("flavorName").(string),
		SSHKey:     d.Get("sshkey").(string),
		Password:   d.Get("password").(bool),
		//RootDisk:  rdisk ,
	// 	//datadisks : d.Get("datadisks").(string),
		Type:             d.Get("type").(string),
		AvailabilityZone: d.Get("availabilityzone").(string),
	// 	//os : d.Get("os").(string),
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	_,err := apiClient.Server.Create(ctx, &item)
	if err != nil {
		return err
	}
	d.SetId(item.Name)
	return nil

}
func resourceInstanceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceInstanceUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceInstanceRead(d, m)
}

func resourceInstanceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

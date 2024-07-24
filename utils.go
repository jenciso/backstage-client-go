package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/tdabasinskas/go-backstage/v2/backstage"
)

func PrintDomains(c *backstage.Client) {

	if domains, _, err := c.Catalog.Entities.List(context.Background(), &backstage.ListEntityOptions{
		Filters: []string{"kind=Domain"},
	}); err != nil {
		log.Fatal(err)
	} else {
		for _, d := range domains {
			fmt.Printf("%s\n", d.Metadata.Name)
		}
	}
}

func PrintAppListWithSystemsAndDomains(c *backstage.Client) {
	if apps, _, err := c.Catalog.Entities.List(context.Background(), &backstage.ListEntityOptions{
		Filters: []string{"kind=Component"},
	}); err != nil {
		log.Fatal(err)
	} else {
		for _, a := range apps {
			if app, _, err := c.Catalog.Components.Get(context.Background(), a.Metadata.Name, ""); err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%s,%s,%s,%s\n",
					app.Metadata.Name,
					strings.ReplaceAll(app.Spec.Owner, "group:", ""),
					strings.ReplaceAll(app.Spec.System, "TBD", ""),
					GetDomain(c, app.Spec.System))
			}
		}
	}
}

func GetDomain(c *backstage.Client, system string) string {
	defer func() {
		recover()
	}()
	var domain string
	if sys, _, err := c.Catalog.Systems.Get(context.Background(), system, ""); err != nil {
		fmt.Println(err)
	} else {
		domain = sys.Spec.Domain
	}
	return domain
}

func PrintDomainsAndSystems(c *backstage.Client) {
	if domains, _, err := c.Catalog.Entities.List(context.Background(), &backstage.ListEntityOptions{
		Filters: []string{"kind=Domain"},
	}); err != nil {
		log.Fatal(err)
	} else {
		// log.Printf("Component entities: %v", entities)
		for _, d := range domains {
			domain := d.Metadata.Name
			filter := `Kind=System,spec.domain=` + domain
			if systems, _, err := c.Catalog.Entities.List(context.Background(), &backstage.ListEntityOptions{
				Filters: []string{filter},
			}); err != nil {
				log.Fatal(err)
			} else {
				for _, s := range systems {
					system := s.Metadata.Name
					fmt.Printf("%s,%s\n", domain, system)

				}
			}

		}

	}

}

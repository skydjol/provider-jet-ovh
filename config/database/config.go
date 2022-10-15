package database

import (
	"github.com/crossplane/terrajet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_database", func(r *config.Resource) {
		r.UseAsync = true

		r.ShortGroup = "database.cloud"
		r.Kind = "Database"
	})
}

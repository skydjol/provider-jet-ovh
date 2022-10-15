package kube

import (
	"github.com/crossplane/terrajet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_kube", func(r *config.Resource) {
		r.UseAsync = true

		r.ShortGroup = "kube.cloud"
		r.Kind = "Kube"
	})
	p.AddResourceConfigurator("ovh_cloud_project_kube_nodepool", func(r *config.Resource) {
		r.References["kube_id"] = config.Reference{
			Type: "ProjectKube",
		}
		r.UseAsync = true

		r.ShortGroup = "kube.cloud"
		r.Kind = "NodePool"
	})
	p.AddResourceConfigurator("ovh_cloud_project_kube_iprestrictions", func(r *config.Resource) {
		r.References["kube_id"] = config.Reference{
			Type: "ProjectKube",
		}

		r.ShortGroup = "kube.cloud"
		r.Kind = "IpRestrictions"
	})
}

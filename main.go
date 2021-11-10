package main

import (
        "github.com/hashicorp/terraform/plugin"
        "github.com/hashicorp/terraform/terraform"
        "terraform-provider-aci/aci"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return aci.Provider()
                },
        })
}
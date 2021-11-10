# terraform-provider-aci

This is the repository for the Terraform Cisco ACI provider which is used to work with Terraform and Cisco ACI.

### Provides:

### Data Sources:

> fv

```
aci_fvAp
aci_fvAEPg
aci_fvBD
aci_fvCtx
aci_fvRsBd
aci_fvRsCons
aci_fvRsCtx
aci_fvRsDomAtt
aci_fvRsProv
aci_fvSubnet
aci_fvTenant
```

> phys

```
aci_phys_dom
```

> tag

```
aci_tagInst
```

> vmm

```
aci_vmm_dom
```

> vz

```
aci_vzBrCP
aci_vzConsLbl
aci_vzConsSubjLbl
aci_vzEntry
aci_vzFilter
aci_vzProvLbl,
aci_vzProvSubjLbl
aci_vzRsSubjFiltAtt
aci_vzSubj
```

### Resources

> fv

```
aci_fvAp
aci_fvAEPg
aci_fvBD
aci_fvCtx
aci_fvRsBd
aci_fvRsCons
aci_fvRsCtx
aci_fvRsDomAtt
aci_fvRsProv
aci_fvSubnet
//"aci_fvTenant
```

> tag

```
aci_tagInst
```

> vz

```
aci_vzBrCP
aci_vzConsLbl
aci_vzConsSubjLbl
aci_vzEntry
aci_vzFilter
aci_vzProvLbl
aci_vzProvSubjLbl
aci_vzRsSubjFiltAtt
aci_vzSubj
```

# Building

Download the latest revision of the master branch then use the go compiler to generate the binary.

```
cd "${GOPATH}
go get gitlab.com:simon.birtles/terraform-providers/terraform-provider-aci
cd ./src/gitlab.com/terraform-providers/terraform-provider-aci
go get
go build -o terraform-provider-aci
```

# Installing

Download the appropriate build for your system from the release page.

Store the binary somewhere on your filesystem such as '/usr/local/bin'.

Then edit the '~/.terraformrc' file of the user running terraform to include the provider's path.

The resulting file should include the following:

```
providers {
    aci = "/path/to/terraform-provider-aci"
}
```

Also see the _Usage_ section if you do not want to use the .terraformrc file and would prefer to use the `-plugin-dir=` option of `terraform init`.

# Debugging

Enable debug mode by `export TF_LOG=DEBUG`.

For addtional options for debugging refer to [terraform documentation](https://www.terraform.io/docs/internals/debugging.html)

# Usage

## Provider

This terraform provider requires at least Terraform v0.11.8 to run.

Run `terraform init -plugin-dir=....` where `....` is where the terraform-provider-aci complied binary is located before plan or apply.

```
variable "aci_username" {
    description = "aci_username, can use env var TF_VAR_ACI_APIC_USERNAME"
}
variable "aci_password" {
    description = "aci_password, can use env var TF_VAR_ACI_APIC_PASSWORD"
}
variable "aci_apic" {
    description = "aci_server as fqdn or ipv4 address, can use env var TF_VAR_ACI_APIC"
}
variable "aci_ignore_ssl" {
    description = "aci_ignore_ssl (True/False), can use env var TF_VAR_ACI_APIC_ALLOW_UNVERIFIED_SSL"
}

provider "aci" {
    username = "${var.aci_username}"
    password = "${var.aci_password}"
    apic = "${var.aci_apic}"
    allow_unverified_ssl = "${var.aci_ignore_ssl}"
    sync_delay = 600
}
```

### Data Sources

#### Tenant [`aci_fvTenant`]

Gets attibutes of an existing Tenant `fv:Tenant`.

- `parent_id` - (Optional) The terraform ID of the parent object. For Tenants this defaults to "uni" and should be the same if provided.
- `desc` - (Exported) The Tenant description.
- `dn` - (Exported) The APIC DN of the Tenant object.
- `name` - (Required) The Tenant name.
- `ownerkey` - (Exported) An arbitary key for enabling clients to own their data for entity correlation.
- `ownertag` - (Exported) A tag for enabling clients to add their own data. For example, to indicate who created this object

```
example
```

#### Application Profile [`aci_fvAp`]

Gets attibutes of an existing Application Profile `fv:Ap` in the provided tenant.

- `tenant_id` - (Required) The terraform ID of the parent Tenant.
- `desc` - (Exported) The Application Profile description.
- `dn` - (Exported) The APIC DN of the Application Profile object.
- `name` - (Required) The Application Profile name.
- `ownerkey` - (Exported) An arbitary key for enabling clients to own their data for entity correlation.
- `ownertag` - (Exported) A tag for enabling clients to add their own data. For example, to indicate who created this object
- `priority_level` - (Exported) The QoS Priority Level for objects contained within this Application Profile.

```
example
```

### Resources

#### Application Profile [`aci_fvAp`]

Creates an Application Profile `fv:Ap` in the provided tenant.

- `tenant_id` - (Required) The terraform ID of the parent Tenant. Forces new on change.
- `desc` - (Optional) The Application Profile description.
- `dn` - (Computed) The APIC DN of the created Application Profile object.
- `name` - (Required) The Application Profile name. Forces new on change.
- `ownerkey` - (Optional) An arbitary key for enabling clients to own their data for entity correlation.
- `ownertag` - (Optional) A tag for enabling clients to add their own data. For example, to indicate who created this object
- `priority_level` - (Optional) The QoS Priority Level for objects contained within this Application Profile.

```
example
```

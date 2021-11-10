package vmm

type DomP struct {
	Dn							string `json:"dn,omitempty" schema:"dn"`
	OwnerKey					string `json:"ownerKey,omitempty" schema:"ownerkey"`
	OwnerTag					string `json:"ownerTag,omitempty" schema:"ownertag"`
	Name						string `json:"name,omitempty" schema:"name"`
}
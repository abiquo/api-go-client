package abiquo_api

type DatastoreTierCollection struct {
	AbstractCollection
	Collection []DatastoreTier
}

type DatastoreTier struct {
	DTO
	ID                      int    `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	Description             string `json:"description,omitempty"`
	Enabled                 bool   `json:"enabled,omitempty"`
	DefaultAllowed          bool   `json:"defaultAllowed,omitempty"`
	StorageAllocationPolicy string `json:"storageAllocationPolicy,omitempty"`
}

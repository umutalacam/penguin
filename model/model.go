package model

import "errors"

type Vault struct {
	Version    string    `json:"version"`
	Type       string    `json:"type"`
	Encryption string    `json:"encryption"`
	Items      []*Entity `json:"items"`
}

type Entity struct {
	Type  string    `json:"type"`
	Name  string    `json:"name"`
	Value *string   `json:"value,omitempty"`
	Items []*Entity `json:"items,omitempty"`
}

func (v *Vault) GetRootEntity() *Entity {
	// We need a root node
	return &Entity{
		Name:  "VaultRootEntity",
		Items: v.Items,
	}
}

// GetEntity Finds the entity in the directory and returns a pointer to entity
func (e *Entity) GetEntity(searchKey string) (*Entity, bool) {
	var credentials = e.Items
	// Linear search in doc
	for i := 0; i < len(credentials); i++ {
		var credential = credentials[i]
		if credential.Name == searchKey {
			return credential, true
		}
	}
	return nil, false
}

func (e *Entity) DeleteEntity(entityName string) error {
	var index = -1
	var items = e.Items
	// Linear search the index
	for i := 0; i < len(items); i++ {
		var item = items[i]
		if item.Name == entityName {
			index = i
		}
	}
	if index == -1 {
		return errors.New("entity does not exist")
	}
	// Delete entity
	e.Items = append(e.Items[:index], e.Items[index+1:]...)
	return nil
}

// CreateCredentialEntity Add credential value to the directory
func (e *Entity) CreateCredentialEntity(entity *Entity) {
	credentialFound, isCredentialFound := e.GetEntity(entity.Name)
	if isCredentialFound {
		// Update if already exist
		credentialFound.Type = entity.Type
		credentialFound.Value = entity.Value
	} else {
		// Create a new one
		e.Items = append(e.Items, entity)
	}
}

// CreateDirectoryEntity Add new directory value to the directory
func (e *Entity) CreateDirectoryEntity(directoryName string) *Entity {
	var directoryEntity = &Entity{
		Type:  "directory",
		Name:  directoryName,
		Items: []*Entity{},
	}
	e.Items = append(e.Items, directoryEntity)
	return directoryEntity
}

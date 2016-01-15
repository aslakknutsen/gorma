package gorma

import "github.com/raphael/goa/design"

// NewStorageGroup creates a StorageGroup structure by parsing
// the APIDefinition and creating all the necessary Stores and Models
func NewStorageGroup(a *design.APIDefinition) (*StorageGroup, error) {
	sg := &StorageGroup{}
	sg.api = a
	sg.RelationalStore = NewRelationalStore()
	err := sg.Parse()
	return sg, err
}

func (sg *StorageGroup) Parse() error {

	err := sg.api.IterateVersions(func(v *design.APIVersionDefinition) error {
		err := v.IterateUserTypes(func(t *design.UserTypeDefinition) error {
			if t.Type.IsObject() {
				name := t.TypeName
				m, err := NewRelationalModel(name, t)
				if err != nil {
					return err
				}
				sg.RelationalStore.Models[name] = m
			}
			return nil
		}) // IterateUserTypes
		return err
	}) // IterateVersions
	if err != nil {
		return err
	}
	err = sg.RelationalStore.ResolveRelationships()

	return err
}
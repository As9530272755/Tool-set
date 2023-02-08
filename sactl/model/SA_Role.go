package model

type SA struct {
	NameSpace string
	Name      string
	API       string
}

func NewSa(namespace, name, api string) *SA {
	return &SA{
		NameSpace: namespace,
		Name:      name,
		API:       api,
	}
}

type RoleBind struct {
	NameSpace string
	Name      string
}

func NewRoleBind(namespace, name string) *RoleBind {
	return &RoleBind{
		namespace,
		name,
	}
}

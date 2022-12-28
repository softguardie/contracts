package contracts

type Identity struct {
	Id        string   `json:"id"`
	Namespace string   `json:"namespaces"`
	Roles     []string `json:"roles"`
}

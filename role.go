package contracts

type Role struct {
	Id         string   `json:"id"`
	Namespaces []string `json:"namespaces"`
}

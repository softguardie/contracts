package contracts

type Policy struct {
	Subject   string `json:"id"`
	Namespace string `json:"namespaces"`
	Object    string `json:"object"`
	Effect    string `json:"effect"`
}

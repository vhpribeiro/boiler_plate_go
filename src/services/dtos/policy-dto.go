package dtos

type PolicyDto struct {
	Role     string `json:"role"`
	Domain   string `json:"domain"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

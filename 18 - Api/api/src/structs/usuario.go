package structs

type Usuario struct {
	Nome     string   `json:"nome,omitempty"`
	Idade    uint8    `json:"idade,omitempty"`
	Endereco Endereco `json:"endereco,omitempty"`
}

type Endereco struct {
	Logradouro string `json:"logradouro,omitempty"`
	Numero     int    `json:"numero,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Cidade     string `json:"cidade,omitempty"`
}

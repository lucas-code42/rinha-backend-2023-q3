package domain

type Pessoa struct {
	Id         string   `json:"id,omitempty"`
	Apelido    string   `json:"apelido,omitempty"`
	Nome       string   `json:"nome,omitempty"`
	Nascimento string   `json:"nascimento,omitempty"`
	Stack      []string `json:"stack,omitempty"`
}

type PessoaDto struct {
	Id         string `json:"id,omitempty"`
	Apelido    string `json:"apelido,omitempty"`
	Nome       string `json:"nome,omitempty"`
	Nascimento string `json:"nascimento,omitempty"`
	Stack      string `json:"stack,omitempty"`
}

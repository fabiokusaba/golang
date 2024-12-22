package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang")

	// Criando um map de usuário
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	// Map aninhado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)
	// Deletando uma chave
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	// Adicionando uma chave
	usuario2["signo"] = map[string]string{
		"nome": "Aquário",
	}
	fmt.Println(usuario2)
}

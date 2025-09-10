# go-convert

**go-convert** é uma ferramenta de linha de comando (CLI) desenvolvida em Go para facilitar conversões rápidas, como codificação/decodificação Base64, decodificação de payloads de tokens JWT e geração de senhas seguras.

## Funcionalidades

- **Base64**: Codifica texto para Base64 ou decodifica Base64 para texto.
- **JWT**: Decodifica e exibe o payload de um token JWT em formato JSON legível.
- **Password**: Gera senhas seguras e aleatórias com opções de personalização.

## Instalação

1. **Clone o repositório**

   ```sh
   git clone https://github.com/GustavoM96/go-convert.git
   cd go-convert
   ```

2. **Compile o projeto**

   ```sh
   go build -o go-convert
   ```

3. O binário `go-convert` estará disponível no diretório do projeto.

## Uso

### Comando Base64

Codifique uma string para Base64:

```sh
./go-convert base64 --encode "texto a ser codificado"
# ou
./go-convert base64 -e "texto a ser codificado"
```

Decodifique uma string Base64:

```sh
./go-convert base64 --decode "dGV4dG8gYSBzZXIgY29kaWZpY2Fkbw=="
# ou
./go-convert base64 -d "dGV4dG8gYSBzZXIgY29kaWZpY2Fkbw=="
```

> **Nota:** As flags `--encode` e `--decode` são mutuamente exclusivas.

### Comando JWT

Decodifique o payload de um token JWT:

```sh
./go-convert jwt --decode "<seu_token_jwt>"
# ou
./go-convert jwt -d "<seu_token_jwt>"
```

O payload será exibido em formato JSON indentado.

### Comando Password

Gere uma senha segura:

```sh
./go-convert password --length 16 --times 3
# ou
./go-convert password -l 16 -t 3
```

## Estrutura do Projeto

```
.
├── main.go                # Ponto de entrada da aplicação
├── go.mod                 # Gerenciamento de dependências
├── cmd/
│   ├── root.go            # Comando raiz e inicialização do CLI
│   ├── base64.go          # Implementação do comando base64
│   ├── jwt.go             # Implementação do comando jwt
│   └── password.go        # Implementação do comando password
└── .gitignore
```

- `main.go`: Executa o comando raiz.
- `cmd/root.go`: Define o comando principal e inicializa a CLI.
- `cmd/base64.go`: Implementa as operações de codificação e decodificação Base64.
- `cmd/jwt.go`: Implementa a decodificação do payload de tokens JWT.
- `cmd/password.go`: Implementa a geração de senhas seguras.

## Dependências

- [Cobra](https://github.com/spf13/cobra): Framework para criação de CLIs em Go.

## Exemplos

Codificar:

```sh
./go-convert base64 -e "Go é incrível!"
```

Decodificar:

```sh
./go-convert base64 -d "R28gw6kgacOtc3RpbW8h"
```

Decodificar JWT:

```sh
./go-convert jwt -d "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiZ29oYSJ9.signature"
```

Gerar senha segura:

```sh
./go-convert password -l 16 -s -n -u
```

## Licença

Este projeto está sob a licença MIT.

---

Para mais detalhes, consulte os arquivos de código-fonte:

- `main.go`
- `cmd/root.go`
- `cmd/base64.go`
- `cmd/jwt.go`
- `cmd/password.go`

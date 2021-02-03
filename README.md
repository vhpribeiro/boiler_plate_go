<h1 align="center"> Controle de Acesso </h1>

## Descrição do Projeto

<p align="justify"> Projeto responsável por cuidar da parte de autorização dos sistemas Ambev. </p>

## Para Começar

Para podermos executar o projeto precisamos instalar alguns programas em nossa máquina para que tudo funcione perfeitamente.

- Visual Studio Code(vscode) é um editor de código-fonte desenvolvido pela Microsoft que você poderá usar para editar e subir seu código. [Download Visual Studio Code](https://code.visualstudio.com/download)

- Go(Golang) é uma linguagem de programação de código aberto que facilita a construção de software simples, confiável e eficiente. [Download Golang](https://golang.org/doc/install#download)

- PostgreSQL é sistema gerenciador de banco de dados objeto relacional, desenvolvido como projeto de código aberto. [Download PostgreSQL](https://www.postgresql.org/download/)

## Como Instalar

Inicialmente precisamos clonar o projeto da seguinte maneira:

`1. git clone https://github.com/vhpribeiro/controle_acesso_core.git`

`2. cd controle_acesso_core`

`3. cd src`

Após clonar de o seguinte comando:

`4. go get -u`

Isso irá baixar e instalar todas as dependências do projeto.

Para executar o projeto, basta executar o seguinte comando:

`5. go run .`

Você saberá quando o projeto estiver rodando quando aparecer a seguinte mensagem em seu Cmd:

```
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.1.17
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8000
```
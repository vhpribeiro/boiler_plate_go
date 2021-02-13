<h1 align="center"> Boiler Plate Go </h1>

## Descrição do Projeto

<p align="justify"> Boiler plate para ser usado em projetos desenvolvidos em go</p>

## Para Começar

Para executar o projeto precisamos instalar alguns programas em nossa máquina para que tudo funcione perfeitamente.

- **Visual Studio Code(vscode)** é um editor de código-fonte desenvolvido pela Microsoft que você poderá usar para editar e subir seu código. [Download Visual Studio Code](https://code.visualstudio.com/download)

- **Go(Golang)** é uma linguagem de programação de código aberto que facilita a construção de software simples, confiável e eficiente. [Download Golang](https://golang.org/doc/install#download)

- **PostgreSQL** é sistema gerenciador de banco de dados objeto relacional, desenvolvido como projeto de código aberto. [Download PostgreSQL](https://www.postgresql.org/download/)

- **Docker** é uma plataforma aberta para desenvolvimento, envio e execução de aplicativos. O Docker permite que você separe seus aplicativos de sua infraestrutura para que possa entregar o software rapidamente. [Downalod Docker](https://docs.docker.com/get-docker/)

## Variáveis de ambiente

Todas variáveis de ambiente que o projeto estão dentro de **src/configuration/environments/environments.go** já com seus valores defaults configurados.

As mesmas variáveis também se encontram configuradas no docker-compose do projeto.

## Como executar o projeto

#### Executar toda aplicação localmente

Inicialmente precisamos clonar o projeto da seguinte maneira:

`1. git clone https://github.com/vhpribeiro/controle_acesso_core.git`

`2. cd boiler_plate_go`

`3. cd src`

Após clonar o projeto, execute o seguinte comando:

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

#### Executar com Docker

Para rodar a aplicação com Docker você precisa estar na pasta raiz no projeto, e rodar o seguinte comando:

`docker-compose up`

Após aplicar esse comando em algum Cmd, ele irá subir todas as imagens que estão dentro dele.

Ao rodar o Docker Compose, irão subir as seguintes imagens:
 - Aplicação / Porta: 8000
 - Postgres / Porta: 5432
 - PgAdmin / Porta: 5050

Caso queira acabar com o funcionamento use `CTRL + C` no Cmd que estiver usando e depois use o comando:

`docker-compose down`

Esse comando fara que suas imagens que estavam rodando parem de rodar.

###### Passo a passo para configurar database

No pgAdmin, é necessário criar um Server com as seguintes configurações:
 - General - Nome do Servidor: Postgres 13
 - Connection - Host name/address: postgres-container
 - Connection - Port - 5432
 - Connection - Maintenance database - postgres
 - Connection - Username - postgres
 - Connection - Password - admin

Feito isso, basta criar uma base com o nome **casbin** e a aplicação poderá ser executada
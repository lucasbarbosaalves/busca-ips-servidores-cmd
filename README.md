# Aplicação em Golang para busca de IP e servidores

Esta aplicação em Golang permite que você busque informações sobre um IP ou servidor através do terminal. Ela utiliza a biblioteca `github.com/urfave/cli` para obter os parâmetros digitados no terminal através da flag `--host`. Caso você queira buscar informações sobre um servidor, é possível passar o nome do servidor como argumento.

## Instruções para rodar

1. Certifique-se de ter o Go instalado em seu sistema. Você pode baixá-lo em https://golang.org/dl/.

2. Clone este repositório para o seu sistema:

   ```shell
   git clone https://github.com/lucasbarbosaalves/busca-ips-servidores-cmd.git

3. cd busca-ips-servidores-cmd

4. Instale as dependências:

   ```shell
   go mod download // Irá Listar as dependências e baixá-las
   ```

5. Compile o projeto:

   ```shell
   go build
   ```

6. Execute o projeto:

   ```shell
   ./busca-ips-servidores-cmd --host <ip ou servidor>
   ```
    Substitua <IP ou nome do servidor> pelo IP ou nome do servidor que você deseja buscar informações.

7. A aplicação exibirá os resultados da busca no terminal.








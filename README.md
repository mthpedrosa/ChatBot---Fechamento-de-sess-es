# Cron de Fechamento de Sessões Antigas

Este projeto implementa um cron simples em Golang que consulta um banco de dados MongoDB para identificar e fechar sessões abertas há mais de 24 horas.

## Funcionalidades

* **Consulta ao MongoDB:** Conecta-se a um banco de dados MongoDB especificado.
* **Identificação de Sessões Antigas:** Localiza sessões cuja data de criação seja anterior a 24 horas.
* **Fechamento de Sessões:** Atualiza o status das sessões identificadas para "fechado" no banco de dados.
* **Agendamento Cron:** Executa a verificação e fechamento de sessões em um intervalo configurável.

## Pré-requisitos

* Go (versão 1.16 ou superior)
* MongoDB
* Driver Go para MongoDB (go.mongodb.org/mongo-driver)

## Instalação

1.  Clone o repositório:

    ```bash
    git clone https://github.com/mthpedrosa/EagleChat-Cron-de-Sessoes
    ```

2.  Navegue até o diretório do projeto:

    ```bash
    cd EagleChat-Cron-de-Sessoes
    ```

3.  Instale as dependências:

    ```bash
    go mod tidy
    ```

4.  Configure as variáveis de ambiente:

    * `MONGODB_URI`: String de conexão do MongoDB.
    * `DATABASE_NAME`: Nome do banco de dados.
    * `COLLECTION_NAME`: Nome da coleção de sessões.
    * `CRON_SCHEDULE`: Expressão cron para agendar a execução (opcional, padrão: "@hourly").

    Você pode definir essas variáveis em um arquivo `.env` na raiz do projeto ou como variáveis de ambiente do sistema.

## Execução

Para executar o cron, utilize o seguinte comando:

```bash
go run main.go

# WeatherCEP

O projeto WeatherCEP é uma aplicação web que permite aos usuários buscar informações meteorológicas atuais através do CEP (Código de Endereçamento Postal) brasileiro. Utilizando uma interface simples e intuitiva, o usuário pode obter dados como temperatura, condições climáticas, umidade, entre outros, de forma rápida e precisa.

## Como Testar

Para testar a funcionalidade do projeto, você tem duas opções:

1. **Acessar Online**

   Você pode acessar a seguinte URL:

   [https://weathercep-epk5h7ir7q-uc.a.run.app/weather/51020000](https://weathercep-epk5h7ir7q-uc.a.run.app/weather/51020000)

   Essa URL irá retornar as informações meteorológicas atuais para o CEP 51020000. Você pode substituir o CEP na URL pelo desejado para buscar informações de diferentes localidades.

2. **Rodar Localmente com Docker Compose**

   Se preferir testar localmente, você pode utilizar o Docker Compose. Para isso, execute o seguinte comando no terminal:

   ```sh
   docker-compose -f docker-compose.dev.yml up
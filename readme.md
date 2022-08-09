# Golang Production Ready Bootstrap Web Server.

This is a startup template for Golang.

## Stracture

### Docker Containers
nginx , mysql and api(Golang).

### Nginx
- https2 configuration
- automate script for /etc/hosts ( so you can use https://www.yourdomain.com for your local development)

### Mysql
Can be swap with any db container but I like to use mysql.

### Golang Rest Api.
- [Air](https://github.com/cosmtrek/air) for hot reload (This will detect changes and rebuild the app on file change)
- folder struct
- DB integration
- example route , controller , 
- middlewares (rest logger , session , cors)

## Installation
### Requirements.
- [docker](https://www.docker.com/)
- [docker-compose](https://docs.docker.com/compose/)

Clone the repository to your computer.

```git
git clone git@github.com:Djancyp/go-rest.git
```

## Usage

### first time runners  (build required)
This will build the required containers and start them. Might take sometime.
```bash
docker-compose up --build
```
### Run
```bash
docker-compose up
``````
🚀 Yuhii app is up and running.
Now you can check server on browser [https://www.golang.loc](https://www.golang.loc)

### Build in example endpoints.
- /example (GET) (Get all examples)
- /example(POST)
```json
{
  "name":"test"
}
```
- /example/{id} (GET)

## Config
All configs can be set from .env file.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

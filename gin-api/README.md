### Example gin application for API

- application environment configuration with viper
- setup database with mongodb behind and using HOF (higher order function) to pass db repository to routes
- use of middleware to support JWT authentication
- two different server for public and admin is running using goroutine
- pre-commit setup for linting and code error checking in local development

### ✔️ Setup for Development

- run `make setup` and read the [Makefile](https://github.com/the-robot/gopher/blob/master/gin-api/Makefile)

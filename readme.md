run with docker run
```docker run -p 8080:8080 --name golangcleanarch --network="my-shared-network" --rm go-clean-arch:1.1.1```
please check the connection string of mysql at `main.go`

run with docker compose
```docker compose up --build -d```
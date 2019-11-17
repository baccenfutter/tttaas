# TTTaaS

*Tic-Tac-Toe-as-a-Service* is a (https://goa.design/)[Goa] Microservice
implementation.

## Requirements

- Go-1.12
- GNU/make
- protoc (v3)

## Installation

```bash
go get -u github.com/baccenfutter/tttaas
go install github.com/baccenfutter/cmd/...
```

### Developers

```bash
git clone https://github.com/baccenfutter/tttaas
make all
```

### Docker

```bash
docker build -t tttaas .
docker run -it --rm tttaas
```

## Usage

Start the server:

```bash
tictactoe
```

### Start New Game

```bash
tictactoe-cli game new
```

### Get Running Game By ID

```bash
tictactoe-cli game get -board <id>
```

### Play A Move

```bash
tictactoe-clie game move -board <id> -square <square>
```

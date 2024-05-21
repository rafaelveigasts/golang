module api-devbook

go 1.22.0

require github.com/gorilla/mux v1.8.1

require (
	github.com/badoux/checkmail v1.2.4 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	golang.org/x/crypto v0.23.0 // indirect
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/tdewolff/parse v2.3.4+incompatible
)

// docker run -it --rm \
//     -w "<DEVBOOK>" \
//     -e "air_wd=<DEVBOOK>" \
//     -v $(pwd):<DEVBOOK> \
//     -p <PORT>:<9000> \
//     cosmtrek/air
//     -c <CONF>

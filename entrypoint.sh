
# wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"
# Watch your .go files and invoke go build if the files changed.
CompileDaemon --build="go build -o SupplySideService/main SupplySideService/main.go"  --command=SupplySideService/main
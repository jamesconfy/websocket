echo -n "Mode: "
read MODE

case $MODE in
dev)
    export $(cat .env.dev | xargs) && make run
    ;;

docker)
    export $(cat .env.docker | xargs) && docker compose up -d --build
    ;;

prod)
    export $(cat .env.prod | xargs) && make deploy
    ;;
*)
    echo -n "Unknown command"
    ;;
esac

# Deploy secrets to fly
# awk '{system("flyctl secrets set " $1)}'

for service in $(docker-compose config --services); do
    echo -e "\n $service \n" | tr '[:lower:]' '[:upper:]';
    eval "docker compose logs $service";
done
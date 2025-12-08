#!/bin/bash

set -e

COMMAND=$1

# Colors
GREEN="\033[0;32m"
CYAN="\033[0;36m"
NC="\033[0m" # No Color

echo -e "${CYAN}== Ledgerly Local Dev Script ==${NC}"

# Load .env if it exists
if [ -f .env ]; then
  echo -e "${GREEN}Loading environment variables from .env...${NC}"
  export $(grep -v '^#' .env | xargs)
else
  echo "No .env file found. Using docker-compose environment."
fi

case "$COMMAND" in

  up)
    echo -e "${GREEN}Starting Ledgerly using docker-compose...${NC}"
    docker compose up -d --build
    echo -e "${CYAN}Containers started. Logs below (Ctrl+C to exit):${NC}"
    docker compose logs -f
    ;;

  down)
    echo -e "${GREEN}Stopping containers...${NC}"
    docker compose down
    ;;

  rebuild)
    echo -e "${GREEN}Forcing rebuild of API image...${NC}"
    docker compose build api
    docker compose up -d
    docker compose logs -f api
    ;;

  logs)
    echo -e "${GREEN}Showing live logs...${NC}"
    docker compose logs -f
    ;;

  reset-db)
    echo -e "${GREEN}⚠ WARNING: This will delete all PostgreSQL data.${NC}"
    read -p "Are you sure? (y/n) " confirm
    if [ "$confirm" = "y" ]; then
      docker compose down
      docker volume rm ledgerly_postgres_data || true
      echo -e "${GREEN}DB volume deleted.${NC}"
      docker compose up -d
    else
      echo "Aborted."
    fi
    ;;

  *)
    echo "Usage: ./local_run.sh [up|down|rebuild|logs|reset-db]"
    exit 1
    ;;
esac

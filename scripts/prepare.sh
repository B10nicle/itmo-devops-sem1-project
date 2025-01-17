#!/bin/bash

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

COMPILE_TO="bin/main"
COMPILE_FROM="./main.go"

export PGPASSWORD="val1dat0r"

run_command() {
  $1 &> /dev/null
  if [ $? -ne 0 ]; then
    echo -e "${RED}$2${NC}"
    exit 1
  fi
  echo -e "${GREEN}$3${NC}"
}

echo -e "${YELLOW}Creating table prices${NC}"
run_command "psql -h localhost -p 5432 -U validator -d project-sem-1 -f migrations/01_create_prices_table.sql" \
            "Failed to create table prices" \
            "Table prices has been successfully created"

echo -e "${YELLOW}Compiling application${NC}"
run_command "go build -o $COMPILE_TO $COMPILE_FROM" \
            "Failed to compile the application" \
            "Application has been successfully compiled to $COMPILE_TO"
#!/bin/bash

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

APP_BIN_PATH="bin/main"

exit_with_error() {
  echo -e "${RED}$1${NC}"
  exit 1
}

if [ ! -f "$APP_BIN_PATH" ]; then
  exit_with_error "Application is not compiled or not found at path $APP_BIN_PATH"
fi

echo -e "${YELLOW}Starting the application${NC}"
./"$APP_BIN_PATH" &

APP_PID=$!

if ! ps -p $APP_PID > /dev/null 2>&1; then
  exit_with_error "Error while starting the application"
fi

echo -e "${GREEN}Application started with PID: $APP_PID${NC}"
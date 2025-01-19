#!/bin/bash

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${YELLOW}Making scripts executable${NC}"

chmod +x ./scripts/prepare.sh
if [ $? -ne 0 ]; then
  echo -e "${RED}Failed to make prepare.sh executable.${NC}"
  exit 1
else
  echo -e "${GREEN}prepare.sh is now executable.${NC}"
fi

chmod +x ./scripts/run.sh
if [ $? -ne 0 ]; then
  echo -e "${RED}Failed to make run.sh executable.${NC}"
  exit 1
else
  echo -e "${GREEN}run.sh is now executable.${NC}"
fi

chmod +x ./scripts/all_tests.sh
if [ $? -ne 0 ]; then
  echo -e "${RED}Failed to make all_tests.sh executable.${NC}"
  exit 1
else
  echo -e "${GREEN}all_tests.sh is now executable.${NC}"
fi

echo -e "${GREEN}Done. All scripts are now executable!${NC}"
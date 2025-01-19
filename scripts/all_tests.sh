#!/bin/bash

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

for level in {1..3}; do
  echo -e "${YELLOW}Running tests for level $level${NC}"
  ./scripts/tests.sh $level
  if [ $? -ne 0 ]; then
    echo -e "${RED}Tests for level $level failed.${NC}"
    exit 1
  else
    echo -e "${GREEN}Tests for level $level passed successfully!${NC}"
  fi
done

echo -e "${GREEN}Done. All tests passed successfully!${NC}"
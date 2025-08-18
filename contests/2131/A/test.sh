#!/usr/bin/env bash

diff --brief <(cat input | go run ./lever.go) <(cat ./output)

if [ $? -eq 1 ]
then
  echo "FAIL"
else
  echo "PASS"
fi

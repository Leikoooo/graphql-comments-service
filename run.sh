#!/bin/bash

set -e

migrate -database ${DATABASE_URL} -path src/migrations up

./app

#!/bin/sh

set -e

cp .env backend/.env
cd backend
go run cmd/main.go

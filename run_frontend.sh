#!/bin/sh

set -e

cp .env frontend/.env
cd frontend
npm run dev

#!/bin/bash
docker build -t michaelpeterswa/shortlink2-backend backend/
docker build -t michaelpeterswa/shortlink2-frontend frontend/
docker push michaelpeterswa/shortlink2-backend
docker push michaelpeterswa/shortlink2-frontend
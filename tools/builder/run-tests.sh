#!/bin/bash

cd agent
go test ./...
cd ..

cd services/agentcentral
go test ./...
cd ../..


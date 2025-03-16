#!/bin/bash

# this file is only for github workflow.

GO=go

cd agent
$GO test ./...
cd ..

cd services/agentcentral
$GO test ./...
cd ../..

cd baselibs/multifs
$GO test ./...
cd ../..

cd baselibs/abstractions
$GO test ./...
cd ../..

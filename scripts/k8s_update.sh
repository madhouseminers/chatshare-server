#!/bin/sh

kubectl set image -f deployments/deployment.yaml chatshare=madhouseminers/chatshare:"$1" --record

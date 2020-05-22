#!/bin/sh

kubectl apply -f deployments/namespace.yaml
kubectl apply -f deployments/secrets.yaml
kubectl apply -f deployments/deployment.yaml
kubectl apply -f deployments/service.yaml
kubectl apply -f deployments/policies.yaml

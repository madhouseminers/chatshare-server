#!/bin/sh

docker build -f build/package/Dockerfile -t madhouseminers/chatshare:"$1" .
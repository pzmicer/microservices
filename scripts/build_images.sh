#! /bin/zsh

minikube image build -t svc-hello ~/Projects/microservices/hello
minikube image build -t svc-second ~/Projects/microservices/second



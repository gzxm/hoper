#!/bin/bash

echo $CA |base64 -d > ca.crt
echo $CACRT |base64 -d > dev.crt
echo $CAKEY |base64 -d > dev.key

server=https://hoper.xyz:6443
kubectl config set-cluster k8s --server=${server} --certificate-authority=ca.crt --embed-certs=true --kubeconfig=/root/.kube/config
kubectl config set-credentials dev --client-certificate=dev.crt --client-key=dev.key --embed-certs=true --kubeconfig=/root/.kube/config
kubectl config set-context dev --cluster=k8s --user=dev --kubeconfig=/root/.kube/config
kubectl config use-context dev --kubeconfig=/root/.kube/config
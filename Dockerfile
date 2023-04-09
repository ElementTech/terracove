FROM ghcr.io/devops-infra/docker-terragrunt:aws-azure-gcp-tf-1.4.4-tg-0.45.2
COPY terracove /usr/bin/terracove
WORKDIR /data
ENTRYPOINT ["/usr/bin/terracove"]
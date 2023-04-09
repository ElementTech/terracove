FROM devopsinfra/docker-terragrunt:tf-1.4.4-tg-0.45.2
COPY terracove /usr/bin/terracove
ENTRYPOINT ["/usr/bin/terracove"]
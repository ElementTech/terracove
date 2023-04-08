FROM alpine:3.17
COPY terracove /usr/bin/terracove
ENTRYPOINT ["/usr/bin/terracove"]
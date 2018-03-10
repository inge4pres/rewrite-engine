FROM debian

COPY rewrite-engine /rewrite-engine

ENTRYPOINT ["/rewrite-engine"]
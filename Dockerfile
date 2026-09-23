ARG BASE_IMAGE=ghcr.io/wntrtech/scratch:v1.0.0
FROM ${BASE_IMAGE}


EXPOSE 8080/tcp
COPY publish/ /

ENTRYPOINT ["/server", "web"]
HEALTHCHECK --start-period=30s --start-interval=5s --interval=30s --timeout=5s --retries=3 CMD ["/server", "health"]

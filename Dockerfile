FROM gcr.io/distroless/static:nonroot

WORKDIR /app/
COPY foobar /app/

ENTRYPOINT ["/app/foobar"]

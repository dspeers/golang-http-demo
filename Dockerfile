FROM scratch
EXPOSE 8080
ENTRYPOINT ["/golang-http-demo"]
COPY ./bin/ /
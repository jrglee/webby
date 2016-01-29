FROM scratch
ADD webby /
EXPOSE 8080
ENTRYPOINT ["/webby"]

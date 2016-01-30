FROM debian:jessie
ADD webby /
EXPOSE 8080
CMD ["/webby"]

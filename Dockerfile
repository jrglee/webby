FROM scratch
ADD webby /
EXPOSE 8080
CMD ["/webby"]

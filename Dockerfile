FROM alpine
ADD cmd/superhero /superhero
ADD configs/config.yml /config.yml
RUN ["chmod", "+x", "/superhero"]
ENTRYPOINT [ "/superhero" ]

FROM alpine
ADD cmd/superhero /superhero
RUN ["chmod", "+x", "/superhero"]
ENTRYPOINT [ "/superhero" ]

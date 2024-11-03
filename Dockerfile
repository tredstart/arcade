FROM archlinux:latest

RUN pacman -Syu --noconfirm && \
    pacman -S --noconfirm go make

WORKDIR /app
COPY . .
RUN go mod tidy
RUN make build

EXPOSE 8080

ENTRYPOINT ["./main"]

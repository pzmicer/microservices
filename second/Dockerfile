FROM rust:1 as builder
WORKDIR /app
COPY . .
RUN cargo install --path .

FROM debian:bullseye-slim 
COPY --from=builder /usr/local/cargo/bin/second /second
ENV ROCKET_ADDRESS=0.0.0.0
ENV ROCKET_PORT=8080
# CMD ["/usr/local/cargo/bin/second"]
EXPOSE ${ROCKET_PORT}
CMD ["/second"]

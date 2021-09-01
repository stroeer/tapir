FROM rust:1.54.0

RUN cargo install fundoc --locked
WORKDIR /opt

ENTRYPOINT /usr/local/cargo/bin/fundoc

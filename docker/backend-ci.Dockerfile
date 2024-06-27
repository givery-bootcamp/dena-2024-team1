FROM ubuntu:22.04
RUN apt update && apt install -y ca-certificates 
WORKDIR /app
COPY backend/myapp /app
RUN chmod +x myapp

ENTRYPOINT ["./myapp"]
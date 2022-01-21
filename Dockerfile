FROM scratch
COPY build /app/bin
WORKDIR /app/bin
CMD ["./main"]

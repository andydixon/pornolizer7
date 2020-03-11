FROM ubuntu
RUN apt update; apt install -y ca-certificates; update-ca-certificates -v
EXPOSE 8080
ADD Pornolizer7 /pornolizer7
RUN chmod +x /pornolizer7
HEALTHCHECK --interval=5s --timeout=3s --retries=3 CMD curl -f http://localhost:8080 || kill -9 1
CMD /pornolizer7


FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /app
ENV TZ="Asia/Ho_Chi_Minh"

COPY build/app .
EXPOSE 9999
ENTRYPOINT [ "/app/app" ]

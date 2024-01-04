FROM golang:1.20.6-alpine3.18 AS build-stage
WORKDIR /app
COPY . ./
RUN go mod download
RUN go build -v -o /api ./cmd/api


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /api /api
COPY --from=build-stage /app/.env /
EXPOSE 8080
CMD [ "/api" ]
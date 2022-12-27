FROM golang:1.19.4-alpine AS build
WORKDIR /project
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o ./api ./main.go

FROM alpine:3.17 AS production
RUN apk --no-cache add ca-certificates
WORKDIR /project
COPY --from=build /project/api .
COPY ./env.example.yml ./env.yml
CMD [ "./api" ]
EXPOSE 5000

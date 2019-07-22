FROM node:10-alpine as front
ENV LANG C.UTF-8

WORKDIR /luizalabs-desafio
ENV API_URL=http://localhost:8000
COPY front .
RUN npm install
# Build app
RUN npm run build

FROM golang:1.12 as back
WORKDIR /luizalabs-desafio
ADD back ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o app-luizalabs main/app.go

FROM scratch
WORKDIR /luizalabs-desafio
COPY --from=front /luizalabs-desafio/dist ./dist
COPY --from=back /luizalabs-desafio/app-luizalabs .

EXPOSE 8000
EXPOSE 8001

ENTRYPOINT ["./app-luizalabs"]
FROM golang:1.13

ADD app app

EXPOSE 8000

CMD ./app


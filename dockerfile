FROM golang:1.20-alpine
WORKDIR .

COPY . .

# retrieves go-ethereum pkg, for ability to use apigen func
RUN go mod download
RUN go get -d github.com/ethereum/go-ethereum/...

# creates new .go contract instance
RUN apigen --sol SimpleToken.sol --pkg main --out SimpleToken.go

CMD [""]
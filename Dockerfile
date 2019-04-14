FROM golang:1.12


WORKDIR $GOPATH/src/github.com/theuncharted/atypeek_blockchain

# copy binary into image
COPY . .
COPY nscli .
COPY sleep.sh .

#ENV GO111MODULE on
#ENV CGO_ENABLED 0
#RUN go get
#RUN make install
RUN ./nsd init --chain-id namechain


# Configure your CLI to eliminate need for chain-id flag
RUN ./nscli config chain-id namechain
RUN ./nscli config output json
RUN ./nscli config indent true
RUN ./nscli config trust-node true

#RUN nohup ./nsd start &
RUN nohup bash -c "./nsd start &" && sleep 4

EXPOSE 8080

#ENTRYPOINT ["./nsd"]
CMD ["sh", "sleep.sh"]
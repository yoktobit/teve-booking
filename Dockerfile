FROM alpine
ADD teve-booking /teve-booking
ENTRYPOINT [ "/teve-booking" ]

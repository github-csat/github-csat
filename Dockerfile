FROM alpine
WORKDIR /app
COPY ./build/ /app/build
CMD [ "/app/build/github-csat" ]
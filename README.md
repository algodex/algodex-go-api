docker build -t algodexidxsvr:latest .
docker run --rm -p 8000:8000 algodexidxsvr:latest /app -debug

to run locally:
go build ./cmd/algodexidxsvr
go build ./cmd/algodexidxsvr-cli

./algodexidxsvr -debug

./algodexidxsvr-cli --url="http://127.0.0.1:8000" account add -p 6APKHESCBZIAAZBMMZYW3MEHWYBIT3V7XDA2MF45J5TUZG5LXFXFVBJSFY

./algodexidxsvr-cli --url="http://localhost:8000" account list -view full

put openapi3.json into that swagger page to get curl exampels et

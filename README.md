# go-tokenizator
This is to be used as an internal API to generate a secure token.

## Description

This API takes a URL and calculates a security token, and returns the URL with the added token.

Make a POST to:
```
http://localhost:3000/tokinize/
```
Headers:
```
api-key: "thisIsAnApiKay"
```
With the body:
```
{
    url: "http://somehost/somepath/avideo.m3u8"
}
```

## Startup
### Clone repository
```
git clone https://github.com/marcos979/go-tokenizator.git
```
sudo docker-compose up

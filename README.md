
<p align="center">
  <p align="center">
    <b>Github OAuth2</b>
  </p>

  <p align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/xjh22222228/github-oauth2" />
    <img src="https://img.shields.io/github/v/release/xjh22222228/github-oauth2" />
    <img src="https://img.shields.io/github/license/xjh22222228/github-oauth2" />
  </p>
</p>


## config
See [config.json](config.json)

```json
{
  "client_id": "xxx",
  "client_secret": "xxx"
}
```


## API

```js
GET localhost:7001/api/oauth/token

param:
code: string
```

## Example For JS

```js
fetch('http://localhost:7001/api/oauth/token?code=xxx', {
  method: 'GET'
}).then(res => {
  console.log(res)
  // ...
})
```


## Build
```bash
# Build Binary
go build main.go


# Run
./main
```




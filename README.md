
<p align="center">
  <p align="center">
    <b>Github OAuth2</b>
  </p>
  <p align="center">Serverless For Vercel</p>

  <p align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/xjh22222228/github-oauth2" />
    <img src="https://img.shields.io/github/v/release/xjh22222228/github-oauth2" />
    <img src="https://img.shields.io/github/license/xjh22222228/github-oauth2" />
  </p>
</p>


## config
See [config.json](api/config.json)

```json
{
  "client_id": "xxx",
  "client_secret": "xxx"
}
```


## Usage
- Fork
- Make sure to make the warehouse private
- https://github.com/apps/vercel
- ...



## Example For JS
```js
fetch('/api/oauth?code=xxx', {
  method: 'GET'
}).then(res => {
  console.log(res)
  // ...
})
```

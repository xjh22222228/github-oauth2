
<p align="center">
  <p align="center">
    <b>Github OAuth2</b>
  </p>
  <p align="center">Serverless For Vercel</p>

  <p align="center">
    <a href="README_zh-CN.md">
      <img src="https://img.shields.io/badge/lang-%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87-red.svg?longCache=true&style=flat-square">
    </a>
    <img src="https://img.shields.io/github/go-mod/go-version/xjh22222228/github-oauth2" />
    <img src="https://img.shields.io/github/v/release/xjh22222228/github-oauth2" />
    <img src="https://img.shields.io/github/license/xjh22222228/github-oauth2" />
  </p>
</p>



## Demo
- [https://github-oauth-opal.vercel.app](https://github-oauth-opal.vercel.app)
- [https://github-oauth-opal.vercel.app/api/oauth](https://github-oauth-opal.vercel.app/api/oauth)
- [https://github-oauth-opal.vercel.app/api/user](https://github-oauth-opal.vercel.app/api/user)




## config
Edit [config.json](api/config.json)

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

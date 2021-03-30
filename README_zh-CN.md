
<p align="center">
  <p align="center">
    <b>Github OAuth2</b>
  </p>
  <p align="center">基于 Vercel Serverless</p>

  <p align="center">
   <a href="README.md">
      <img src="https://img.shields.io/badge/lang-%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87-red.svg?longCache=true&style=flat-square">
    </a>
    <img src="https://img.shields.io/github/go-mod/go-version/xjh22222228/github-oauth2" />
    <img src="https://img.shields.io/github/v/release/xjh22222228/github-oauth2" />
    <img src="https://img.shields.io/github/license/xjh22222228/github-oauth2" />
  </p>
</p>



## 例子
- [https://github-oauth-opal.vercel.app](https://github-oauth-opal.vercel.app)
- [https://github-oauth-opal.vercel.app/api/oauth](https://github-oauth-opal.vercel.app/api/oauth)




## 配置
编辑 [config.json](api/config.json)

```json
{
  "client_id": "xxx",
  "client_secret": "xxx"
}
```


## 使用
- Fork
- 务必将当前仓库设为私用，否则会泄露 client_secret
- 打开 https://github.com/apps/vercel 根据流程去跑
- ...



## JS
```js
fetch('/api/oauth?code=xxx', {
  method: 'GET'
}).then(res => {
  console.log(res)
  // ...
})
```

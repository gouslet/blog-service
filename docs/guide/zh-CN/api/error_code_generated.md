# 错误码
！！IAM 系统错误码列表，由 `codegen -type=int -doc` 命令生成，不要对此文件做任何更改。
## 功能说明
如果返回结果中存在 `code` 字段，则表示调用 API 接口失败。例如：
```json
{
  "code": 100101,
  "message": "Database error"
}
```
上述返回中 `code` 表示错误码，`message` 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500(Internal Server Error)。
## 错误码列表
IAM 系统支持的错误码列表如下：
| Identifier | Code | HTTP Code | Description |
| ---------- | ---- | --------- | ----------- |
| ErrArticleNotFound | 110101 | 400 | Article not found |
| ErrArticleAlreadyExist | 110102 | 400 | Article already exists |
| ErrArticleTagNotFound | 110201 | 400 | ArticleTag not found |
| ErrArticleTagAlreadyExist | 110202 | 400 | ArticleTag already exists |
| ErrAuthNotFound | 110301 | 400 | Auth not found |
| ErrAuthAlreadyExist | 110302 | 400 | Auth already exists |
| Success | 100000 | 200 | OK |
| ServerError | 100001 | 500 | Internal error |
| InvalidParams | 100002 | 400 | Parameters error |
| TooManyRequests | 100003 | 400 | Too many requests |
| UnauthorizedAuthNotExist | 100101 | 400 | Authorization failed: can't find AppKey and AppSecret |
| UnauthorizedTokenError | 100102 | 400 | Authorization failed: token error |
| UnauthorizedTokenTimeout | 100103 | 400 | Authorization failed: token timeout |
| UnauthorizedTokenGenerate | 100104 | 400 | Authorization failed: failed generating a token |
| ErrTagNotFound | 110001 | 400 | Tag not found |
| ErrTagAlreadyExist | 110002 | 400 | Tag already exists |
| ErrUploadFile | 120001 | 400 | Failed uploading files |


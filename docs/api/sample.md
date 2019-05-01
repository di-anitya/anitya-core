# 認証

## トークンの取得
- [`POST`] /v1/auth/tokens

ユーザ名とパスワードからトークンを払い出す。

### リクエスト
#### パラメータ
| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| nocatalog (Optional) | query | string | (Since v3.1) The authentication response excludes the service catalog. By default, the response includes the service catalog. |
| domain | body | object | A domain object |
| name (Optional) | body | string | The user name. Required if you do not specify the ID of the user. If you specify the user name, you must also specify the domain, by ID or name. |
| auth | body | object | An auth object. |
| user | body | object | A user object. |
| password | body | object | The password object, contains the authentication information. |
| id (Optional) | body | string | The ID of the user. Required if you do not specify the user name. |
| identity | body | object | An identity object. |
| methods | body | array | The authentication method. For password authentication, specify password. |

#### 例
```
{
    "auth": {
        "identity": {
            "methods": [
                "password"
            ],
            "password": {
                "user": {
                    "name": "admin",
                    "domain": {
                        "name": "Default"
                    },
                    "password": "devstacker"
                }
            }
        }
    }
}
```
### 応答
#### パラメータ

| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| X-Subject-Token | header | string | The authentication token. An authentication response returns the token ID in this header rather than in the response body. |
| domain | body | object | A domain object |
| methods | body | array | The authentication method. For password authentication, specify password. |
| expires_at | body | string | The date and time when the token expires.<br>The date and time stamp format is ISO 8601:<br>```CCYY-MM-DDThh:mm:ss.sssZ```<br>For example, `2015-08-27T09:49:58.000000Z`.<br>A `null` value indicates that the token never expires. |
| token | body | object | A `token` object. |
| user | body | object | A `user` object. |
| audit_ids | body | array | A list of one or two audit IDs. An audit ID is a unique, randomly generated, URL-safe string that you can use to track a token. The first audit ID is the current audit ID for the token. The second audit ID is present for only re-scoped tokens and is the audit ID from the token before it was re-scoped. A re- scoped token is one that was exchanged for another token of the same or different scope. You can use these audit IDs to track the use of a token or chain of tokens across multiple requests and endpoints without exposing the token ID to non-privileged users. |
| issued_at | body | string | The date and time when the token was issued.<br>The date and time stamp format is ISO 8601:<br>```CCYY-MM-DDThh:mm:ss.sssZ```<br>For example, `2015-08-27T09:49:58.000000Z`. |
| id (Optional) | body | string | The ID of the user. Required if you do not specify the user name. |
| name (Optional) | body | string | The user name. Required if you do not specify the ID of the user. If you specify the user name, you must also specify the domain, by ID or name. |

#### ステータスコード
##### 成功時
| コード | 事由 |
|:--|:--|
| 201 - Created | Resource was created and is ready to use. |

##### 失敗時
| コード | 事由 |
|:--|:--|
| 400 - Bad Request | Some content in the request was invalid. |
| 401 - Unauthorized | User must authenticate before making a request. |
| 403 - Forbidden | Policy does not allow current user to do this operation. |
| 404 - Not Found | The requested resource could not be found. |

#### 例
```
{
    "token": {
        "methods": [
            "password"
        ],
        "expires_at": "2015-11-06T15:32:17.893769Z",
        "user": {
            "domain": {
                "id": "default",
                "name": "Default"
            },
            "id": "423f19a4ac1e4f48bbb4180756e6eb6c",
            "name": "admin",
            "password_expires_at": null
        },
        "audit_ids": [
            "ZzZwkUflQfygX7pdYDBCQQ"
        ],
        "issued_at": "2015-11-06T14:32:17.893797Z"
    }
}
```
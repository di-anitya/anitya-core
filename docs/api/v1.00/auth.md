# 認証 (Auth)
## 《パスワード認証》トークンの発行
- [`POST`] /api/v1/auth/tokens

ユーザ名とパスワードからトークンを払い出す。

### リクエスト
#### パラメータ
| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| user_name | body | string | トークンを払い出すユーザ名。 |
| password | body | string | トークンを払い出すパスワード。 |
| project_id | body | string | トークンを払い出すプロジェクトID。 |

#### 例
```
{
    "user_name": "taro123", 
    "password": "password", 
    "project_id": "435aea53-37cb-4ce3-b76f-ac4674c80f3c"
}
```

### 応答
#### パラメータ

| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| auth | body | object | `auth` オブジェクト。 |
| &nbsp;&nbsp;&nbsp;&nbsp;UserID | body | string | ユーザID。 |
| &nbsp;&nbsp;&nbsp;&nbsp;ProjectID | body | string | プロジェクトID。 |
| &nbsp;&nbsp;&nbsp;&nbsp;token | body | string | トークン文字列。<br>トークン認証APIをリクエストする際にヘッダー部に `Authorization: "Bearer {token}"` を指定する。 |
| status | body | boolean | 処理結果の判定。 |
| message | body | string | 処理結果に関するメッセージ |

#### ステータスコード
##### 成功時
| コード | 事由 |
|:--|:--|
| 200 OK | 正常に処理が完了。 |

##### 失敗時
| コード | 事由 |
|:--|:--|
| 400 Bad Request | リクエスト方法に誤りがある。 |

#### 例
```
{
    "auth": {
        "UserID": "21fef860-5e1b-4860-9efc-02f7864814be",
        "ProjectID": "435aea53-37cb-4ce3-b76f-ac4674c80f3c",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIyMWZlZjg2MC01ZTFiLTQ4NjAtOWVmYy0wMmY3ODY0ODE0YmUiLCJQcm9qZWN0SUQiOiI0MzVhZWE1My0zN2NiLTRjZTMtYjc2Zi1hYzQ2NzRjODBmM2MiLCJ0b2tlbiI6IiJ9.qXwU4GYGs8NO06IYulC-y68Z1gKjbGjhg12jQuGTgtM"
    },
    "message": "success",
    "status": true
}
```

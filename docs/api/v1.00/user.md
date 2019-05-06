# ユーザ (User)
## 《認証なし》ユーザの作成
- [`POST`] /api/v1/user

新規ユーザを作成する。

### リクエスト
#### パラメータ
| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| user_name | body | string | （必須）ユーザ名。 |
| email | body | string | （必須）メールアドレス。 |
| password | body | string | （必須）パスワード。 |
| project_id | body | string | （必須）プロジェクトID。 |
| role_id | body | string | （必須）ロールID。 |
| first_name | body | string | （任意）ユーザフルネームの名前。 |
| last_name | body | string | （任意）ユーザフルネームの苗字。 |

#### 例
```
{
	"user_name": "taro123",
	"first_name":"Taro", 
	"last_name":"Yamada",
	"email": "hanako@sample.com", 
	"password": "password", 
	"project_id": "435aea53-37cb-4ce3-b76f-ac4674c80f3c", 
	"role_id": "71827858-190b-4d8f-94af-05793c2dab7f"
}
```

### 応答
#### パラメータ

| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| user | body | object | `user` オブジェクト。 |
| &nbsp;&nbsp;&nbsp;&nbsp;id | body | string | ユーザID |
| &nbsp;&nbsp;&nbsp;&nbsp;user_name | body | string | ユーザ名。 |
| &nbsp;&nbsp;&nbsp;&nbsp;email | body | string | メールアドレス。 |
| &nbsp;&nbsp;&nbsp;&nbsp;password | body | string | 暗号化されたパスワード。 |
| &nbsp;&nbsp;&nbsp;&nbsp;project_id | body | string | プロジェクトID。 |
| &nbsp;&nbsp;&nbsp;&nbsp;role_id | body | string | ロールID。 |
| &nbsp;&nbsp;&nbsp;&nbsp;first_name | body | string | ユーザフルネームの名前。 |
| &nbsp;&nbsp;&nbsp;&nbsp;last_name | body | string | ユーザフルネームの苗字。 |
| &nbsp;&nbsp;&nbsp;&nbsp;created_at | body | datetime | 作成日時。 |
| &nbsp;&nbsp;&nbsp;&nbsp;updated_at | body | datetime | 更新日時。 |
| status | body | boolean | 処理結果の判定。 |
| message | body | string | 処理結果に関するメッセージ |

#### ステータスコード
##### 成功時
| コード | 事由 |
|:--|:--|
| 200 OK | 正常に処理が完了。 |

##### 失敗時
\* TBD.

#### 例
```
{
    "message": "success",
    "status": true,
    "user": {
        "id": "21fef860-5e1b-4860-9efc-02f7864814be",
        "created_at": "2019-05-06T16:49:00.728456+09:00",
        "update_at": "2019-05-06T16:49:00.728456+09:00",
        "project_id": "435aea53-37cb-4ce3-b76f-ac4674c80f3c",
        "role_id": "71827858-190b-4d8f-94af-05793c2dab7f",
        "user_name": "taro123",
        "first_name": "Taro",
        "last_name": "Yamada",
        "email": "hanako@sample.com",
        "password": "$2a$10$tGR2ar1E2xktPn7TR9H6/.FY5bbSEaKDiWTwl5PQDQvoD7zaJBeuu"
    }
}
```

---
## 《JWT認証》ユーザ一覧の表示
- [`GET`] /api/v1/user

ユーザの属しているプロジェクトに紐づくユーザの一覧を表示する。

### リクエスト
#### パラメータ
| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| Authorization | header | string | （必須）トークン文字列。`Bearer {token}` の形式で指定する。 |

### 応答
#### パラメータ

| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| users | body | array | `users` 配列オブジェクト。 |
| &nbsp;&nbsp;&nbsp;&nbsp;id | body | string | ユーザID |
| &nbsp;&nbsp;&nbsp;&nbsp;user_name | body | string | ユーザ名。 |
| &nbsp;&nbsp;&nbsp;&nbsp;email | body | string | メールアドレス。 |
| &nbsp;&nbsp;&nbsp;&nbsp;password | body | string | 暗号化されたパスワード。 |
| &nbsp;&nbsp;&nbsp;&nbsp;project_id | body | string | プロジェクトID。 |
| &nbsp;&nbsp;&nbsp;&nbsp;role_id | body | string | ロールID。 |
| &nbsp;&nbsp;&nbsp;&nbsp;first_name | body | string | ユーザフルネームの名前。 |
| &nbsp;&nbsp;&nbsp;&nbsp;last_name | body | string | ユーザフルネームの苗字。 |
| &nbsp;&nbsp;&nbsp;&nbsp;created_at | body | datetime | 作成日時。 |
| &nbsp;&nbsp;&nbsp;&nbsp;updated_at | body | datetime | 更新日時。 |
| status | body | boolean | 処理結果の判定。 |
| message | body | string | 処理結果に関するメッセージ |

#### ステータスコード
##### 成功時
| コード | 事由 |
|:--|:--|
| 200 OK | 正常に処理が完了。 |

##### 失敗時
\* TBD.

#### 例
```
{
    "message": "success",
    "status": true,
    "users": [
        {
            "id": "21fef860-5e1b-4860-9efc-02f7864814be",
            "created_at": "2019-05-06T07:49:00Z",
            "update_at": "2019-05-06T07:49:00Z",
            "project_id": "435aea53-37cb-4ce3-b76f-ac4674c80f3c",
            "role_id": "71827858-190b-4d8f-94af-05793c2dab7f",
            "user_name": "taro123",
            "first_name": "Taro",
            "last_name": "Yamada",
            "email": "hanako@sample.com",
            "password": "$2a$10$tGR2ar1E2xktPn7TR9H6/.FY5bbSEaKDiWTwl5PQDQvoD7zaJBeuu"
        }
    ]
}
```

---
## 《JWT認証》ユーザ詳細情報の表示
- [`GET`] /api/v1/user/`{user_id}`

ユーザの属しているプロジェクトに紐づくユーザの詳細情報を表示する。

### リクエスト
#### パラメータ
| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| Authorization | header | string | （必須）トークン文字列。`Bearer {token}` の形式で指定する。 |

### 応答
#### パラメータ

| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| user | body | object | `user` オブジェクト。 |
| &nbsp;&nbsp;&nbsp;&nbsp;id | body | string | ユーザID |
| &nbsp;&nbsp;&nbsp;&nbsp;user_name | body | string | ユーザ名。 |
| &nbsp;&nbsp;&nbsp;&nbsp;email | body | string | メールアドレス。 |
| &nbsp;&nbsp;&nbsp;&nbsp;password | body | string | 暗号化されたパスワード。 |
| &nbsp;&nbsp;&nbsp;&nbsp;project_id | body | string | プロジェクトID。 |
| &nbsp;&nbsp;&nbsp;&nbsp;role_id | body | string | ロールID。 |
| &nbsp;&nbsp;&nbsp;&nbsp;first_name | body | string | ユーザフルネームの名前。 |
| &nbsp;&nbsp;&nbsp;&nbsp;last_name | body | string | ユーザフルネームの苗字。 |
| &nbsp;&nbsp;&nbsp;&nbsp;created_at | body | datetime | 作成日時。 |
| &nbsp;&nbsp;&nbsp;&nbsp;updated_at | body | datetime | 更新日時。 |
| status | body | boolean | 処理結果の判定。 |
| message | body | string | 処理結果に関するメッセージ |

#### ステータスコード
##### 成功時
| コード | 事由 |
|:--|:--|
| 200 OK | 正常に処理が完了。 |

##### 失敗時
\* TBD.

#### 例
```
{
    "message": "success",
    "status": true,
    "user": {
        "id": "21fef860-5e1b-4860-9efc-02f7864814be",
        "created_at": "2019-05-06T07:49:00Z",
        "update_at": "2019-05-06T07:49:00Z",
        "project_id": "435aea53-37cb-4ce3-b76f-ac4674c80f3c",
        "role_id": "71827858-190b-4d8f-94af-05793c2dab7f",
        "user_name": "taro123",
        "first_name": "Taro",
        "last_name": "Yamada",
        "email": "hanako@sample.com",
        "password": "$2a$10$tGR2ar1E2xktPn7TR9H6/.FY5bbSEaKDiWTwl5PQDQvoD7zaJBeuu"
    }
}
```

---
## 《JWT認証》ユーザ情報の更新
- [`PATCH`] /api/v1/user/`{user_id}`

ユーザの属しているプロジェクトに紐づくユーザの情報を更新する。

### リクエスト
#### パラメータ
| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| Authorization | header | string | （必須）トークン文字列。`Bearer {token}` の形式で指定する。 |
| user_name | body | string | （任意）ユーザ名。 |
| email | body | string | （任意）メールアドレス。 |
| password | body | string | （任意）パスワード。 |
| project_id | body | string | （任意）プロジェクトID。 |
| role_id | body | string | （任意）ロールID。 |
| first_name | body | string | （任意）ユーザフルネームの名前。 |
| last_name | body | string | （任意）ユーザフルネームの苗字。 |

#### 例
```
{
	"email": "abc@sample.com"
}
```

### 応答
#### パラメータ

| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| user | body | object | `user` オブジェクト。 |
| &nbsp;&nbsp;&nbsp;&nbsp;id | body | string | ユーザID |
| &nbsp;&nbsp;&nbsp;&nbsp;user_name | body | string | ユーザ名。 |
| &nbsp;&nbsp;&nbsp;&nbsp;email | body | string | メールアドレス。 |
| &nbsp;&nbsp;&nbsp;&nbsp;password | body | string | 暗号化されたパスワード。 |
| &nbsp;&nbsp;&nbsp;&nbsp;project_id | body | string | プロジェクトID。 |
| &nbsp;&nbsp;&nbsp;&nbsp;role_id | body | string | ロールID。 |
| &nbsp;&nbsp;&nbsp;&nbsp;first_name | body | string | ユーザフルネームの名前。 |
| &nbsp;&nbsp;&nbsp;&nbsp;last_name | body | string | ユーザフルネームの苗字。 |
| &nbsp;&nbsp;&nbsp;&nbsp;created_at | body | datetime | 作成日時。 |
| &nbsp;&nbsp;&nbsp;&nbsp;updated_at | body | datetime | 更新日時。 |
| status | body | boolean | 処理結果の判定。 |
| message | body | string | 処理結果に関するメッセージ |

#### ステータスコード
##### 成功時
| コード | 事由 |
|:--|:--|
| 200 OK | 正常に処理が完了。 |

##### 失敗時
\* TBD.

#### 例
```
{
    "message": "success",
    "status": true,
    "user": {
        "id": "21fef860-5e1b-4860-9efc-02f7864814be",
        "created_at": "2019-05-06T07:49:00Z",
        "update_at": "2019-05-06T07:49:00Z",
        "project_id": "435aea53-37cb-4ce3-b76f-ac4674c80f3c",
        "role_id": "71827858-190b-4d8f-94af-05793c2dab7f",
        "user_name": "taro123",
        "first_name": "Taro",
        "last_name": "Yamada",
        "email": "abc@sample.com",
        "password": "$2a$10$tGR2ar1E2xktPn7TR9H6/.FY5bbSEaKDiWTwl5PQDQvoD7zaJBeuu"
    }
}
```

---
## 《JWT認証》ユーザの削除
- [`DELETE`] /api/v1/user/`{user_id}`

ユーザの属しているプロジェクトに紐づくユーザを削除する。

### リクエスト
#### パラメータ

| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| Authorization | header | string | （必須）トークン文字列。`Bearer {token}` の形式で指定する。 |

### 応答
#### パラメータ

| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| status | body | boolean | 処理結果の判定。 |
| message | body | string | 処理結果に関するメッセージ |

#### ステータスコード
##### 成功時
| コード | 事由 |
|:--|:--|
| 200 OK | 正常に処理が完了。 |

##### 失敗時
\* TBD.

#### 例
```
| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| Authorization | header | string | （必須）トークン文字列。`Bearer {token}` の形式で指定する。 |
```

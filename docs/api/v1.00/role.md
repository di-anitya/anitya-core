# ロール (Role)
## 《認証なし》ロールの作成
- [`POST`] /api/v1/role

新規ロールを作成する。

### リクエスト
#### パラメータ
| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| role_name | body | string | （必須）ロール名。 |
| project_id | body | string | （必須）プロジェクトID。 |
| is_admin | body | string | （任意）管理者フラグ。デフォルト値は`FALSE` |

#### 例
```
{
    "role_name": "admin",
    "project_id": "435aea53-37cb-4ce3-b76f-ac4674c80f3c",
    "is_admin": true
}
```

### 応答
#### パラメータ

| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| role | body | object | `role` オブジェクト。 |
| &nbsp;&nbsp;&nbsp;&nbsp;id | body | string | ロールID |
| &nbsp;&nbsp;&nbsp;&nbsp;role_name | body | string | ロール名。 |
| &nbsp;&nbsp;&nbsp;&nbsp;project_id | body | string | プロジェクトID。 |
| &nbsp;&nbsp;&nbsp;&nbsp;is_admin | body | string | 管理者フラグ。 |
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
    "role": {
        "id": "71827858-190b-4d8f-94af-05793c2dab7f",
        "created_at": "2019-05-06T16:42:58.155511+09:00",
        "update_at": "2019-05-06T16:42:58.155511+09:00",
        "project_id": "435aea53-37cb-4ce3-b76f-ac4674c80f3c",
        "role_name": "admin",
        "is_admin": true
    },
    "status": true
}
```

---
## 《JWT認証》ロール一覧の表示
- [`GET`] /api/v1/role

ユーザの属しているプロジェクトのロールの一覧を表示する。

### リクエスト
#### パラメータ
| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| Authorization | header | string | （必須）トークン文字列。`Bearer {token}` の形式で指定する。 |

### 応答
#### パラメータ
\* TBD.

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
    "roles": [
        {
            "id": "71827858-190b-4d8f-94af-05793c2dab7f",
            "created_at": "2019-05-06T07:42:58Z",
            "update_at": "2019-05-06T07:42:58Z",
            "project_id": "435aea53-37cb-4ce3-b76f-ac4674c80f3c",
            "role_name": "admin",
            "is_admin": true
        }
    ],
    "status": true
}
```

---
## 《JWT認証》ロール詳細情報の表示
- [`GET`] /api/v1/role/`{role_id}`

ユーザの属しているプロジェクトのロールの詳細情報を表示する。

### リクエスト
#### パラメータ
| パラメータ名 | 設定位置 | 型 | 説明 |
|:--|:--|:--|:--|
| Authorization | header | string | （必須）トークン文字列。`Bearer {token}` の形式で指定する。 |

### 応答
#### パラメータ
\* TBD.

#### ステータスコード
##### 成功時
| コード | 事由 |
|:--|:--|
| 200 OK | 正常に処理が完了。 |

##### 失敗時
\* TBD.

#### 例
```
* TBD.
```

---
## 《JWT認証》ロール情報の更新
- [`PATCH`] /api/v1/role/`{role_id}`

ユーザの属しているプロジェクトのロール情報を更新する。

### リクエスト
#### パラメータ
\* TBD.

#### 例
```
* TBD.
```

### 応答
#### パラメータ
\* TBD.

#### ステータスコード
##### 成功時
| コード | 事由 |
|:--|:--|
| 200 OK | 正常に処理が完了。 |

##### 失敗時
\* TBD.

#### 例
```
* TBD.
```

---
## 《JWT認証》ロールの削除
- [`DELETE`] /api/v1/role/`{role_id}`

ユーザの属しているプロジェクトのロールを削除する。

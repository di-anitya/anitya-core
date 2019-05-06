
# [プロジェクト (Project)](project.md)
| 項目 | API |
|:--|:--|
| 《認証なし》プロジェクトの作成 | [`POST`] /api/v1/project |
| 《JWT認証》プロジェクト一覧の表示 | [`GET`] /api/v1/project |
| 《JWT認証》プロジェクト詳細情報の表示 | [`GET`] /api/v1/project/`{project_id}` |
| 《JWT認証》プロジェクト情報の更新 | [`PATCH`] /api/v1/project/`{project_id}` |
| 《JWT認証》プロジェクトの削除 | [`DELETE`] /api/v1/project/`{project_id}` |

# [ロール (Role)](role.md)
| 項目 | API |
|:--|:--|
| 《認証なし》ロールの作成 | [`POST`] /api/v1/role |
| 《JWT認証》ロール一覧の表示 | [`GET`] /api/v1/role |
| 《JWT認証》ロール詳細情報の表示 | [`GET`] /api/v1/role/`{role_id}` |
| 《JWT認証》ロール情報の更新 | [`PATCH`] /api/v1/role/`{role_id}` |
| 《JWT認証》ロールの削除 | [`DELETE`] /api/v1/role/`{role_id}` |

# [ユーザ (User)](user.md)
| 項目 | API |
|:--|:--|
| 《認証なし》ユーザの作成 | [`POST`] /api/v1/user |
| 《JWT認証》ユーザ一覧の表示 | [`GET`] /api/v1/user |
| 《JWT認証》ユーザ詳細情報の表示 | [`GET`] /api/v1/user/`{user_id}` |
| 《JWT認証》ユーザ情報の更新 | [`PATCH`] /api/v1/user/`{user_id}` |
| 《JWT認証》ユーザの削除 | [`DELETE`] /api/v1/user/`{user_id}` |

# [認証 (Auth)](auth.md)
| 項目 | API |
|:--|:--|
| 《パスワード認証》トークンの発行 | [`POST`] /api/v1/auth/tokens |

# プローブポイント (ProbePoint)
| 項目 | API |
|:--|:--|
| 《JWT認証》プローブポイントの作成 | [`POST`] /api/v1/probepoint |
| 《JWT認証》プローブポイント一覧の表示 | [`GET`] /api/v1/probepoint |
| 《JWT認証》プローブポイントの詳細表示 | [`GET`] /api/v1/probepoint/`{probepoint_id}` |
| 《JWT認証》プローブポイントの更新 | [`PATCH`] /api/v1/probepoint/`{probepoint_id}` |
| 《JWT認証》プローブポイントの削除 | [`DELETE`] /api/v1/probepoint/`{probepoint_id}` |

# [監視](monitoring.md)
| 項目 | API |
|:--|:--|
| 《JWT認証》HTTP(S) 監視の作成 | [`POST`] /api/v1/monitoring/http |
| 《JWT認証》HTTP(S) 監視一覧の表示 | [`GET`] /api/v1/monitoring/http |
| 《JWT認証》HTTP(S) 監視の詳細表示 | [`GET`] /api/v1/monitoring/http/`{http_id}` |
| 《JWT認証》HTTP(S) 監視の更新 | [`PATCH`] /api/v1/monitoring/http/`{http_id}` |
| 《JWT認証》HTTP(S) 監視の削除 | [`DELETE`] /api/v1/monitoring/http/`{http_id}` |
| 《JWT認証》DNS 監視の作成 | [`POST`] /api/v1/monitoring/dns |
| 《JWT認証》DNS 監視一覧の表示 | [`GET`] /api/v1/monitoring/dns |
| 《JWT認証》DNS 監視の詳細表示 | [`GET`] /api/v1/monitoring/dns/`{dns_id}` |
| 《JWT認証》DNS 監視の更新 | [`PATCH`] /api/v1/monitoring/dns/`{dns_id}` |
| 《JWT認証》DNS 監視の削除 | [`DELETE`] /api/v1/monitoring/dns/`{dns_id}` |

# [アラート通知履歴 (Notification)](notification.md)
| 項目 | API |
|:--|:--|
| 《JWT認証》アラート通知履歴一覧の表示 | [`GET`] /api/v1/notification/history |
| 《JWT認証》アラート通知履歴の詳細表示 | [`GET`] /api/v1/notification/history/`{notification_id}` |
| 《JWT認証》アラート通知履歴の削除 | [`DELETE`] /api/v1/notification/history/`{notification_id}` |
| 《JWT認証》Slack通知設定の作成 | [`POST`] /api/v1/notification/slack |
| 《JWT認証》Slack通知設定一覧の表示 | [`GET`] /api/v1/notification/slack |
| 《JWT認証》Slack通知設定の詳細表示 | [`GET`] /api/v1/notification/slack/`{notification_id}` |
| 《JWT認証》Slack通知設定の更新 | [`PATCH`] /api/v1/notification/slack/`{notification_id}` |
| 《JWT認証》Slack通知設定の削除 | [`DELETE`] /api/v1/notification/slack/`{notification_id}` |

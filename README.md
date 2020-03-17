# sanma
2020 spring, internship

## api
### Users
* ユーザー作成(signup): post, /api/users
* AuthN(signin): post, /api/authentication

### Articles
* 記事作成: post, api/articles
* 全記事取得: get, /api/articles
* 特定ユーザーの記事取得: get, /api/articles?user_id=""
* 特定の記事取得: get, /api/articles/{{ .article_id }}
* 記事更新: put, api/articles/{{ .article_id }}
* 記事削除: delete, api/articles/{{ .article_id }}

### Crawl target urls
* ターゲット追加: get, api/crawl-targets

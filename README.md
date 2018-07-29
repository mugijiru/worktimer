# これは何?

指定した日の Slack の投稿時刻の最初と最後を拾って来て表示するやつ

# インストール

```
$ go get -u github.com/mugijiru/worktimer
```

# 初期設定

https://api.slack.com/custom-integrations/legacy-tokens
からトークンを取得して
環境変数 SLACK_TOKEN に設定してください


# 使い方

`worktimer [date]`

コマンドライン引数がない場合はその日の最初の投稿時間と最後の投稿時間を表示します。
第一引数に日付を指定できます(形式: YYYY/MM/DD)。
日付を指定した場合には、指定日の最初の投稿時間と最後の投稿時間を表示します。

# 出力例

```
2018/07/11      09:54:52        20:10:00
```

# License

MIT License

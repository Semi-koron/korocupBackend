# korocupBackend  
## 開発環境の準備
```
# レポジトリをクローンする
$ git clone git@github.com:Semi-koron/korocupBackend2024.git

# 作業ディレクトリに移動する
$ cd korocupBackend2024

# 依存モジュールをインストールする
$ go mod download

# サーバーを起動
$ docker compose up -d

# 以下のURLを開き、起動しているか確認
http://localhost:8080/

#サーバーの停止
$ docker compose down
```

## 作業手順
```
# ローカル環境を最新にする
$ git switch main
$ git pull

# .envファイルを用意する
# .env.sampleを複製して名前を.envに変更する
# Discordで共有している内容に書き換える

# ブランチを切り替える
# ブランチ名はそのブランチの内容が分かるように
# 例: simesaba/login, Taku/issue3
$ git switch -c <ブランチ名>

# 作業内容をcommit&pushする
# コミットメッセージは内容が分かれば日本語英語問わず好きな形でOK

# そのブランチで作業が終わればmainブランチにプルリクエストを出す
# 原則として一つのブランチにつき一つの機能
# 例: loginブランチでログイン機能が完成したらプルリクを出し、ログアウト機能はlogoutブランチに切り替えてブランチを別にする
```

## gitのコマンドたち  
今いるブランチの確認  
`git branch`  
ブランチの変更  
`git switch (ブランチ名)`  
ブランチの作成と移動  
`git switch -c (ブランチ名)`  
コミットするファイルを追加(ファイル名に . で一括追加)  
`git add (ファイル名)`  
コミット  
`git commit -m "(コミットメッセージ)"`  
コミットしたファイルをpush  
`git push origin (ブランチ名)`  
ローカルのコードを最新にする(mainブランチで)  
`git pull`  

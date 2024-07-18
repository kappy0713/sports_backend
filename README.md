## ExercisePlusのバックエンドリポジトリ

フロントエンドのリポジトリは[こちら](https://github.com/kappy0713/sports_frontend)から

**概要**：日々の運動を記録することで運動習慣の定着をサポートするサービス<br/>　　　自分の記録を分かりやすく可視化

資料は[こちら](https://github.com/kappy0713/sports_backend/blob/main/exerciseplus.pdf)

### デモ動画
https://github.com/user-attachments/assets/ce4184d1-11ea-4785-8122-785d64d46109


### 使用技術
#### フロント
[![My Skills](https://skillicons.dev/icons?i=ts,nextjs,react)](https://skillicons.dev)

#### バック
[![My Skills](https://skillicons.dev/icons?i=go)](https://skillicons.dev)

#### DB・インフラ
[![My Skills](https://skillicons.dev/icons?i=postgres,docker)](https://skillicons.dev)


### 環境構築手順
#### 1. リポジトリをクローン
```
git clone git@github.com:kappy0713/sports_backend.git
```
#### 2.  `.env`ファイルで環境変数の設定
```env
SERVER_URL = YOUR_API_SERVER_URL
POSTGRES_USER = YOUR_POSTGRES_USER
POSTGRES_PASSWORD = YOUR_POSTGRES_PASSWORD
POSTGRES_DB = YOUR_POSTGRESS_DB
POSTGRES_HOST = YOUR_POSTGRESS_HOST
POSTGRES_PORT = YOUR_POSTGRESS_PORT

FRONT_URL = YOUR_FRONT_URL

JWT_SECRET_KEY = YOUR_JWT_SECRET_KEY
```
#### 3. コンテナを起動
```
docker compose up backend --build
```
#### コンテナを停止
```
docker compose down
```
### DBの確認
#### 1. ターミナルからDBコンテナに入る
```
docker compose exec db bash
```
#### 2. DB接続
```
psql -U postgres -d db
```
#### 例)Usersテーブルの中身を確認
```
select * from users;
```

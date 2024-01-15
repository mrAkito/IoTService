# ベースとなるDockerイメージを指定
FROM node:14

# アプリケーションのディレクトリを作成
WORKDIR /usr/src/app

# アプリケーションの依存関係をインストールするファイルをコピー
COPY package*.json ./

# アプリケーションの依存関係をインストール
RUN npm install

# アプリケーションのソースをコピー
COPY . .

# アプリケーションがリッスンするポートを指定
EXPOSE 8080

# コンテナを起動するとき、ターミ
CMD [ "node", "server.js" ]
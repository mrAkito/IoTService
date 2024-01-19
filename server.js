const express = require('express');
const cors = require('cors');

const app = express();

// CORS設定
app.use(cors({
    origin: '*', // 全てのオリジンからのリクエストを許可
    methods: ['GET', 'POST', 'PUT', 'DELETE'], // GET, POST, PUT, DELETEメソッドを許可
  }));
  

// ここからサーバのルーティングを定義する

// 例: GETリクエストのハンドラ
app.get('/', (req, res) => {
    // レスポンスを返す
    // クロスオリジンを全て許可する
    res.header('Access-Control-Allow-Origin', '*');
    // helloworldを返す
    // res.send('hellowo');
    res.sendFile(__dirname + '/init.html');
});

// サーバを起動する
app.listen(8080, () => {
    console.log('Server is running on port 8080');
});

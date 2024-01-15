const express = require('express');
const cors = require('cors');

const app = express();

// CORS設定
app.use(cors());

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
app.listen(3000, () => {
    console.log('Server is running on port 3000');
});

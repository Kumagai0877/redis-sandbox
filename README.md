# redis-sandbox

## 概要

- redisディレクトリでredisの構築
- apiディレクトリでgo-redisでのredisからzrangeで値の取得

## 動作方法

- 起動  
  `make up`

- データ投入

```bash
% docker exec -it redis bash
root@5d6fd6bb6dac:/data# redis-cli 
127.0.0.1:6379> zadd key 100 aaa
(integer) 1
127.0.0.1:6379> zadd key 200 bbb
(integer) 1
```

- apiへリクエスト  
  `curl "localhost:8081/?key=key"`
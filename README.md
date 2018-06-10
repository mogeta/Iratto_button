Press IoT button , send data for Firebase.


* https://firebase.google.com/docs/firestore/quickstart


```
# バイナリ作成
env GOARCH=amd64 GOOS=linux go build myfirestore.go struct.go

# テスト環境で実行
sam local invoke -e event.json -t template.yaml --env-vars sam_env.json

# Lambdaに上げるファイルの作成。secret_keyをアップロードするのを忘れずに
zip myfirestore.zip myfirestore firebase_secret_key.json
```
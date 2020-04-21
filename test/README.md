# test

各エミュレータが正常に機能しているかどうかをテストする．

## 実行

```bash
all_test.sh
```

説明：

1. docker コンテナの削除
2. .data -> .tmpdata に移動 (データの保持を行うため)
3. `SERVICES=""` を指定して，docker コンテナの再立ち上げ
4. `go test` を試行して `runtest.log` にテスト実行時のログを保存．
(各エミュレータごとのテストスクリプトは *_test.go を参照)
5. テスト時の .data を削除して， .tmpdata -> .data に移動 (テスト前のデータの復元)

[runtest.log](./runtest.log) からテスト実行結果を確認する。

## 各テストスクリプト詳細

### datastore_test.go

Test flow:

1. データ作成 
2. データを取得して整合性を確認
3. データを削除して整合性を確認
4. コンテナを再起動してデータが保持されていることを確認

注意点

- ローカルエミュレータのデータは `.data` ディレクトリ内の `local_db.bin` に保存される
- `local_db.bin` の生成とデータの反映には実行からラグがあるため，本スクリプトでは`local_db.bin` が作成するまで待ったのち，30sec おいてから 4.コンテナの再起動を行っている

### firestore_test.go

### pubsub_test.go

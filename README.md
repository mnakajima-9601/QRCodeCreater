# QRCodeCreator
QRコードを作成するプログラムです。

## ローカルでの起動手順
conf.xml内でパラメータ設定できます。

```
<?xml version="1.0" encoding="UTF-8"?>
<List>
	<Out>/home/.../_source/go/QRCode/QRCode/</Out>
	<CsvFile>/home/.../_source/go/QRCode/list.txt</CsvFile>
	<Size>100</Size>
	<TtfFile>・・・.ttf</TtfFile>
</List>
```
|パラメーター|内容|
|----|----|
|Out|出力先のフォルダ名|
|ListFile|バーコードにしたい文字列、入力文字列が保存されたCSVファイル名|
|Size|QRコードのサイズ|
|TtfFile|フォントファイル(実行ファイル化に配置)|

直接引数に指定できます。
```
go run main.go 出力先フォルダ CSVファイルパス
```


## dockerでの起動手順
ビルド後、docker run　で実行します。
※出力先のディレクトリ（②）とCSVファイルが配置されているディレクトリ（③）は
同じディレクトリに作成してください。(①)

```
docker run -it --name コンテナ名 --mount type=bind,src=①,dst=コンテナディレクトリ  qrcodecreater_qrcode コンテナディレクトリ/② コンテナディレクトリ/③/CSVファイル名
```
以下ではカレントディレクトリの想定です。
```
docker-compose build --no-cache
docker run -it --name qrcode_creater --mount type=bind,src=$PWD,dst=/pwd  qrcodecreater_qrcode /pwd/cmd/qrCode/output /pwd/cmd/qrCode/csv/CSVファイル名
```

# QRCodeCreator
QRコードを作成するプログラムです。

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
|Out|出力先のフォルダアドレス|
|ListFile|バーコードにしたい文字列、入力文字列が保存されたCSVファイルのアドレス|
|Size|QRコードのサイズ|
|TtfFile|フォントファイル|
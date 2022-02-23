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
|Out|出力先のフォルダ名|
|ListFile|バーコードにしたい文字列、入力文字列が保存されたCSVファイル名|
|Size|QRコードのサイズ|
|TtfFile|フォントファイル(実行ファイル化に配置)|

出力先はデスクトップ指定になります。
指定フォルダの場合は、以下の部分はデスクトップのパス以下を入力して下さい。
Out
CsvFile
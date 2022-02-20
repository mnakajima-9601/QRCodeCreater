# QRCodeCreator
QRコードを作成するプログラムです。

conf.xml内でパラメータ設定できます。

```
<?xml version="1.0" encoding="UTF-8"?>
<List>
	<Out>/home/.../_source/go/QRCode/QRCode/</Out>
	<ListFile>/home/.../_source/go/QRCode/list.txt</ListFile>
	<Size>100</Size>
</List>
```
|パラメーター|内容|
|----|----|
|Out|出力先のフォルダアドレス|
|ListFile|バーコードにしたい文字列が保存されたファイルのアドレス|
|Size|QRコードのサイズ|


list.txtにQRコードにしたい文字列リストを入れます。
Swag公式

https://github.com/swaggo/swag?tab=readme-ov-file

こちらの記事を参考にいたしました

https://engineering.nifty.co.jp/blog/10123

こちらでswag initの問題（Error parsing type definition　... cannot find type definition: json.Number / ParseComment error in file ... annot find type definition: json.Number）を解決いたしました

https://github.com/swaggo/swag/issues/373

https://github.com/swaggo/swag/issues/817

[使い方の注意点]

・Successアノテーションを使う際は{}内の型に対応するデータ型（構造体など）を記入すること → ParseCommentエラーの原因

・titleやversionアノテーションはmain関数に紐付けて記入する（versionはAPIのバージョンであり、Swaggerのバージョンではない）

・swag initを実行して、「warning: failed to evaluate const mProfCycleWrap at ~ reflect: call of reflect.Value.Len on zero Value」といったエラーが発生したらインストールしたSwaggoのバージョンの問題もあるため、ダウングレードを行なって対応する

https://stackoverflow.com/questions/76582283/swag-init-generates-nothing-but-general-api-information


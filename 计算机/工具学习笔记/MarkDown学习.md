# MarkDown 学习

## 标题
使用`#` 表示标题

## 注意
使用两个以上的空格+换行表示换行`</br>` 一个以上的空行表示一个段落 `</p>`


作用：将回车换为两个空格加回车
作者：吴思澎
链接：https://www.zhihu.com/question/22524345/answer/102385371
来源：知乎
著作权归作者所有，转载请联系作者获得授权。
```js
text=text.replace(/\n+/g,function(ns){
    if(ns.length==1)
        return '  '+ns
    
    return ns
});

text=Markdown.toHtml(text);
```
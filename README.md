# 我的笔记

---
> 记录自己所学所想

**1. 能够输出的才是自己的知识**

**2. 能够做到的才是自己的本领**  

:heart: **不忘初心，方得始终**

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


## markdown测试

### 一、各种流程图
1. 时序图

```seq
Alice->Bob: Hello Bob, how are you?
Note right of Bob: Bob thinks
Bob-->Alice: I am good thanks!
```

2. 流程图

```flow
st=>start: Start
op=>operation: Your Operation
cond=>condition: Yes or No?
e=>end

st->op->cond
cond(yes)->e
cond(no)->op
```

3. 甘特图

```gantt
    title 项目开发流程
    section 项目确定
        需求分析       :a1, 2016-06-22, 3d
        可行性报告     :after a1, 5d
        概念验证       : 5d
    section 项目实施
        概要设计      :2016-07-05, 5d
        详细设计      :2016-07-08, 10d
        编码          :2016-07-15, 10d
        测试          :2016-07-22, 5d
    section 发布验收
        发布: 2d
        验收: 3d
```

4. Mermaid 流程图

```graphLR
    A[Hard edge] -->|Link text| B(Round edge)
    B --> C{Decision}
    C -->|One| D[Result one]
    C -->|Two| E[Result two]
```

5. Mermaid 序列图

```sequence
    Alice->John: Hello John, how are you?
    loop every minute
        John-->Alice: Great!
    end
```
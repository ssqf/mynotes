# 数据库SQL学习

<!-- TOC -->

- [数据库SQL学习](#%E6%95%B0%E6%8D%AE%E5%BA%93sql%E5%AD%A6%E4%B9%A0)
    - [基本概念](#%E5%9F%BA%E6%9C%AC%E6%A6%82%E5%BF%B5)
    - [常见注意点](#%E5%B8%B8%E8%A7%81%E6%B3%A8%E6%84%8F%E7%82%B9)
    - [数据库操作](#%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C)
    - [操作表](#%E6%93%8D%E4%BD%9C%E8%A1%A8)
    - [操作数据](#%E6%93%8D%E4%BD%9C%E6%95%B0%E6%8D%AE)
        - [插入 insert](#%E6%8F%92%E5%85%A5-insert)
        - [删除 delete](#%E5%88%A0%E9%99%A4-delete)
        - [更新](#%E6%9B%B4%E6%96%B0)
        - [查询](#%E6%9F%A5%E8%AF%A2)
    - [分组和聚集](#%E5%88%86%E7%BB%84%E5%92%8C%E8%81%9A%E9%9B%86)
    - [条件逻辑](#%E6%9D%A1%E4%BB%B6%E9%80%BB%E8%BE%91)
    - [连接](#%E8%BF%9E%E6%8E%A5)
    - [视图](#%E8%A7%86%E5%9B%BE)
    - [索引与约束](#%E7%B4%A2%E5%BC%95%E4%B8%8E%E7%BA%A6%E6%9D%9F)
    - [事务](#%E4%BA%8B%E5%8A%A1)
    - [存储过程](#%E5%AD%98%E5%82%A8%E8%BF%87%E7%A8%8B)
    - [触发器](#%E8%A7%A6%E5%8F%91%E5%99%A8)
    - [常用函数](#%E5%B8%B8%E7%94%A8%E5%87%BD%E6%95%B0)

<!-- /TOC -->

## 基本概念

1. SQL (Structured Query Language)
2. CURD (create update read delet)
3. ACD
4. PL/SQL (Procedural Language/SQL)

## 常见注意点

1. 不区分大小写

## 数据库操作

1. 创建数据库
2. 删除数据库
3. 数据库的属性设置
4. psql特殊操作
5. mysql特殊操作

## 操作表

1. 创建表
2. 删除表
3. 数据类型
4. 主键设置
5. 其他约束

## 操作数据

### 插入 insert 

insert into 表名 (列名) values (值)[,(值)],可以有单行或者多行值。如：

```sql
INSERT INTO EMPLOYEES(ID, NAME, AGE, ADDRESS, SALARY)
VALUES
(1, 'Maxsu', 25, '海口市人民大道2880号', 109990.00 ),
(2, 'minsu', 25, '广州中山大道 ', 125000.00 ),
(3, '李洋', 21, '北京市朝阳区', 185000.00),
(4, 'Manisha', 24, 'Mumbai', 65000.00),
(5, 'Larry', 21, 'Paris', 85000.00);
```

### 删除 delete

delete from 表名 where 条件，如果没有where子句，则删除所有数据。
如：`delete from table where name='tako'`

### 更新

update 表名 set 列名=值 where 条件，如果没有where条件子句，则修改所有行数据
如：`update table set name='tako',city='西安' where id=1`

### 查询

1. `select` 选择需要显示的数据列，
    - *表示所有列，如：`selet * from table;`
    - 可以给列起别名，如：`selet name person_name from table;`
    - 指定定值
    - 表达式
    - 内置函数
    - 用户自定义函数调用
    - **distinct** 去除重复行
2. `from` 指定数据来源

数据来源主要指：表和表的连接方式，如，`select * from table;`， `slect t1.name,t2.city from table1 t1 JOIN table2 t2 ON t1.id=t2.id`

3. `where` 主要用于过滤掉不需要的数据
    - 字符相等 `where xxx='yyy'`
    - 
4. `group by` 和 `having`
5. `order by `
    - 升序 `order by xxx asc` 默认升序不需要加asc
    - 降序 `order by xxx desc`
    - 根据表达排序 `order by right(xxx,2)` 根据xxx的右边两位进行排序
    - 根据数字站位符排序，不常用 `order by 2,4` 表示用select的2和4列排序
6. 子查询

## 分组和聚集

## 条件逻辑

## 连接


## 视图

## 索引与约束


## 事务

## 存储过程

## 触发器

## 常用函数

1. SUM
2. COUNT
3.  
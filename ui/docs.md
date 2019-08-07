# 全语种网页正文抽取📑

对文章类网页抽取正文、标题、发布时间、正文图片、作者、语种、地区、关键词等信息，支持全球各语种网站。

### HTTP请求

`POST https://semantics.work/article/api`

### 请求参数

参数 | 描述
--------- | -------
`url` | `字符串`，要进行正文抽取的网页URL

### 状态码说明

status | 说明
--------- | -------
`ok` | 查找成功
`fail` | 失败，此时`message`保存错误信息

### 正文信息说明

字段 | 说明
--- | ---
url|网页URL
title|文章标题
text|文章正文
html|文章正文，HTML版本
publish_date|发布时间
images|正文图片
language|网页语言
location|国家，`ISO 3166-1 alpha-2 Country Codes`
author|作者
tags|关键词


# blog-service
一个博客后端服务

# Introduction
blog-service是一个Web服务，对外提供RESTful API接口，完成Article、Tag和ArticleTag三种REST资源的增删查改

## Article API
|接口名称|接口功能|
|---|---|
|/api/v1/articles [get]|查询article列表|
|/api/v1/articles/:id [get]|查看article|
|/api/v1/articles [post]|新建article|
|/api/v1/articles/:id [put]|修改article|
|/api/v1/articles/:id [delete]|删除article|
|/api/v1/articles/:id [patch]|禁用（启用）article|

## Tag API
|接口名称|接口功能|
|---|---|
|/api/v1/tags [get]|查询tag列表|
|/api/v1/tags/:id [get]|查看tag|
|/api/v1/tags [post]|新建tag|
|/api/v1/tags/:id [put]|修改tag|
|/api/v1/tags/:id [delete]|删除tag|
|/api/v1/tags/:id [patch]|禁用（启用）tag|

## ArticleTag API
|接口名称|接口功能|
|---|---|
|/api/v1/articletags [get]|查询articletag列表|
|/api/v1/articletags/:id [get]|查看articletag|
|/api/v1/articletags [post]|新建articletag|
|/api/v1/articletags/:id [put]|修改articletag|
|/api/v1/articletags/:id [delete]|删除articletag|
|/api/v1/articletags/:id [patch]|禁用（启用）articletag|

# Usage
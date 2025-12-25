# Supply_and_Demand ER图 documentation
## Summary

- [Introduction](#introduction)
- [Database Type](#database-type)
- [Table Structure](#table-structure)
	- [users](#users)
	- [user_infos](#user_infos)
	- [products](#products)
	- [category](#category)
	- [address](#address)
	- [browse_history](#browse_history)
	- [user_follow](#user_follow)
	- [user_like_products](#user_like_products)
	- [product_comments](#product_comments)
	- [user_product_review](#user_product_review)
	- [login_logs](#login_logs)
	- [verification_codes](#verification_codes)
	- [goods_image](#goods_image)
	- [comment](#comment)
	- [comment_reply](#comment_reply)
	- [comment_like](#comment_like)
	- [collection](#collection)
	- [price_watch](#price_watch)
	- [chat_session_product](#chat_session_product)
	- [chat-message](#chat-message)
	- [safe_tip](#safe_tip)
	- [chat_session_notify](#chat_session_notify)
	- [user_homepage](#user_homepage)
	- [user_homepage_visit](#user_homepage_visit)
	- [user_homepage_top](#user_homepage_top)
	- [user_homepage_album](#user_homepage_album)
	- [user_homepage_album_image](#user_homepage_album_image)
	- [order](#order)
- [Relationships](#relationships)
- [Database Diagram](#database-diagram)

## Introduction

> [🔗数据库设计json文件](./instruction/datebase_ER_map.json) 
> 将*`./instruction/database_ER_map.json`文件导入到[drawDB](https://www.drawdb.app/editor)中可查看

**数据库概览图**
![数据库整体图](../image/datebase_note-1.png)

## Database type

- **Database system:** MySQL
## Table structure

### users

| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement | fk_users_id_user_infos |用户ID |
| **status** | ENUM | not null, default: normal |  |用户状态 1 - 正常，0 - 注销，-1 - 违规 |
| **nick_name** | VARCHAR(255) | not null, unique |  |昵称 |
| **password** | VARCHAR(255) | not null |  |密码 |
| **email** | VARCHAR(255) | null, unique |  |邮箱 |
| **scorecard** | FLOAT | not null, default: 0 |  |积分 |
| **credit_score** | FLOAT | not null, default: 100 |  |信用分 | 

#### Enums
##### status

- 1
- 0
- -1


### user_infos
用户信息
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |用户ID |
| **sex** | ENUM | not null |  |性别 m - 男， w - 女 |
| **head_portrait_url** | VARCHAR(1024) | not null, default: 默认头像链接 |  |头像 |
| **registration_date** | DATE | null |  |注册时间 |
| **last_online_date** | DATE | null |  |最后登录时间 |
| **fans_count** | INTEGER | not null, default: 0 |  |粉丝数量 |
| **followers_count** | INTEGER | not null, default: 0 |  |关注数量 |
| **publish_product_count** | INTEGER | not null, default: 0 |  |发布商品数量 |
| **complete_order_count** | INTEGER | not null, default: 0 |  |完成订单数量 |
| **ip_address** |  | not null |  |最近的ip地址 | 

#### Enums
##### sex

- 'm'
- 'w'


### products
商品列表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |商品ID |
| **user_id** | BIGINT | not null | fk_products_user_id_users |用户ID |
| **category_id** | BIGINT | not null | fk_products_category_id_category |分类ID |
| **title** | VARCHAR(255) | not null |  |标题 |
| **intro_text** | TEXT(65535) | not null |  |描述 |
| **price** | DECIMAL | not null |  |价格 |
| **original_price** | DECIMAL | not null |  |原价（打折前） |
| **main_image_url** | VARCHAR(1024) | not null |  |主图 |
| **publish_date** | DATE | not null |  |发布时间 |
| **status** | ENUM | not null, default: normal |  |商品状态 |
| **view_count** | INTEGER | not null, default: 0 |  |浏览量 |
| **want_count** | INTEGER | not null, default: 0 |  |想要的数量 |
| **condition** | ENUM | not null |  |成色 |
| **trade_type** | ENUM | not null |  |交易方式 |
| **free_shipping** |  | not null |  |邮寄方式 |
| **collect_count** |  | not null, default: 收藏的数量 |  | |
| **creat_time** |  | not null |  |商品发布的时间 | 


### category
商品类别
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |分类ID |
| **image** | VARCHAR(1024) | null |  |分类图 |
| **sort_order** | INTEGER | not null, default: 0 |  |排序 |
| **tag_name** |  | not null |  |分类名称 | 


### address
地址
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |地址ID |
| **user_id** | BIGINT | not null | fk_address_user_id_users |用户ID |
| **phone** | VARCHAR(255) | null |  |电话 |
| **is_default** | BOOLEAN | null, default: False |  |是否默认 |
| **province** | VARCHAR(255) | not null |  |省份 |
| **city** | VARCHAR(255) | not null |  |城市 |
| **district** | VARCHAR(255) | not null |  |县/区 |
| **detail_address** | TEXT(65535) | not null |  |详细地址 | 


### browse_history
浏览历史
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |浏览ID |
| **user_id** | BIGINT | not null | fk_browse_history_user_id_users |用户ID |
| **target_id** | BIGINT | not null | fk_browse_history_target_id_products |目标ID |
| **target_type** | ENUM | not null |  |目标类型 |
| **browse_time** | DATE | not null |  |浏览时间 | 

#### Enums
##### target_type

- 'product'
- 'user_info'


### user_follow
用户关注的人
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |用户关注关系ID |
| **follower_id** | BIGINT | not null | fk_user_follow_follower_id_users |关注用户ID |
| **followed_id** | BIGINT | not null | fk_user_follow_follower_id_users |被关注用户ID |
| **follow_time** | DATE | not null |  |关注时间 |
| **status** | TINYINT | not null, default: 1 |  |关注状态 1-关注 0-取消 -1-拉黑 | 


### user_like_products
用户喜欢的商品
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | INTEGER | 🔑 PK, not null, unique, autoincrement |  |喜欢商品关系ID |
| **user_id** | BIGINT | not null | fk_user_like_product_user_id_users |用户ID |
| **product_id** | BIGINT | not null | fk_user_like_product_product_id_products |商品ID |
| **like_date** | DATE | not null |  |喜欢的时间 | 


### product_comments
商品留言
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |留言ID |
| **user_id** | BIGINT | not null | fk_product_comment_user_id_users |用户ID |
| **product_id** | BIGINT | not null | fk_product_comment_product_id_products |商品ID |
| **content** | BIGINT | not null |  |留言 |
| **created_at** | DATE | not null |  |创建时间 |
| **status** | TINYINT | not null, default: 1 |  |留言状态 1 - 正常,0 - 删除,-1 - 违规 | 


### user_product_review
用户评价
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |用户评价ID |
| **reviewer_user_id** | BIGINT | not null | fk_user_product_review_reviewer_user_id_users |买家用户ID |
| **reviewed_user_id** | BIGINT | not null | fk_user_product_review_reviewed_user_id_users |卖家用户ID |
| **product_id** | BIGINT | not null | fk_user_product_review_product_id_products |商品ID |
| **order_id** | BIGINT | not null |  |订单ID |
| **rating** | TINYINT | not null |  |评分 |
| **content** | TEXT(65535) | not null |  |评价文本 |
| **created_at** | DATE | null |  |评价时间 | 


### login_logs
登录日志
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **log_id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |登录日志id |
| **user_id** | BIGINT | null | fk_login_logs_user_id_users |用户id |
| **login_method** |  | null |  |登录方式：手机号/邮箱 |
| **login_result** |  | null |  |登录结果：成功-1 失败-0 |
| **login_account** |  | null |  |登录账号：手机号/邮箱 |
| **client_ip** |  | null |  |客户端IP地址：IPV6/IPV4 | 


### verification_codes
验证码
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **code_id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |验证码id |
| **user_id** | BIGINT | null | fk_verification_codes_user_id_users |用户id |
| **code_type** |  | null |  |验证码类型：登录/注册 |
| **receiver_type** |  | null |  |接收方类型 |
| **receiver_address** |  | null |  |接收方地址 |
| **code_hash** |  | null |  |验证码值 |
| **sent_at** |  | null |  |发送时间 |
| **used_at** |  | null |  |使用时间 |
| **expires_at** |  | null |  |过期时间 | 


### goods_image
商品图片表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | INT | 🔑 PK, not null, unique, autoincrement |  |商品图片ID |
| **goods_id** | BIGINT | null | fk_goods_image_goods_id_products |外键->goods.id  找到照片对应商品 |
| **image_url** | VARCHAR(255) | null |  |图片储存地址 |
| **sort_num** | SMALLINT | null |  |图片展示顺序 | 


### comment
留言表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | INT | 🔑 PK, not null, unique, autoincrement |  |留言表主键ID |
| **goods_id** | BIGINT | null | fk_comment_goods_id_products |外键->goods_id |
| **user_id** | BIGINT | null |  |外键->user_id |
| **is_seller** | BOOLEAN | null |  |0 = 普通用户 / 1 = 卖家 |
| **content** | TEXT(65535) | null |  |留言内容（支持表情包） |
| **publish_time** | DATETIME | null |  |发布时间 |
| **region** | VARCHAR(50) | null |  |留言用户地区 |
| **like_count** | INT | null, default: 0 |  |点赞数 |
| **is_hidden** | BOOLEAN | null |  |0 = 显示 / 1 = 隐藏 | 


### comment_reply
留言回复表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | INT | 🔑 PK, not null, unique, autoincrement |  |留言表回复主键ID |
| **comment_id** | INT | null | fk_comment_reply_comment_id_comment |外键->comment_id |
| **user_id** | BIGINT | null |  |外键->user.id |
| **content** | TEXT(65535) | null |  |回复内容 |
| **reply_time** | DATETIME | null |  |回复时间 |
| **replied_user_id** | BIGINT | null |  |被回复的用户id | 


### comment_like
留言点赞表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | INT | 🔑 PK, not null, unique, autoincrement |  |留言点赞主键ID |
| **comment_id** | BIGINT | null | fk_comment_like_comment_id_comment |外键->comment.id |
| **user_id** | BIGINT | null |  |外键->user.id |
| **like_time** | DATETIME | null |  |点赞时间 | 


### collection
收藏表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | INT | 🔑 PK, not null, unique, autoincrement |  |收藏主键ID |
| **user_id** | BIGINT | null | fk_collection_user_id_users |外键->user.id |
| **goods_id** | BIGINT | null | fk_collection_goods_id_products |外键 ->goods.id |
| **collection_time** | DATETIME | null |  |收藏时间 | 


### price_watch
蹲降价表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | INT | 🔑 PK, not null, unique, autoincrement |  |蹲降价主键ID |
| **user_id** | BIGINT | null |  |外键->user.id |
| **goods_id** | BIGINT | null | fk_price_watch_goods_id_products |外键->goods.id |
| **expect_price** | DECIMAL(10,2) | null |  |期望蹲价 |
| **is_notified** | BOOLEAN | null |  |0 = 未通知 / 1 = 已通知 |
| **create_time** | DATETIME | null |  |创建时间 | 


### chat_session_product
聊天会话表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **chat_id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |会话唯一ID |
| **buyer_id** | BIGINT | null | fk_chat_session_product_buyer_id_users |买家唯一ID |
| **seller_id** | BIGINT | null | fk_chat_session_product_seller_id_users |卖家唯一ID |
| **goods_id** | BIGINT | null | fk_chat_session_product_goods_id_products | |
| **last_msg_time** |  | null, default: 最后一条信息时间 |  | |
| **is_sticky** |  | null |  | |
| **is_mute** |  | null |  | | 


#### Indexes
| Name | Unique | Fields |
|------|--------|--------|
| chat_session_index_0 |  |  |
### chat-message
聊天消息表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **msg_id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |消息唯一ID |
| **chat_id** | BIGINT | null | fk_chat-message_chat_id_chat_session_product |所属会话ID |
| **sender_id** | BIGINT | null |  |发送者ID |
| **type** |  | null |  | |
| **content** |  | null |  |消息内容 |
| **send_time** |  | null |  |发送时间 |
| **is_read** |  | null |  |是否已读（0/1） | 


### safe_tip
安全信息表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **tip_id** | INTEGER | 🔑 PK, not null, unique, autoincrement |  |提示ID |
| **content** |  | null |  |提示内容 |
| **type** |  | null |  |提示类型（防诈骗/交易须知） |
| **session_id** | BIGINT | null | fk_safe_tip_session_id_chat_session_product,fk_safe_tip_session_id_chat_session_notify | | 


### chat_session_notify

| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **chat_id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |会话id |
| **user_id** | BIGINT | null | fk_chat_session_notify_user_id_users |消息接受方id |
| **last_msg_time** |  | null |  |最后消息时间 |
| **is_sticky** |  | null |  | |
| **is_mute** |  | null |  | | 


### user_homepage
个人主页核心配置表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **user_id** | BIGINT | not null, unique | fk_homepage_seller_order,fk_user_homepage_user_id_users |用户ID |
| **id** | BIGINT | 🔑 PK, null, unique, autoincrement | fk_homepage_visit,fk_homepage_top,fk_homepage_id_album,fk_album_id_album_image |唯一标识符 |
| **background_url** | VARCHAR(1024) | null, default: '' |  |背景图片链接 |
| **avatar_banner** | VARCHAR(1024) | null, default: '' |  |头像图片 |
| **signature** | VARCHAR(255) | null, default: '' |  |个性签名 |
| **personal_label** | VARCHAR(255) | null, default: '' |  |个人标签 |
| **is_public** | TINYINT | null, default: 1 |  |主页公开状态 |
| **show_fans** | TINYINT | null, default: 1 |  |粉丝数展开开关 |
| **show_follow** | TINYINT | null, default: 1 |  |关注数展开开关 |
| **show_content** | TINYINT | null, default: 1 |  |发布内容展示开关 |
| **layout_content** | DECIMAL | null |  |主页布局 |
| **update_time** | DATETIME | null |  |配置更新最后时间 |
| **create_time** | DATETIME | null |  |配置创建时间 | 


### user_homepage_visit
个人主页访问统计表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |访客记录唯一ID |
| **homepage_id** | BIGINT | null |  |主表ID |
| **visitor_user_id** | BIGINT | null, default: Null |  |访客用户id |
| **visitor_ip** | VARCHAR(64) | null |  |访客IP地址 |
| **visitor_device** | VARCHAR(255) | null |  |访客设备信息 |
| **visit_time** | DATETIME | null |  |访问时间 |
| **stady_durdation** | INTEGER | null, default: 0 |  |停留时长 |
| **entry_page** | VARCHAR(255) | null |  |入口页面 | 


### user_homepage_top
个人主页内容置顶表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |置顶记录唯一ID |
| **homepage_id** | BIGINT | null |  |主表ID |
| **content_type** | TINYINT | null, default: 1 |  |内容类型 |
| **content_id** | BIGINT | null |  |对应类型内容ID(商品ID,文章ID) |
| **top_sort** | TINYINT | null |  |置顶排序 |
| **top_start_time** | DATETIME | null |  |置顶开始时间 |
| **top_end_time** | DATETIME | null, default: Null |  |置顶结束时间 |
| **create_time** | DATETIME | null |  |记录创建时间 | 


### user_homepage_album
个人主页专属相册表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |相册唯一ID |
| **homepage_id** | BIGINT | null |  |主表ID |
| **album_name** | VARCHAR(64) | null |  |相册名称 |
| **album_cover** | VARCHAR(1024) | null |  |相册封面图 |
| **is_public** | TINYINT | null, default: 1 |  |相册公开状态 |
| **sort_order** | TINYINT | null, default: 1 |  |相册排序 |
| **create_time** | DATETIME | null |  |相册创建时间 |
| **update_time** | DATETIME | null |  |相册最后更新时间 | 


### user_homepage_album_image
个人主页相册图片明细表
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |图片唯一ID |
| **album_id** | BIGINT | null |  |相册ID |
| **image_url** | VARCHAR(1024) | null |  |图片URL |
| **image_desc** | VARCHAR(255) | null |  |图片描述 |
| **sort_num** | TINYINT | null |  |图片排序 |
| **upload_time** | DATETIME | null |  |图片上传时间 | 


### order
订单
| Name        | Type          | Settings                      | References                    | Note                           |
|-------------|---------------|-------------------------------|-------------------------------|--------------------------------|
| **id** | BIGINT | 🔑 PK, not null, unique, autoincrement |  |订单唯一ID |
| **order_no** |  | null |  |订单编号
 |
| **product_id** |  | null |  |关联二手商品表的商品ID |
| **buyer_id** |  | null |  |买家ID |
| **seller_id** | BIGINT | null |  |卖家用户ID |
| **create_time** |  | null |  |订单创建时间 |
| **order_status** |  | null |  |订单状态 |
| **pay_amount** |  | null |  |实付金额 |
| **pay_type** |  | null |  |支付方式 |
| **pay_time** |  | null |  |支付时间 |
| **confirm_receive** |  | null |  |收货确认状态 |
| **receive_time** |  | null |  |买家收货确认时间 |
| **cancel_time** |  | null |  |订单取消时间 |
| **cancel_reason** |  | null |  |取消订单原因 |
| **remark** |  | null |  |订单备注 | 


## Relationships

- **users to user_infos**: one_to_one
- **products to users**: many_to_one
- **products to category**: many_to_one
- **address to users**: many_to_one
- **browse_history to users**: many_to_one
- **browse_history to products**: many_to_one
- **user_follow to users**: many_to_one
- **user_follow to users**: many_to_one
- **user_like_products to users**: many_to_one
- **user_like_products to products**: many_to_one
- **product_comments to users**: many_to_one
- **product_comments to products**: many_to_one
- **user_product_review to users**: many_to_one
- **user_product_review to users**: many_to_one
- **user_product_review to products**: many_to_one
- **verification_codes to users**: many_to_one
- **login_logs to users**: many_to_one
- **comment_reply to comment**: many_to_one
- **comment_like to comment**: many_to_one
- **price_watch to products**: many_to_one
- **goods_image to products**: many_to_one
- **collection to users**: many_to_one
- **collection to products**: many_to_one
- **comment to products**: many_to_one
- **chat_session_product to users**: many_to_one
- **chat_session_product to users**: many_to_one
- **chat_session_product to products**: many_to_one
- **chat-message to chat_session_product**: many_to_one
- **safe_tip to chat_session_product**: many_to_one
- **safe_tip to chat_session_notify**: many_to_one
- **chat_session_notify to users**: many_to_one
- **user_homepage to user_homepage_visit**: one_to_many
- **user_homepage to user_homepage_top**: one_to_many
- **user_homepage to user_homepage_album**: one_to_many
- **user_homepage to user_homepage_album_image**: one_to_many
- **user_homepage to order**: one_to_many
- **user_homepage to users**: one_to_one

## Database Diagram

```mermaid
erDiagram
	users ||--|| user_infos : references
	products }o--|| users : references
	products }o--|| category : references
	address }o--|| users : references
	browse_history }o--|| users : references
	browse_history }o--|| products : references
	user_follow }o--|| users : references
	user_follow }o--|| users : references
	user_like_products }o--|| users : references
	user_like_products }o--|| products : references
	product_comments }o--|| users : references
	product_comments }o--|| products : references
	user_product_review }o--|| users : references
	user_product_review }o--|| users : references
	user_product_review }o--|| products : references
	verification_codes }o--|| users : references
	login_logs }o--|| users : references
	comment_reply }o--|| comment : references
	comment_like }o--|| comment : references
	price_watch }o--|| products : references
	goods_image }o--|| products : references
	collection }o--|| users : references
	collection }o--|| products : references
	comment }o--|| products : references
	chat_session_product }o--|| users : references
	chat_session_product }o--|| users : references
	chat_session_product }o--|| products : references
	chat-message }o--|| chat_session_product : references
	safe_tip }o--|| chat_session_product : references
	safe_tip }o--|| chat_session_notify : references
	chat_session_notify }o--|| users : references
	user_homepage ||--o{ user_homepage_visit : references
	user_homepage ||--o{ user_homepage_top : references
	user_homepage ||--o{ user_homepage_album : references
	user_homepage ||--o{ user_homepage_album_image : references
	user_homepage ||--o{ order : references
	user_homepage ||--|| users : references

	users {
		BIGINT id
		ENUM status
		VARCHAR(255) nick_name
		VARCHAR(255) password
		VARCHAR(255) email
		FLOAT scorecard
		FLOAT credit_score
	}

	user_infos {
		BIGINT id
		ENUM sex
		VARCHAR(1024) head_portrait_url
		DATE registration_date
		DATE last_online_date
		INTEGER fans_count
		INTEGER followers_count
		INTEGER publish_product_count
		INTEGER complete_order_count
		 ip_address
	}

	products {
		BIGINT id
		BIGINT user_id
		BIGINT category_id
		VARCHAR(255) title
		TEXT(65535) intro_text
		DECIMAL price
		DECIMAL original_price
		VARCHAR(1024) main_image_url
		DATE publish_date
		ENUM status
		INTEGER view_count
		INTEGER want_count
		ENUM condition
		ENUM trade_type
		 free_shipping
		 collect_count
		 creat_time
	}

	category {
		BIGINT id
		VARCHAR(1024) image
		INTEGER sort_order
		 tag_name
	}

	address {
		BIGINT id
		BIGINT user_id
		VARCHAR(255) phone
		BOOLEAN is_default
		VARCHAR(255) province
		VARCHAR(255) city
		VARCHAR(255) district
		TEXT(65535) detail_address
	}

	browse_history {
		BIGINT id
		BIGINT user_id
		BIGINT target_id
		ENUM target_type
		DATE browse_time
	}

	user_follow {
		BIGINT id
		BIGINT follower_id
		BIGINT followed_id
		DATE follow_time
		TINYINT status
	}

	user_like_products {
		INTEGER id
		BIGINT user_id
		BIGINT product_id
		DATE like_date
	}

	product_comments {
		BIGINT id
		BIGINT user_id
		BIGINT product_id
		BIGINT content
		DATE created_at
		TINYINT status
	}

	user_product_review {
		BIGINT id
		BIGINT reviewer_user_id
		BIGINT reviewed_user_id
		BIGINT product_id
		BIGINT order_id
		TINYINT rating
		TEXT(65535) content
		DATE created_at
	}

	login_logs {
		BIGINT log_id
		BIGINT user_id
		 login_method
		 login_result
		 login_account
		 client_ip
	}

	verification_codes {
		BIGINT code_id
		BIGINT user_id
		 code_type
		 receiver_type
		 receiver_address
		 code_hash
		 sent_at
		 used_at
		 expires_at
	}

	goods_image {
		INT id
		BIGINT goods_id
		VARCHAR(255) image_url
		SMALLINT sort_num
	}

	comment {
		INT id
		BIGINT goods_id
		BIGINT user_id
		BOOLEAN is_seller
		TEXT(65535) content
		DATETIME publish_time
		VARCHAR(50) region
		INT like_count
		BOOLEAN is_hidden
	}

	comment_reply {
		INT id
		INT comment_id
		BIGINT user_id
		TEXT(65535) content
		DATETIME reply_time
		BIGINT replied_user_id
	}

	comment_like {
		INT id
		BIGINT comment_id
		BIGINT user_id
		DATETIME like_time
	}

	collection {
		INT id
		BIGINT user_id
		BIGINT goods_id
		DATETIME collection_time
	}

	price_watch {
		INT id
		BIGINT user_id
		BIGINT goods_id
		DECIMAL(10,2) expect_price
		BOOLEAN is_notified
		DATETIME create_time
	}

	chat_session_product {
		BIGINT chat_id
		BIGINT buyer_id
		BIGINT seller_id
		BIGINT goods_id
		 last_msg_time
		 is_sticky
		 is_mute
	}

	chat-message {
		BIGINT msg_id
		BIGINT chat_id
		BIGINT sender_id
		 type
		 content
		 send_time
		 is_read
	}

	safe_tip {
		INTEGER tip_id
		 content
		 type
		BIGINT session_id
	}

	chat_session_notify {
		BIGINT chat_id
		BIGINT user_id
		 last_msg_time
		 is_sticky
		 is_mute
	}

	user_homepage {
		BIGINT user_id
		BIGINT id
		VARCHAR(1024) background_url
		VARCHAR(1024) avatar_banner
		VARCHAR(255) signature
		VARCHAR(255) personal_label
		TINYINT is_public
		TINYINT show_fans
		TINYINT show_follow
		TINYINT show_content
		DECIMAL layout_content
		DATETIME update_time
		DATETIME create_time
	}

	user_homepage_visit {
		BIGINT id
		BIGINT homepage_id
		BIGINT visitor_user_id
		VARCHAR(64) visitor_ip
		VARCHAR(255) visitor_device
		DATETIME visit_time
		INTEGER stady_durdation
		VARCHAR(255) entry_page
	}

	user_homepage_top {
		BIGINT id
		BIGINT homepage_id
		TINYINT content_type
		BIGINT content_id
		TINYINT top_sort
		DATETIME top_start_time
		DATETIME top_end_time
		DATETIME create_time
	}

	user_homepage_album {
		BIGINT id
		BIGINT homepage_id
		VARCHAR(64) album_name
		VARCHAR(1024) album_cover
		TINYINT is_public
		TINYINT sort_order
		DATETIME create_time
		DATETIME update_time
	}

	user_homepage_album_image {
		BIGINT id
		BIGINT album_id
		VARCHAR(1024) image_url
		VARCHAR(255) image_desc
		TINYINT sort_num
		DATETIME upload_time
	}

	order {
		BIGINT id
		 order_no
		 product_id
		 buyer_id
		BIGINT seller_id
		 create_time
		 order_status
		 pay_amount
		 pay_type
		 pay_time
		 confirm_receive
		 receive_time
		 cancel_time
		 cancel_reason
		 remark
	}
```
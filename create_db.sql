CREATE TABLE IF NOT EXISTS `owner` (
  `owner_id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `phone` varchar(16) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `salt` char(16) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `password` char(64) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `last_login_dt` bigint(20) DEFAULT NULL,
  `last_logout_dt` bigint(20) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`owner_id`),
  KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci COMMENT='사장님 정보입니다.';

CREATE TABLE IF NOT EXISTS `product` (
  `product_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `owner_id` char(36) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `category` varchar(50) COLLATE utf8mb4_unicode_520_ci NOT NULL COMMENT '카테고리',
  `price` int(11) NOT NULL COMMENT '가격',
  `cost` int(11) NOT NULL COMMENT '원가',
  `name` varchar(200) COLLATE utf8mb4_unicode_520_ci NOT NULL COMMENT '이름',
  `description` varchar(2000) COLLATE utf8mb4_unicode_520_ci NOT NULL COMMENT '설명',
  `barcode` varchar(13) COLLATE utf8mb4_unicode_520_ci NOT NULL COMMENT '바코드, EAN-13',
  `expiration_time` int(11) NOT NULL COMMENT '유통기한(시간 단위)',
  `size` char(5) COLLATE utf8mb4_unicode_520_ci NOT NULL COMMENT '사이즈(small or large)',
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`product_id`) USING BTREE,
  KEY `owner_id` (`owner_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci COMMENT='상품 정보입니다.';

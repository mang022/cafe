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

CREATE FUNCTION `chosung`(
	`str` VARCHAR(200)
) RETURNS varchar(200) CHARSET utf8mb4
    DETERMINISTIC
    COMMENT '초성 추출'
BEGIN
	DECLARE ret VARCHAR(200); 
	DECLARE tmp VARCHAR(1);
	DECLARE cnt INT; 
	DECLARE i INT; 
	DECLARE j INT; 
	DECLARE cur INT;
	
	IF str IS NULL THEN
		RETURN ''; 
	END IF; 
	
	SET cnt = LENGTH(str);
	SET i = 1;
	SET cur = 1;
	
	WHILE cur <= cnt DO 
		SET tmp = SUBSTRING(str,i,1); 
		
		IF tmp RLIKE '^(ㄱ|ㄲ)' OR ( tmp >= '가' AND tmp < '나' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㄱ');
			SET j = 3;
		ELSEIF tmp RLIKE '^ㄴ' OR ( tmp >= '나' AND tmp < '다' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㄴ');
			SET j = 3;
		ELSEIF tmp RLIKE '^(ㄷ|ㄸ)' OR ( tmp >= '다' AND tmp < '라' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㄷ');
			SET j = 3;
		ELSEIF tmp RLIKE '^ㄹ' OR ( tmp >= '라' AND tmp < '마' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㄹ');
			SET j = 3;
		ELSEIF tmp RLIKE '^ㅁ' OR ( tmp >= '마' AND tmp < '바' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㅁ');
			SET j = 3;
		ELSEIF tmp RLIKE '^ㅂ' OR ( tmp >= '바' AND tmp < '사' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㅂ');
			SET j = 3;
		ELSEIF tmp RLIKE '^(ㅅ|ㅆ)' OR ( tmp >= '사' AND tmp < '아' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㅅ');
			SET j = 3;
		ELSEIF tmp RLIKE '^ㅇ' OR ( tmp >= '아' AND tmp < '자' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㅇ');
			SET j = 3;
		ELSEIF tmp RLIKE '^(ㅈ|ㅉ)' OR ( tmp >= '자' AND tmp < '차' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㅈ');
			SET j = 3;
		ELSEIF tmp RLIKE '^ㅊ' OR ( tmp >= '차' AND tmp < '카' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㅊ');
			SET j = 3;
		ELSEIF tmp RLIKE '^ㅋ' OR ( tmp >= '카' AND tmp < '타' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㅋ');
			SET j = 3;
		ELSEIF tmp RLIKE '^ㅌ' OR ( tmp >= '타' AND tmp < '파' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㅌ');
			SET j = 3;
		ELSEIF tmp RLIKE '^ㅍ' OR ( tmp >= '파' AND tmp < '하' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㅍ');
			SET j = 3;
		ELSEIF tmp RLIKE '^ㅎ' OR ( tmp >= '하' AND tmp <= '힣' ) THEN 
			SET ret = CONCAT(IFNULL(ret,''), 'ㅎ');
			SET j = 3;
		ELSE
			SET j = 1;
		END IF;
		
		SET cur=cur+j;
		SET i=i+1; 
	END WHILE;
	
	RETURN ret; 
END;

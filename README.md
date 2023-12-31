
# Cafe
카페 관리를 위한 유저 로그인, 상품 관리를 위한 서비스입니다.

# 실행방법
- `config.json` 파일을 만듭니다.
```
{
	"host": {
		"address": "localhost",	// 원하는 서버의 호스트를 입력합니다.
		"port": 10001			// 원하는 서버의 포트를 입력합니다.
	},
	"db": {
		"host": "localhost",	// DB 주소를 입력합니다.
		"port": 3306,			// DB 포트를 입력합니다.
		"user": "",				// DB 유저이름을 입력합니다.
		"pwd": ""				// DB 비밀번호를 입력합니다.
		},
	"jwt": {
		"secret": "08d90a2e7c49a160e46f14b618c350651754dc6fc0f101ed7b73130f5a1ad170"	// 임의의 256bit 입력합니다.
	}
}
```
- `docker build -t docker.io/mystery1348/cafe:test -f ./Test.Dockerfile .`
주의) 본인의 docker hub 주소를 사용해주시기 바랍니다. 
tag는 아무거나 하셔도 됩니다.
Test.Dockerfile도 원하시면 변경하셔도 됩니다.

- `docker push docker.io/mystery1348/cafe:test`
주의) build에 사용한 docker hub 주소와 tag를 사용해주시기 바랍니다.

- `docker run -d --name cafe-test -p 10001:10001 --network="host" docker.io/mystery1348/cafe:test`
주의) --name 옵션은 본인이 짓고 싶은 이름을 입력하시면 됩니다.
-p 포트는 제가 임의로 10001로 지정했습니다.
--network 옵션은 local DB에 접속하기 위해서 사용하였습니다. 원격 DB를 사용하시면 지워도 됩니다.

# API 리스트
### 회원가입 | POST | /signup
#### request body
```
phone | string | required | ^01([0|1|6|7|8|9])(-)([0-9]{3,4})(-)([0-9]{4})$
password | string | required | 8-128 자리
```  
#### response body
```
200 - 성공
400 - 휴대폰번호, 비밀번호 잘못 입력 / 이미 존재하는 휴대폰번호
503 - 내부 에러
```

### 로그인 | POST | /signin
#### request body
```
phone | string | required | ^01([0|1|6|7|8|9])(-)([0-9]{3,4})(-)([0-9]{4})$
password | string | required | 8-128 자리
```  
#### response body
```
200 - 성공
jwt | string | required | JWT 토큰
400 - 휴대폰번호, 비밀번호 잘못 입력
503 - 내부 에러
```

### 로그아웃 | POST | /owner/:id/signout
#### header
```
Authorization | string | required | Bearer {JWT 토큰}
```
#### request body

#### response body
```
200 - 성공
503 - 내부 에러
```

### 상품등록 | POST | /owner/:id/product
#### header
```
Authorization | string | required | Bearer {JWT 토큰}
```
#### request body
```
category | string | required | 1-50 글자
price | int | required | 0 이상
cost | int | required | 0 이상
name | string | required | 1-200 글자
description | string | required | 1-2000 글자
barcode | string | required | 1-13 글자
expiration_time | int | required | 유통기한(시간 단위)
size | string | required | small or large
```  
#### response body
```
200 - 성공
400 - 상품 정보 잘못 입력
503 - 내부 에러
```

### 상품 정보 수정 | PUT | /owner/:id/product/:pid
#### header
```
Authorization | string | required | Bearer {JWT 토큰}
```
#### request body
```
category | string | optional | 1-50 글자
price | int | optional | 0 이상
cost | int | optional | 0 이상
name | string | optional | 1-200 글자
description | string | optional | 1-2000 글자
barcode | string | optional | 1-13 글자
expiration_time | int | optional | 유통기한(시간 단위)
size | string | optional | small or large
```  
#### response body
```
200 - 성공
400 - 상품 정보 잘못 입력
503 - 내부 에러
```

### 상품 제거 | DELETE | /owner/:id/product/:pid
#### header
```
Authorization | string | required | Bearer {JWT 토큰}
```
#### request body

#### response body
```
200 - 성공
400 - pid 잘못 입력
503 - 내부 에러
```

### 상품 상세 조회 | GET | /owner/:id/product/:pid
#### header
```
Authorization | string | required | Bearer {JWT 토큰}
```
#### request body

#### response body
```
200 - 성공
product | struct | requried |
	id | int | required | 상품 ID
	category | string | required | 카테고리
	price | int | required | 가격
	cost | int | required | 원가
	name | string | required | 이름
	description | string | required | 설명
	barcode | string | required | 바코드 번호
	expiration_time | int | required | 유통기한(시간 단위)
	size | string | required | small or large
400 - pid 잘못 입력
503 - 내부 에러
```

### 상품 목록 조회 | GET | /owner/:id/product/
#### header
```
Authorization | string | required | Bearer {JWT 토큰}
```
#### parameter
```
last_id | int | optional | 현재 조회한 ID 중 가장 작은 ID
keyword | string | optional | 검색어
```
#### request body

#### response body
```
200 - 성공
products | array | required
	id | int | required | 상품 ID
	category | string | required | 카테고리
	price | int | required | 가격
	name | string | required | 이름
400 - last_id 잘못 입력
503 - 내부 에러
```


# DB 설계

기본 문자열은 utf8mb4_unicode_520_ci로 선택했습니다. 이모티콘도 들어갈 것을 고려하였습니다.

### 테이블 owner (사장님)

- owner_id | 사장님 ID입니다. uuid를 사용하였습니다. 유일합니다.

- phone | 전화번호입니다. `-`를 저장하지 않을 것입니다. 로그인할 때 검색을 위하여 index를 추가하였습니다. 유일합니다.

- salt | 비밀번호에 추가될 salt입니다. 알파벳, 숫자 16자리 조합을 이용합니다.

- password | 비밀번호입니다. SHA256 해시를 사용할 것입니다.

- last_login_dt | 최근 로그인 시간입니다. 로그인을 했는지 확인하기 위하여 추가하였습니다.

- last_logout_dt | 최근 로그아웃 시간입니다. 로그아웃을 했는지 확인하기 위하여 추가하였습니다.

- created_at | 생성 시간

- updated_at | 수정 시간

- deleted_at | 삭제 시간


### 테이블 product (상품)

- proudct_id | 상품 ID입니다.

- owner_id | owner 테이블의 사장님 ID입니다. 외래키를 걸지 않았습니다. 개발 초기 단계라고 판단했기 때문에 외래키로 인한 DB 수정의 어려움을 야기시키지 않기를 위함입니다.

- category | 카테고리입니다. 따로 데이터에 대한 언급이 없어서 문자열로 판단하였습니다.

- price | 가격입니다. 따로 언급이 없어서 국내 서비스라고 판단하였습니다. 소수는 고려하지 않았습니다.

- cost | 원가입니다. 마찬가지로 따로 언급이 없어서 국내 서비스라고 판단하였습니다. 소수는 고려하지 않았습니다.

- name | 이름입니다.

- description | 설명입니다.

- barcode | 바코드입니다. 언급이 없어서 EAN-13 바코드라고 가정했습니다.

- expiration_time | 유통기한입니다. 따로 업급이 없어서 시간 단위를 입력하기로 하였습니다.

- size | 사이즈입니다.

- created_at | 생성 시간

- updated_at | 수정 시간

- deleted_at | 삭제 시간

  

# MySQL 5.7 설치 (Ubuntu 20.04)

AWS RDB만 이용하다보니 로컬 환경에 DB가 설치되지 않았다는 것을 알았습니다.

1. Debian package 다운로드
Ubuntu 20.04 APT repository에는 MySQL 8.0 밖에 없기 때문에 5.7 버전이 있는 repository를 설치해야 합니다.
`wget https://dev.mysql.com/get/mysql-apt-config_0.8.12-1_all.deb`
주의) 가장 최근 Debian package는 5.7을 지원해주지 않습니다. (https://dev.mysql.com/get/mysql-apt-config_0.8.29-1_all.deb)
`sudo dpkg -i mysql-apt-config_0.8.12-1_all.deb`
프롬프트가 나오면 Ubuntu Bionic / MySQL Server & Cluster / mysql-5.7 순으로 선택하면 됩니다.
`sudo apt update`
주의) 서명이 올바르지 않다고 나올 수 있습니다. 23년 12월에 public key가 변경되었습니다. 그런 경우에는
`sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys B7B3B788A8D3785C`
를 실행한 후 다시 sudo apt update를 하면 됩니다.
`sudo apt list --all-versions mysql-client`
버전을 확인했을 때 5.7 버전이 있으면 됩니다.

2. MySQL 설치
`sudo apt list --all-versions mysql-client`
버전을 확인해야 합니다.
`sudo apt install -f mysql-client=5.7.42-1ubuntu18.04 mysql-community-server=5.7.42-1ubuntu18.04 mysql-server=5.7.42-1ubuntu18.04`
주의) 순서가 중요합니다.
프롬프트가 나오면 root의 초기 비밀번호를 입력합니다.
`sudo mysql_secure_installation`
보안 모듈을 설치합니다.
`mysql -u root -p`
실행하여 확인합니다.
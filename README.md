# cafe

# DB 설계
기본 문자열은 utf8mb4_unicode_520_ci로 선택했습니다. 이모티콘도 들어갈 것을 고려하였습니다.

- 테이블 owner (사장님)
- owner_id 사장님 ID입니다. uuid를 사용하였습니다. 유일합니다.
- phone 전화번호입니다. `-`를 저장하지 않을 것입니다. 로그인할 때 검색을 위하여 index를 추가하였습니다. 유일합니다.
- salt 비밀번호에 추가될 salt입니다. 알파벳, 숫자 16자리 조합을 이용합니다.
- password 비밀번호입니다. SHA256 해시를 사용할 것입니다.
- last_login_dt 최근 로그인 시간입니다. 로그인을 했는지 확인하기 위하여 추가하였습니다.
- last_logout_dt 최근 로그아웃 시간입니다. 로그아웃을 했는지 확인하기 위하여 추가하였습니다.
- created_at 생성 시간
- updated_at 수정 시간
- deleted_at 삭제 시간

- 테이블 product (상품)
- proudct_id 상품 ID입니다.
- owner_id owner 테이블의 사장님 ID입니다. 외래키를 걸지 않았습니다. 개발 초기 단계라고 판단했기 때문에 외래키로 인한 DB 수정의 어려움을 야기시키지 않기를 위함입니다.
- category 카테고리입니다. 따로 데이터에 대한 언급이 없어서 문자열로 판단하였습니다.
- price 가격입니다. 따로 언급이 없어서 국내 서비스라고 판단하였습니다. 소수는 고려하지 않았습니다.
- cost 원가입니다. 마찬가지로 따로 언급이 없어서 국내 서비스라고 판단하였습니다. 소수는 고려하지 않았습니다.
- name 이름입니다.
- description 설명입니다.
- barcode 바코드입니다. 언급이 없어서 EAN-13 바코드라고 가정했습니다.
- expiration_dt 유통기한입니다. unix-timestamp로 시간을 비교하는 게 낫다고 생각했습니다.
- size 사이즈입니다. 
- created_at 생성 시간
- updated_at 수정 시간
- deleted_at 삭제 시간

# MySQL 5.7 설치 (Ubuntu 20.04)
AWS RDB만 이용하다보니 로컬 환경에 DB가 설치되지 않았다는 것을 알았습니다.

1. Debian package 다운로드
Ubuntu 20.04 APT repository에는 MySQL 8.0 밖에 없기 때문에 5.7 버전이 있는 repository를 설치해야 합니다.

- wget https://dev.mysql.com/get/mysql-apt-config_0.8.12-1_all.deb
주의) 가장 최근 Debian package는 5.7을 지원해주지 않습니다. (https://dev.mysql.com/get/mysql-apt-config_0.8.29-1_all.deb)
- sudo dpkg -i mysql-apt-config_0.8.12-1_all.deb
프롬프트가 나오면 Ubuntu Bionic / MySQL Server & Cluster / mysql-5.7 순으로 선택하면 됩니다.
- sudo apt update
주의) 서명이 올바르지 않다고 나올 수 있습니다. 23년 12월에 public key가 변경되었습니다. 그런 경우에는
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys B7B3B788A8D3785C
를 실행한 후 다시 sudo apt update를 하면 됩니다.
- sudo apt list --all-versions mysql-client
버전을 확인했을 때 5.7 버전이 있으면 됩니다.

2. MySQL 설치

- sudo apt list --all-versions mysql-client
버전을 확인해야 합니다.
- sudo apt install -f mysql-client=5.7.42-1ubuntu18.04 mysql-community-server=5.7.42-1ubuntu18.04 mysql-server=5.7.42-1ubuntu18.04
순서가 중요합니다.
프롬프트가 나오면 root의 초기 비밀번호를 입력합니다.
- sudo mysql_secure_installation
보안 모듈을 설치합니다.
- mysql -u root -p
실행하여 확인합니다.
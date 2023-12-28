# cafe

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
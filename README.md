# I-love-reading
![主页](static/img/home.png)

![分类](static/img/category.png)


# centos mysql安装
1. yum -y install mysql mysql-server mysql-devel
2. rpm -qi mysql-server
3. service mysqld start
4. create user apple@'%' identified by 'apple';
5. GRANT ALL PRIVILEGES ON `%`.* TO 'apple'@'%' IDENTIFIED BY 'apple' WITH GRANT OPTION;
   GRANT ALL PRIVILEGES ON `%`.* TO 'apple'@'localhost' IDENTIFIED BY 'apple' WITH GRANT OPTION;
6. flush privileges;
7. /etc/init.d/iptables stop

# Nginx
1. sudo yum install epel-release
2. yum --enablerepo=epel install nginx
3. /etc/init.d/nginx start

# Nginx ssl
1. 阿里云申请签名（搜索安全）
2. 修改nginx中config.d中的ssl.config
3. 修改.pem和.key文件的位置
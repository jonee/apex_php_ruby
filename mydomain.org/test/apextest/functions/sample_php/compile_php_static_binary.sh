#!/bin/sh

# you should compile php in the same AMI template used by AWS Lambdas eg amzn-ami-hvm-2016.03.3.x86_64-gp2 (ami-a59b49c6)

PWD=pwd

#VERSION="5.6.5"
#VERSION="7.1.4"
VERSION="5.6.29"

yum groupinstall "Development tools"
yum install -y libexif-devel libjpeg-devel gd-devel curl-devel openssl-devel libxml2-devel autoconf

# does not work in ubuntu!
# apt install libexif-dev libjpeg-dev curl openssl libxml2-dev libcurl4-openssl-dev pkg-config libssl-dev libsslcommon2-dev libpng-dev libfreetype6-dev build-essential libssl1.0.0
#sudo ln -s /lib/x86_64-linux-gnu/libssl.so.1.0.0 /usr/lib/libssl.so.10
#sudo ln -s /lib/x86_64-linux-gnu/libcrypto.so.1.0.0 /usr/lib/libcrypto.so.10


cd /tmp

# if download is needed
#wget http://ro1.php.net/get/php-$VERSION.tar.gz/from/this/mirror -O php-$VERSION.tar.gz
rm -rf php-$VERSION/
tar xzf php-$VERSION.tar.gz



# mongo legacy php driver since the mongodb driver cannot be statically compiled https://jira.mongodb.org/browse/PHPC-759
# mongo legacy though does not seem to work with php7
#wget http://pecl.php.net/get/mongo-1.6.14.tar 
## or if pecl is installed ../compiled/bin/pecl download mongo
##tar xzf mongo-*.* # if tgz
tar xf mongo-*.tar
mv mongo-*/ php-$VERSION/ext/mongo/


# would have been the supported mongodb library
#wget http://pecl.php.net/get/mongodb-1.2.9.tar
## or if pecl is installed ../compiled/bin/pecl download mongodb
#tar xf mongodb-*.tar
#mv mongodb-*/ php-$VERSION/ext/mongodb/




cd php-$VERSION
rm configure
./buildconf --force

# removed because aws likely does not have libs -> --with-exif
# we are doomed if openssl would not work since mongo connection from php needs it 
# ./configure --prefix=/tmp/php-$VERSION/compiled/ --enable-shared=no --enable-static=yes --enable-phar --enable-json --disable-all --with-openssl --with-curl --enable-libxml --enable-simplexml --enable-xml --with-mhash --with-gd --enable-exif --enable-mbstring --enable-sockets --with-pear=/tmp/lib/php --enable-mongo

./configure --prefix=/tmp/php-$VERSION/compiled/ --enable-shared=no --enable-static=yes --enable-phar --enable-json --disable-all --with-openssl --with-curl --enable-libxml --enable-simplexml --enable-xml --with-mhash --with-gd --enable-mbstring --enable-sockets --with-pear=/tmp/lib/php --enable-mongo


make
make install 









cd $PWD



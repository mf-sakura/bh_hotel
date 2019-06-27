# Waiting mysql
until mysqladmin ping -uroot -ppassword -h mysql_hotel --silent; do
    echo 'waiting for mysqld to be connectable...' && sleep 3;
done

mysqladmin -uroot -ppassword -h mysql_hotel create bh_hotel

goose -env=compose -path=. up

mysql -f  -h mysql_hotel -uroot -ppassword bh_hotel < testdata/testdata.sql
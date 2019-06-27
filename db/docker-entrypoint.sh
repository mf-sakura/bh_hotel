# Waiting mysql
until mysqladmin ping -u root -p password -h mysql_hotel --silent; do
    echo 'waiting for mysqld to be connectable...' && sleep 3;
done


goose -env compose up

mysql -f  -h mysql_hotel -u root -p password bh_hotel < testdata/testdata.sql
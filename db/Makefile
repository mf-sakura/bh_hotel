prepare-db:
	mysqladmin -u root create bh_hotel
	mysqladmin -u root create bh_hotel_test

down:
	goose down
	goose -env=test down

up:
	goose up
	goose -env=test up
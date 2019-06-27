

-- Hotel
INSERT INTO `hotel` (`id`,`name`) VALUES (1, "箱根旅館");

-- Plan
INSERT INTO `plan` (`id`, `hotel_id`, `description`, `date_unix`, `total`,`available`, `cost` ) VALUES
(1, 1, "激安",1561906800,10,0,1000),
(2, 1,"高級",1561906800,5,4,100000);
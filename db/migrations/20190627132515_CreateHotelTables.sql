
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS `hotel` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL comment 'ホテル名',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `name` (`name` ASC)
)ENGINE = InnoDB
DEFAULT CHARSET=utf8mb4
comment='ホテルのマスター情報';

-- 日帰りのみ扱う(テーブル設計ミス)
CREATE TABLE IF NOT EXISTS `plan` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `hotel_id` INT UNSIGNED NOT NULL,
  `description` MEDIUMTEXT NOT NULL comment 'プラン説明',
  `date_unix` BIGINT UNSIGNED NOT NULL comment 'プランの日程',
  `total` INT UNSIGNED NOT NULL comment 'プランの総数',
  `available` INT UNSIGNED NOT NULL comment 'プランの残り利用数',
  `cost` INT UNSIGNED NOT NULL comment '金額(JPY)',
  PRIMARY KEY (`id`),
  INDEX `hotel_id` (`hotel_id` ASC),
  INDEX `hotel_id_date_unix` (`hotel_id` ASC, `date_unix` ASC),
  FOREIGN KEY (`hotel_id`) REFERENCES `hotel` (`id`) ON DELETE CASCADE
  )ENGINE = InnoDB
DEFAULT CHARSET=utf8mb4
comment='ホテルのプラン情報';

CREATE TABLE IF NOT EXISTS `reservation` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` INT UNSIGNED NOT NULL comment '宿泊者のUserID',
  `plan_id` INT UNSIGNED NOT NULL comment 'PlanのID',
  `sequence_id` VARCHAR(255) NOT NULL comment 'マイクロサービスのシーケンスID',
  PRIMARY KEY (`id`),
  INDEX `plan_id` (`plan_id` ASC),
  INDEX `user_id` (`user_id` ASC),
  INDEX `plan_id_user_id` (`plan_id` ASC, `user_id` ASC),
  UNIQUE INDEX `sequence_id` (`sequence_id` ASC),
  FOREIGN KEY (`plan_id`) REFERENCES `plan` (`id`) ON DELETE CASCADE
)ENGINE = InnoDB
DEFAULT CHARSET=utf8mb4
comment='ホテルの予約情報';



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE `reservation`;
DROP TABLE `plan`;
DROP TABLE `hotel`;

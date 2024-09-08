CREATE TABLE user (
                      id bigint AUTO_INCREMENT,
                      name varchar(128) NOT NULL DEFAULT '' COMMENT 'The username',
                      gender tinyint(4) NOT NULL DEFAULT 0 COMMENT 'The gender',
                      password varchar(128) NOT NULL DEFAULT '' COMMENT 'The password',
                      phone varchar(128) NOT NULL DEFAULT '' COMMENT 'The phone',
                      type tinyint(1) NULL DEFAULT 0 COMMENT 'The userdd type, 0:normal,1:vip, for test golang keyword',
                      create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'userdd table';


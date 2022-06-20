CREATE DATABASE expenses;
USE expenses;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
                         `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                         `name` varchar(50) NOT NULL,
                         `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `users` VALUES
                        ('1','joao', '2020-08-09 10:27:22'),
                        ('2','maria', '2020-08-09 10:27:23'),
                        ('3','pedrinho', '2020-08-09 10:27:24');

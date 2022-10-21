CREATE SCHEMA IF NOT EXISTS `appointments_go`;

DROP TABLE IF EXISTS `appointments_go`.`dentists`;
CREATE TABLE `appointments_go`.`dentists` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `lastname` VARCHAR(50) NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `license` VARCHAR(50) NOT NULL,
  PRIMARY KEY (`id`))
  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
  
DROP TABLE IF EXISTS `appointments_go`.`patients`;
CREATE TABLE `appointments_go`.`patients` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `lastname` VARCHAR(50) NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `identity_document` VARCHAR(50) NOT NULL,
  `entry_date` DATE NOT NULL,
  PRIMARY KEY (`id`))
  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
  
DROP TABLE IF EXISTS `appointments_go`.`appointments`;
CREATE TABLE `appointments_go`.`appointments` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `dentist_id` INT NOT NULL,
  `patient_id` INT NOT NULL,
  `date` DATE NOT NULL,
  `time` VARCHAR(50) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `dentist_id` FOREIGN KEY (`dentist_id`) REFERENCES `dentists` (`id`),
  CONSTRAINT `patient_id` FOREIGN KEY (`patient_id`) REFERENCES `patients` (`id`)
  )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
  
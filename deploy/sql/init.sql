CREATE TABLE `Department` (
    `gorm.Model` varchar(255) NOT NULL
);

CREATE TABLE `Hierarchy`  (
    `gorm.Model` varchar(255) NOT NULL,
    `company` varchar(255) NULL,
    `department` varchar(255) NULL,
    PRIMARY KEY (`gorm.Model`),
    INDEX `organization`(`company`, `department`)
);

CREATE TABLE `Role`  (
    `gorm.Model` varchar(255) NOT NULL,
    `name` varchar(255) NULL,
    PRIMARY KEY (`gorm.Model`)
);

CREATE TABLE `User`  (
    `gorm.Model` varchar(255) NOT NULL,
    `name` varchar(255) NULL,
    `password` varchar(255) NULL,
    `phonenum` int NULL,
    PRIMARY KEY (`gorm.Model`)
);

#
# Copyright (C) 2019 Vdaas.org Vald team ( kpango, kmrmt, rinx )
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
DROP SCHEMA IF EXISTS `vald` ;
CREATE SCHEMA IF NOT EXISTS `vald` DEFAULT CHARACTER SET utf8 ;

USE `vald` ;

CREATE TABLE IF NOT EXISTS `vald`.`meta_vector` (
    `uuid` VARCHAR(255) NOT NULL,
    `vector` TEXT NOT NULL,
    `meta` VARCHAR(1024) NOT NULL,
    PRIMARY KEY (`uuid`),
    UNIQUE INDEX `meta_unique` (`meta` ASC)
)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `vald`.`pod_ip` (
    `uuid` VARCHAR(255) NOT NULL,
    `ip` VARCHAR(64) NOT NULL,
    PRIMARY KEY (`uuid`, `ip`),
    INDEX `ip_index` (`ip` ASC)
)
ENGINE = InnoDB;

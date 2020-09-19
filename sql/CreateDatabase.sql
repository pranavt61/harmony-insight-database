CREATE DATABASE IF NOT EXISTS HarmonyDashboard;

USE HarmonyDashboard;

CREATE TABLE IF NOT EXISTS BlockTransactionCount (
    shard_id      TINYINT UNSIGNED,
    block_height  INT UNSIGNED,
    tx_count      SMALLINT UNSIGNED,

    PRIMARY KEY (shard_id, block_height)
);

CREATE TABLE IF NOT EXISTS UserVisits (
  request_type  VARCHAR(7),
  user_ip       VARCHAR(21),
  time          INT(11) UNSIGNED
);

CREATE TABLE IF NOT EXISTS Validators (
  name    VARCHAR(255),
  website VARCHAR(255),
  address VARCHAR(42) UNIQUE
);

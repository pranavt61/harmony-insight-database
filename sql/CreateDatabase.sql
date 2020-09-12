CREATE DATABASE IF NOT EXISTS HarmonyDashboard;

USE HarmonyDashboard;

CREATE TABLE IF NOT EXISTS BlockTransactionCount (
    shard_id TINYINT UNSIGNED,
    block_height INT UNSIGNED,
    tx_count SMALLINT UNSIGNED,
    PRIMARY KEY (shard_id, block_height)
);

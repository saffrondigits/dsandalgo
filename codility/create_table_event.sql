-- Given a table events with the following structure:
-- create table events (
-- sensor_id integer not null, event_type integer not null, value float not null, time timestamp not null
-- write an SQL query that returns a set of all sensors (sensor_id) with the number of different event types (event_type) registered by each of them, ordered by sensor_id (in ascending order).
-- For example, given the following data:
-- sensor_id
-- | event_type | value
-- time
-- 2
-- 1 3.45
-- 1-149.42
-- 2
-- | 4.12
-- 3
-- 2
-- 13.47
-- 2
-- 3
-- 24.54
-- your query should return the following rowset:
-- sensor id
-- types
-- 2014-02-13
-- 12:42:00
-- 2014-02-13
-- 12:42:00
-- 2014-02-13
-- 12:54:39
-- 2014-02-13
-- 12:42:00
-- 2014-02-13 13:32:36
-- 2
-- 3
-- ★
-- The names of the columns in the rowset don't matter, but their order does.

SELECT sensor_id, COUNT(DISTINCT event_type) AS types
FROM events
GROUP BY sensor_id
ORDER BY sensor_id ASC;

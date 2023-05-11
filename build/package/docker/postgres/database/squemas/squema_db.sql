CREATE SCHEMA pdi_go_kafka_db;

GRANT ALL PRIVILEGES ON DATABASE "go-kafka-db" TO "postgres";

GRANT USAGE ON SCHEMA pdi_go_kafka_db TO "postgres";
ALTER USER "postgres" SET search_path = 'pdi_go_kafka_db';

SET SCHEMA 'go-kafka-db';
ALTER DEFAULT PRIVILEGES
    IN SCHEMA pdi_go_kafka_db
GRANT SELECT, UPDATE, INSERT, DELETE ON TABLES
    TO "postgres";

ALTER DEFAULT PRIVILEGES
    IN SCHEMA pdi_go_kafka_db
GRANT USAGE ON SEQUENCES
    TO "postgres";
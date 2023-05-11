CREATE EXTENSION "uuid-ossp" SCHEMA public;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS users CASCADE;
create table users (
  id uuid DEFAULT(public.uuid_generate_v4()),
  name text);
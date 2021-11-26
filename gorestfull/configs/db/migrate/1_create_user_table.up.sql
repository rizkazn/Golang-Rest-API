CREATE TABLE public.user (
	id serial NOT NULL,
	name varchar(100) NOT NULL,
	username varchar(50) NOT NULL,
	email varchar(50) NOT NULL,
	password varchar(255) NOT NULL,
	role varchar(10) NOT NULL,
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT user_pk PRIMARY KEY (id)
);
CREATE TABLE public.category (
	id serial NOT NULL,
	category_name varchar(100) NOT NULL,
	category_image varchar(255) NOT NULL,
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT category_pk PRIMARY KEY (id)
);
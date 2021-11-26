CREATE TABLE public.product (
	id serial NOT NULL,
	product_name varchar(100) NOT NULL,
	price int4 NOT NULL,
	product_image varchar(255) NOT NULL,
	category_id int4 NOT NULL,
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT product_pk PRIMARY KEY (id)
    CONSTRAINT product_fk FOREIGN KEY (category_id) 
    REFERENCES public.category(id)
    ON DELETE CASCADE ON UPDATE CASCADE
);



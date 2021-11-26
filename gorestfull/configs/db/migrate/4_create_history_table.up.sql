CREATE TABLE public.history (
	id serial NOT NULL,
	no_invoice varchar(20) NOT NULL,
	cashier varchar(50) NOT NULL,
	date date NOT NULL DEFAULT CURRENT_DATE,
	orders int4 NOT NULL,
	amount int4 NOT NULL,
	CONSTRAINT history_pk PRIMARY KEY (id)
    CONSTRAINT history_fk FOREIGN KEY (orders)
    REFERENCES public.product(id)
    ON DELETE CASCADE ON UPDATE CASCADE
);



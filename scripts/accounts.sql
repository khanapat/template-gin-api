CREATE TABLE public.accounts (
	id varchar(36) NOT NULL,
	first_name varchar(255) NOT NULL,
	last_name varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	balance numeric NOT NULL DEFAULT 0,
	role_id int4 NOT NULL,
	created_date_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_date_time timestamp NULL,
	CONSTRAINT accounts_email_key UNIQUE (email),
	CONSTRAINT accounts_pkey PRIMARY KEY (id),
	CONSTRAINT role_id_fkey FOREIGN KEY (role_id) REFERENCES public."role"(id)
);

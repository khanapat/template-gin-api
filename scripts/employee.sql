CREATE TABLE public.employee (
	id serial4 NOT NULL,
	username varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	created_date_time timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_date_time timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT employee_email_key UNIQUE (email),
	CONSTRAINT employee_pkey PRIMARY KEY (id)
);

CREATE TABLE public."role" (
	id serial4 NOT NULL,
	title varchar(100) NOT NULL,
	description varchar(255) NOT NULL,
	CONSTRAINT role_pkey PRIMARY KEY (id)
);

INSERT INTO public."role"
(title, description)
values
('admin', 'admin role'),
('user', 'user role');

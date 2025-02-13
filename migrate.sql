CREATE TABLE USERS (
    id BIGINT PRIMARY KEY IDENTITY(1,1), 
    first_name NVARCHAR(255) NOT NULL DEFAULT '',
    last_name NVARCHAR(255) NOT NULL DEFAULT '',
    email VARCHAR(255) NOT NULL DEFAULT '',
    hash_pass VARCHAR(255) NOT NULL DEFAULT '',
    created_at DATETIMEOFFSET DEFAULT SYSDATETIMEOFFSET(),
    updated_at DATETIMEOFFSET DEFAULT SYSDATETIMEOFFSET(),
    deleted_at DATETIMEOFFSET NULL
);

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


CREATE TABLE ACTS  (
    id BIGINT PRIMARY KEY IDENTITY (1,1),
    act_number VARCHAR(255) NOT NULL DEFAULT '',
    sender_name NVARCHAR(255) NOT NULL DEFAULT '',
    sender_position NVARCHAR(255) NOT NULL DEFAULT '',
    sender_organization NVARCHAR(255) NOT NULL DEFAULT '',
    [data] DATE NOT NULL DEFAULT GETDATE(),
    receiver_name NVARCHAR(255) NOT NULL DEFAULT '',
    receiver_position NVARCHAR(255) NOT NULL DEFAULT '',
    receiver_organization NVARCHAR(255) NOT NULL DEFAULT '',
    created_at DATETIMEOFFSET DEFAULT SYSDATETIMEOFFSET(),
    updated_at DATETIMEOFFSET DEFAULT SYSDATETIMEOFFSET(),
    deleted_at DATETIMEOFFSET NULL
)

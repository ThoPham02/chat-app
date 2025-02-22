
CREATE TABLE conversations (
    conversation_id BIGINT,
    name VARCHAR(255),
    created_by BIGINT NOT NULL,
    created_at BIGINT,
    PRIMARY KEY (conversation_id)
);

CREATE TABLE conversation_members (
    id BIGINT,
    conversation_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    joined_at BIGINT,
    PRIMARY KEY (id)
);
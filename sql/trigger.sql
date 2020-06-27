CREATE OR REPLACE FUNCTION user_insert_trigger()
RETURNS TRIGGER AS $$
BEGIN
    IF ( NEW.updated_at_at is null ) THEN
        INSERT INTO new_users VALUES (NEW.*);
    ELSE
        INSERT INTO old_users VALUES (NEW.*);
    END IF;
    RETURN NULL;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER insert_user_trigger
    BEFORE INSERT ON users
    FOR EACH ROW EXECUTE PROCEDURE user_insert_trigger();
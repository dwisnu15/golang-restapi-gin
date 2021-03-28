CREATE TABLE IF NOT EXISTS items(
    id serial not null PRIMARY KEY,
    name varchar(225) not null,
    price bigint not null
);
/* PostgreSQL basically character varying is the alias name of varchar*/
/* */

CREATE VIEW view_item(item_id, item_name, item_price) AS
SELECT id,
       name,
       price
FROM items;

CREATE FUNCTION add_items(name_param character varying, price_param numeric) returns boolean
    language plpgsql
as $$
    BEGIN
        INSERT INTO items(name, price)
        VALUES(name_param, price_param);
        return true;
    end;
    $$;


/*i should create another entity if i have the time;
  to truly create view from joined table.
  i also dont really need to create this func, just as an example
  */

CREATE FUNCTION get_items()
    RETURNS TABLE(item_id numeric, item_name character varying, item_price numeric)
    language plpgsql
as
    $$
    BEGIN
        return query
        SELECT item_id, item_name, item_price FROM view_item LIMIT 20;
    end;
    $$;

/*i fear that this is a really bad way to implement patch.
  really, two queries for one update?
  */
CREATE FUNCTION update_items(item_id_param numeric, item_name_param character varying, item_price_param numeric)
returns boolean
    language plpgsql
as
    $$
    BEGIN
        if item_name_param IS NULL OR item_price_param IS NULL THEN
            return false;
        end if;
        return true;
--         if item_name_param IS NULL AND item_price_param IS NULL THEN
--             return false;
--         end if;
--         if item_name_param IS NOT NULL THEN
--             UPDATE items
--             SET name = item_name_param
--             WHERE id = item_id_param;
--         end if;
--         if item_price_param IS NOT NULL THEN
--             UPDATE items
--             SET price = item_price_param
--             WHERE id = item_id_param;
--         end if;
--         return true;
    end;
    $$;
/*
 i dont know yet how to do this better
 */
CREATE FUNCTION delete_item(item_id_param numeric)
    returns boolean
    language plpgsql
as
    $$
    DECLARE
        deletedrow numeric;
    BEGIN
        DELETE FROM items
        WHERE id = item_id_param returning id as deletedrow;
        if deletedrow < 1 THEN
            return false;
        end if;
        return true;
    end;
    $$
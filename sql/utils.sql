-- Active: 1696055093828@@127.0.0.1@5432@postgres
CREATE OR REPLACE FUNCTION get_ads(
    p_offset INTEGER,
    p_limit INTEGER,
    p_age INTEGER DEFAULT NULL,
    p_gender CHAR DEFAULT NULL,
    p_country CHAR(2) DEFAULT NULL,
    p_platform VARCHAR DEFAULT NULL
)
RETURNS TABLE(title VARCHAR, end_at TEXT) AS $$
DECLARE
    v_sql TEXT;
BEGIN
    -- Base FROM clause
    v_sql := 'SELECT sub_a.title, TO_CHAR(sub_a.end_at, ''YYYY-MM-DD HH24:MM:SS'') ';

    -- Filter by date
    v_sql := v_sql || 'FROM (
        SELECT a.id, a.title, a.end_at
        FROM ads a
        WHERE a.end_at > NOW() AND a.start_at < NOW() 
        ) AS sub_a ';

    -- Conditionally append joins
    IF p_age != 0 THEN
        v_sql := v_sql || 'LEFT JOIN ad_ages aa ON sub_a.id = aa.ad_id ';
    END IF;

    IF p_gender != '' THEN
        v_sql := v_sql || 'LEFT JOIN ad_genders ag ON sub_a.id = ag.ad_id ';
    END IF;

    IF p_country != '' THEN
        v_sql := v_sql || 'LEFT JOIN ad_countries ac ON sub_a.id = ac.ad_id ';
    END IF;

    IF p_platform != '' THEN
        v_sql := v_sql || 'LEFT JOIN ad_platforms ap ON sub_a.id = ap.ad_id ';
    END IF;

    -- Conditionally append WHERE conditions
    v_sql := v_sql || 'WHERE 1=1 ';
    IF p_age != 0 THEN
        v_sql := v_sql || 'AND (aa.age_start IS NULL OR (aa.age_start <= $1 AND aa.age_end >= $1)) ';
    END IF;
    IF p_gender != '' THEN
        v_sql := v_sql || 'AND (ag.gender IS NULL OR ag.gender = $2) ';
    END IF;
    IF p_country != '' THEN
        v_sql := v_sql || 'AND (ac.country_code IS NULL OR ac.country_code = $3) ';
    END IF;
    IF p_platform != '' THEN
        v_sql := v_sql || 'AND (ap.platform IS NULL OR ap.platform = $4) ';
    END IF;

    -- order by end_at
    v_sql := v_sql || 'ORDER BY sub_a.end_at DESC ';

    -- Append OFFSET and LIMIT
    v_sql := v_sql || 'OFFSET $5 ';

    IF p_limit != 0 THEN
        v_sql := v_sql || ' LIMIT $6 ';
    ELSE
        v_sql := v_sql || ' LIMIT 5 ';
    END IF;

    -- Execute the constructed query
    RETURN QUERY EXECUTE v_sql USING p_age, p_gender, p_country, p_platform, p_offset, p_limit;
END;
$$ LANGUAGE plpgsql;

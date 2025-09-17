-- Basic key-value pairs with colon syntax
SELECT json_object('id': 1);
SELECT json_object('name': 'John', 'age': 30);
SELECT json_object('active': true, 'score': null);

-- VALUE keyword syntax
SELECT json_object('id' VALUE 1);
SELECT json_object('name' VALUE 'Alice', 'age' VALUE 25);
SELECT json_object(KEY 'city' VALUE 'New York');

-- Mixed data types
SELECT json_object(
    'string': 'text',
    'number': 42,
    'float': 3.14,
    'boolean': false,
    'null': null
);

-- Nested objects
SELECT json_object(
    'user': json_object('id': 1, 'name': 'Bob'),
    'metadata': json_object('created': '2024-01-01', 'version': 2)
);

-- With expressions
SELECT json_object(
    'sum': 1 + 2,
    'concat': 'hello' || ' world',
    'date': CURRENT_DATE
);

-- Column references in FROM clause
SELECT json_object('id': t.id, 'value': t.value) FROM generate_series(1, 3) AS t(id);

-- Empty object
SELECT json_object();

-- Special characters in keys
SELECT json_object('key-with-dash': 1, 'key.with.dots': 2, 'key with spaces': 3);

-- Basic arrays
SELECT json_array(1, 2, 3);
SELECT json_array('a', 'b', 'c');
SELECT json_array(true, false, null);

-- Empty array
SELECT json_array();

-- Nested arrays
SELECT json_array(json_array(1, 2), json_array(3, 4));

-- Arrays with objects
SELECT json_array(
    json_object('id': 1, 'name': 'Item1'),
    json_object('id': 2, 'name': 'Item2')
);

-- With expressions
SELECT json_array(1 + 1, 2 * 3, 10 / 2);

-- Column values
SELECT json_array(t.n, t.n * 2, t.n * 3) FROM generate_series(1, 3) AS t(n);

-- Mixed types
SELECT json_array(1, 'string', true, null, 3.14, CURRENT_DATE);

-- Basic aggregation
SELECT json_arrayagg(n) FROM generate_series(1, 5) AS t(n);

-- With ORDER BY
SELECT json_arrayagg(n ORDER BY n DESC) FROM generate_series(1, 5) AS t(n);

-- With GROUP BY
SELECT
    category,
    json_arrayagg(value)
FROM (VALUES
    ('A', 1), ('A', 2),
    ('B', 3), ('B', 4)
) AS t(category, value)
GROUP BY category;

-- Aggregate objects
SELECT json_arrayagg(
    json_object('id': id, 'name': name)
    ORDER BY id
)
FROM (VALUES
    (1, 'Alice'),
    (2, 'Bob'),
    (3, 'Charlie')
) AS users(id, name);

-- With FILTER clause
SELECT json_arrayagg(n) FILTER (WHERE n > 2)
FROM generate_series(1, 5) AS t(n);

-- NULL handling
SELECT json_arrayagg(n) FROM (VALUES (1), (NULL), (3)) AS t(n);
SELECT json_arrayagg(n ABSENT ON NULL) FROM (VALUES (1), (NULL), (3)) AS t(n);

-- Basic object aggregation
SELECT json_objectagg(key: value)
FROM (VALUES
    ('a', 1),
    ('b', 2),
    ('c', 3)
) AS t(key, value);

-- VALUE syntax
SELECT json_objectagg(key VALUE value)
FROM (VALUES
    ('x', 10),
    ('y', 20)
) AS t(key, value);

-- With GROUP BY
SELECT
    department,
    json_objectagg(name: salary)
FROM (VALUES
    ('IT', 'John', 50000),
    ('IT', 'Jane', 60000),
    ('HR', 'Bob', 45000),
    ('HR', 'Alice', 48000)
) AS employees(department, name, salary)
GROUP BY department;

-- Nested aggregation
SELECT json_objectagg(
    category: json_arrayagg(item)
)
FROM (VALUES
    ('fruits', 'apple'),
    ('fruits', 'banana'),
    ('vegetables', 'carrot'),
    ('vegetables', 'lettuce')
) AS t(category, item)
GROUP BY category;

-- With NULL handling
SELECT json_objectagg(k: v ABSENT ON NULL)
FROM (VALUES ('a', 1), ('b', NULL), ('c', 3)) AS t(k, v);

-- Key uniqueness (last value wins)
SELECT json_objectagg(k: v)
FROM (VALUES ('a', 1), ('a', 2), ('a', 3)) AS t(k, v);

-- Nested object with arrays
SELECT json_object(
    'user': json_object(
        'id': 1,
        'name': 'John',
        'tags': json_array('admin', 'user', 'developer')
    ),
    'posts': json_array(
        json_object('id': 101, 'title': 'First Post'),
        json_object('id': 102, 'title': 'Second Post')
    )
);

-- Complex aggregation query
WITH sales_data AS (
    SELECT * FROM (VALUES
        ('2024-01', 'North', 'ProductA', 100),
        ('2024-01', 'North', 'ProductB', 150),
        ('2024-01', 'South', 'ProductA', 200),
        ('2024-02', 'North', 'ProductA', 120),
        ('2024-02', 'South', 'ProductB', 180)
    ) AS t(month, region, product, amount)
)
SELECT json_object(
    'summary': json_object(
        'total': SUM(amount),
        'count': COUNT(*)
    ),
    'by_month': (
        SELECT json_objectagg(
            month: json_object(
                'total': SUM(amount),
                'regions': json_arrayagg(DISTINCT region)
            )
        )
        FROM sales_data
        GROUP BY month
    ),
    'by_region': (
        SELECT json_objectagg(
            region: json_arrayagg(
                json_object('product': product, 'amount': amount)
            )
        )
        FROM sales_data
        GROUP BY region
    )
) FROM sales_data;

-- CTE with JSON construction
WITH json_data AS (
    SELECT json_object(
        'id': id,
        'data': json_array(id, id * 2, id * 3)
    ) AS obj
    FROM generate_series(1, 3) AS t(id)
)
SELECT json_arrayagg(obj) FROM json_data;

-- Multiple CTEs
WITH
users AS (
    SELECT * FROM (VALUES
        (1, 'Alice'),
        (2, 'Bob')
    ) AS t(id, name)
),
orders AS (
    SELECT * FROM (VALUES
        (101, 1, 50.00),
        (102, 1, 75.00),
        (103, 2, 100.00)
    ) AS t(order_id, user_id, amount)
)
SELECT json_arrayagg(
    json_object(
        'user': u.name,
        'orders': (
            SELECT json_arrayagg(
                json_object('id': order_id, 'amount': amount)
            )
            FROM orders o
            WHERE o.user_id = u.id
        )
    )
) FROM users u;

-- Unicode and special characters
SELECT json_object('emoji': '😀', '中文': '你好', 'special': E'\t\n\r');

-- Very long keys/values
SELECT json_object(
    'very_long_key_name_that_exceeds_normal_length': 'value',
    'key': repeat('x', 1000)
);

-- Numeric edge cases
SELECT json_object(
    'max_int': 2147483647,
    'min_int': -2147483648,
    'float': 1.23e-10,
    'infinity': 'Infinity'::float,
    'nan': 'NaN'::float
);

-- Date and time types
SELECT json_object(
    'date': DATE '2024-01-01',
    'time': TIME '12:34:56',
    'timestamp': TIMESTAMP '2024-01-01 12:34:56',
    'interval': INTERVAL '1 day 2 hours'
);

-- Subqueries in JSON construction
SELECT json_object(
    'count': (SELECT COUNT(*) FROM pg_tables),
    'random': (SELECT random()),
    'subquery_array': (
        SELECT json_arrayagg(tablename)
        FROM pg_tables
        WHERE schemaname = 'pg_catalog'
        LIMIT 5
    )
);

-- CASE expressions in JSON
SELECT json_object(
    'status': CASE
        WHEN 1 > 0 THEN 'positive'
        ELSE 'negative'
    END,
    'value': CASE 2
        WHEN 1 THEN 'one'
        WHEN 2 THEN 'two'
        ELSE 'other'
    END
);

-- RETURNING clause (if your grammar supports it)
SELECT json_object('id': 1) RETURNING JSON;
SELECT json_array(1, 2, 3) RETURNING JSONB;

-- NULL handling options
SELECT json_object('a': 1, 'b': NULL, 'c': 3) WITH UNIQUE KEYS;
SELECT json_arrayagg(n NULL ON NULL) FROM (VALUES (1), (NULL), (3)) AS t(n);

-- Format options
SELECT json_object('nested': json_object('key': 'value')) FORMAT JSON;
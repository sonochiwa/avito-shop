package repository

const (
	getInfo = `
SELECT row_to_json(row)
FROM (SELECT u.balance                                                                     AS balance,
             coalesce((SELECT json_agg(rows)
                       FROM (SELECT p.title AS product_title, i.quantity
                             FROM inventory AS i
                                      INNER JOIN products AS p ON p.id = i.product_id
                             WHERE i.user_id = $1) AS rows), '[]')                          AS inventory,
             (SELECT json_build_object(
                             'received', coalesce((SELECT json_agg(rows)
                                                   FROM (SELECT u.username AS from_user, amount
                                                         FROM ledger AS l
                                                                  INNER JOIN users AS u ON u.id = l.from_user
                                                         WHERE to_user = $1) rows), '[]'),
                             'sent', coalesce((SELECT json_agg(rows)
                                               FROM (SELECT u.username AS to_user, amount
                                                     FROM ledger AS l
                                                              INNER JOIN users AS u ON u.id = l.to_user
                                                     WHERE l.from_user = $1) rows), '[]'))) AS ledger
      FROM users AS u
      WHERE u.id = $1) row;
`
)

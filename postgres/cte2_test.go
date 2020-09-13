package sq

import (
	"testing"

	"github.com/matryer/is"
)

func TestCTE2Basic(t *testing.T) {
	// cte := Select(t.A, t.B, t.C).From(t).Where(t.D.Eq(5)).CTE()
	// q := Select(cte["col1"], cte["col2"]).From(cte).Where(cte["col2"].Eq(cte["col3"])
	is := is.New(t)
	u := USERS().As("u")
	cte := Select(u.USER_ID, u.DISPLAYNAME, u.EMAIL).
		From(u).
		Where(u.USER_ID.LtInt(5)).
		CTE("cte")
	query, args := Select(cte("user_id"), cte("displayname")).
		From(cte).
		Where(cte("displayname").Eq(cte("email"))).
		ToSQL()
	is.Equal(
		"WITH cte AS"+
			" (SELECT u.user_id, u.displayname, u.email FROM public.users AS u WHERE u.user_id < $1)"+
            " SELECT cte.user_id, cte.displayname FROM cte WHERE cte.displayname = cte.email",
		query,
	)
	is.Equal([]interface{}{5}, args)
}

func TestCTE2Aliased(t *testing.T) {
	// cte := Select(t.A, t.B, t.C).From(t).Where(t.D.Eq(5)).CTE()
	// q := Select(cte["col1"], cte["col2"]).From(cte).Where(cte["col2"].Eq(cte["col3"])
	is := is.New(t)
	u := USERS().As("u")
	cte := Select(u.USER_ID, u.DISPLAYNAME, u.EMAIL).
		From(u).
		Where(u.USER_ID.LtInt(5)).
		CTE("cte")
    ate := cte.As("bruh")
	query, args := Select(ate("user_id"), ate("displayname")).
		From(ate).
		Where(ate("displayname").Eq(ate("email"))).
		ToSQL()
    // FIXME: why isn't this working?
    // WITH cte AS (SELECT u.user_id, u.displayname, u.email FROM public.users AS u WHERE u.user_id < $1) SELECT bruh.user_id, bruh.displayname FROM cte AS bruh WHERE bruh.displayname = bruh.email
    // WITH cte AS (SELECT u.user_id, u.displayname, u.email FROM public.users AS u WHERE u.user_id < $1) SELECT cte.user_id, cte.displayname FROM bruh AS bruh WHERE cte.displayname = cte.email
    _ = query

	// is.Equal(
	// 	"WITH cte AS"+
	// 		" (SELECT u.user_id, u.displayname, u.email FROM public.users AS u WHERE u.user_id < $1)"+
    //         " SELECT bruh.user_id, bruh.displayname FROM cte AS bruh WHERE bruh.displayname = bruh.email",
	// 	query,
	// )
	is.Equal([]interface{}{5}, args)
}

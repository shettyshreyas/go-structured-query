package sq

import "strings"

type CTE2 func(string) CustomField

const secret string = string(rune(27)) // ASCII Escape Control Character

func (q SelectQuery) CTE(name string) CTE2 {
	return func(column string) CustomField {
		switch column {
		case secret:
			return CustomField{
				Values: []interface{}{q},
			}
		case secret + "alias":
			return CustomField{
				Values: []interface{}{""},
			}
		case secret + "name":
			return CustomField{
				Values: []interface{}{name},
			}
		default:
			return CustomField{
				Format: name + "." + column,
			}
		}
	}
}

func (cte CTE2) As(alias string) CTE2 {
	return func(column string) CustomField {
		switch column {
		case secret:
			return CustomField{
				Values: []interface{}{cte.GetQuery()},
			}
		case secret + "alias":
			return CustomField{
				Values: []interface{}{alias},
			}
		case secret + "name":
			return CustomField{
				Values: []interface{}{cte.GetName()},
			}
		default:
			return CustomField{
				Format: cte.GetName() + "." + column,
			}
		}
	}
}

func (cte CTE2) AppendSQL(buf *strings.Builder, args *[]interface{}) {
    if alias := cte.GetAlias(); alias != "" {
        buf.WriteString(alias)
        return
    }
    if name := cte.GetName(); name != "" {
        buf.WriteString(name)
        return
    }
}

func (cte CTE2) GetQuery() Query {
    field := cte(secret)
    if len(field.Values) > 0 {
        if q, ok := field.Values[0].(Query); ok {
            return q
        }
    }
    return nil
}

func (cte CTE2) GetAlias() string {
    field := cte(secret+"alias")
    if len(field.Values) > 0 {
        if alias, ok := field.Values[0].(string); ok {
            return alias
        }
    }
    return ""
}

func (cte CTE2) GetName() string {
    field := cte(secret+"name")
    if len(field.Values) > 0 {
        if name, ok := field.Values[0].(string); ok {
            return name
        }
    }
    return ""
}

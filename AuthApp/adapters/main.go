package adapters

type Adapters struct {
	SQL SQLAdapter
}

func (h *Adapters) ActivateSql(conn string) {
	var zeroDbLayer SQLAdapter
	sqlConn, _ := newSqlContext(conn)
	if h.SQL != zeroDbLayer {
		h.SQL.DB = sqlConn
	} else if h.SQL == zeroDbLayer {
		h.SQL = SQLAdapter{
			DB: sqlConn,
		}
	}
}

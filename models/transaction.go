package models

type transaction struct {
	OrdenID    string `json:"orden_id"`
	SessionID  string `json:"session_id"`
	Monto      int    `json:"monto"`
	UrlRetorno string `json:"url_retorno"`
}

var DatoTransaction = []transaction{
	{
		OrdenID:    "123AB",
		SessionID:  "345ab",
		Monto:      1500,
		UrlRetorno: "http://localhost:8080/commit",
	},
}

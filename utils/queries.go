package utils

const (
	// USER QUERIES
	GET_ALL_USER   = "SELECT user_name, password, email from m_user where user_name like ? order by ? ? limit ?, ?"
	GET_USER_BY_ID = "SELECT user_name, password, email FROM m_user WHERE user_name = ?"
	ADD_USER       = "INSERT INTO m_user(user_name, password, email) VALUES(?, ?, ?)"
	UPDATE_USER    = ""
	DELETE_USER    = ""

	// ARTICLE QUERIES
	GET_ALL_ARTICLE   = "SELECT * FROM m_article where user_creator = ? order by created_at"
	GET_ARTICLE_BY_ID = "SELECT * FROM m_article WHERE user_creator = ? and article_id = ? "
	ADD_ARTICLE       = "INSERT INTO m_article(article_id, article_name, article_content, user_creator, created_at) VALUES(?, ?, ?, ?, ?)"
	UPDATE_ARTICLE    = "UPDATE m_article SET article_name = ?, article_content = ? where user_creator = ? and article_id = ?"
	DELETE_ARTICLE    = "DELETE from m_article where user_creator = ? and article_id = ?"
)

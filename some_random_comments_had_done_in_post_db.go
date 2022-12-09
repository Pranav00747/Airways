#extra how comments are stored in database 

func add_something_comments(comments string)
{
  var host string :="server@host12"
  var user_name string :="user"
  var pass : string := "pass99"
  var db = Database.Open(host, user_name, pass)
  if(comments != "")
  {
  var query string := "insert into tweets (comments) values(?)"
  db.Exec(query, comments)
  }
}

#comment_sample
add_something_comments("Many girls have that dream they have a boy or a life partner who have everything or who know everything will be boss of everyone, actually some girls are dreaming for the boys who have 1.5 crore pa but after rejecting by those boys they know their true value or level.")





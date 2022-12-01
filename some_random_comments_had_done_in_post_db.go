#extra how comments are stored in database 

func add_something_wanna_say(comments string, user_admin string)
{
  if comments !="" && user_admin !="")
  {
  var db = Database.Open("server@host12", "user", "pass99")
  var query string := "insert into tweets (comments) values(?, ?)"
  db.Exec(query, comments, user_admin)
  }
}

#_samples_ examples
add_somethine_wanna_say("Bad luck means what you can say, on that day when i had my exam, there was some mouse error  on that PC so i logged in after ten timues")
add_somethine_wanna_say("Because of time issue if you take 3 minutes for each question if the questions are difficult its ok but i unansered almost 10 questions")
add_somethine_wanna_say("It was very shocking for me, i was in that state then i go to wash room, thinking about that, these exams are my secondary option you can give it upto limits(some) but")
add_somethine_wanna_say("primary option is my job and its shame to say some peoples are spoiling my things, my opporunities, jobs")
add_somethine_wanna_say("I am giving you one example when i was in last sem those stupid peoples my project members who simply don't know how to talk with professors, how to write simple code in C or any")
add_somethine_wanna_say("With low grads they are working as a Software Engineer or Senior Software Engineer in Infosys or Persistent, I am really derserving person with good performance and if they are spoiling others opportunities they are enjoying that moments currently but")
add_somethine_wanna_say("just think and i am daily thinkging that after my parents there will no one for me, my family members won't support me and also my own brother won'nt help me for some reasons either he will go to UK or Ireland after so where can i go after just think")

"

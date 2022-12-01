You can get the details like this with the help of database like SQL in Kothlin or Go

fun get_flight_det(source : string, to : string, date_str: string) : boolean
{
  val database : SqlLiteDatabase
  val db = database.readableDatabase
  val query = "select flight_name, available_seats, min_fair from @com_dom_air where from = "+source+" and dest = "+to+" date_string = "+date_str
  val cur = db.rawQuery(query)
  val fn = cur.getIndex(0)
  val avl = cur.getIndex(1)
  val fair = cur.getIndex(2)
  if(fn !="" && avl !="" && fair!="")
  {
    return true
  }
   return false
 }
 
 e.g
 get_flight_det("BLR", "JKF", "25-12-22")
 get_flight_det("BOM", "LHR", "25-12-22")
 get_flight_det("GMR", "AMS", "25-12-22")

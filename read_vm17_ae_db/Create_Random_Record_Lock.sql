-- Select a user at random
Select
   top 1 ts_id
From
   ts_Users
Where
   ts_id > Round((((Select Max(ts_id) from ts_Users) - 2 - 1) * Rand() + 2) , 0)
;
   
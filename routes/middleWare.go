package routes

//func IsLoggedin(f http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		user := models.User{}
//		sessId := r.Header.Get("sessionId")
//		_ = json.NewDecoder(r.Body).Decode(&user)
//		db := server.Conn().Model(&user).Where("session_id = ?", sessId).First(&user).Value
//		dbUser ,err := json.Marshal(db)
//		if err != nil{
//			fmt.Println(err)
//		}
//
//		err = json.Unmarshal(dbUser, user)
//
//		if (user.SessionId).String() == "00000000-0000-0000-0000-000000000000" {
//			fmt.Println("please login frist")
//		} else {
//			fmt.Println("already logged in")
//		}
//		f(w, r)
//	}
//}


package dao

import "time"

type Thread struct {
	Id        int64     `json:"id"`
	Uuid      string    `json:"uuid"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserId    int64     `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetThreadList() (threads []Thread, err error) {
	rows, err := DB.Query("SELECT id, uuid, title, content, user_id, created_at, updated_at FROM thread ORDER BY id DESC")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		thread := Thread{}
		err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Title, &thread.Content, &thread.UserId, &thread.CreatedAt, &thread.UpdatedAt)
		if err != nil {
			return
		}

		threads = append(threads, thread)
	}

	return
}

func GetThreadById(id int64) (thread Thread, err error) {
	err = DB.QueryRow("SELECT id, uuid, title, content, user_id, created_at, updated_at "+
		"FROM thread WHERE id = ? ORDER BY id DESC", id).Scan(&thread.Id, &thread.Uuid, &thread.Title, &thread.Content, &thread.UserId, &thread.CreatedAt, &thread.UpdatedAt)
	return
}

func Insert(thread *Thread) (id int64, err error) {
	result, err := DB.Exec("insert into thread(uuid, title, content, user_id) values (?, ?, ?, ?)",
		thread.Uuid, thread.Title, thread.Content, thread.UserId)
	if err != nil {
		return
	}

	id, _ = result.LastInsertId()
	return
}

func Update(thread *Thread) (err error) {
	_, err = DB.Exec("update thread set title = ?, content = ? where id = ?", thread.Title, thread.Content, thread.Id)
	return
}

func Delete(threadId int64) (err error) {
	_, err = DB.Exec("delete from thread where id = ?", threadId)
	return
}

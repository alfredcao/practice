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

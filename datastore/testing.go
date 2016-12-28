package datastore

const (
	InsertImageRegex       = "^INSERT INTO images \\(id, caption\\) VALUES \\(\\?, \\?\\)$"
	SelectImageByIDRegex   = "^SELECT id, caption FROM images WHERE id = \\?$"
	SelectLatestImageRegex = "^SELECT id, caption FROM images ORDER BY id DESC LIMIT 1$"
	SelectRandomImageRegex = "^SELECT id, caption FROM images WHERE id = \\(SELECT id FROM images ORDER BY RANDOM\\(\\) LIMIT 1\\)$"
)

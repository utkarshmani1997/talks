package main

const (
	INSERT_PULL_REQUEST_STATEMENT = `
		INSERT INTO pull_request(url, id, node_id, http_url, diff_url, patch_url,
		issue_url, commit_url, review_comments_url, review_comment_url, comments_url,
		statuses_url, number, state, title, body, creation_date, last_update_date,
		merged_date, closed_date, comments, review_comments, commits) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15,
		$16, $17, $18, $19, $20, $21, $22, $23, $24, $25)
	`
)

type PullRequest struct {
	url, id, node_id, commits, commit_url string
	...
}

func(p *PullRequest) Store(db *sql.DB) error {
    return db.Exec(p.url, p.id, p.node_id....)
}

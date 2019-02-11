package db

import "github.com/kongebra/go-practice/models"

// GetAllPosts returns all posts in the database
func (d *Database) GetAllPosts() ([]models.Post, error) {
	// Create an empty slice of posts
	posts := make([]models.Post, 0)

	// Run SQL-query
	rows, err := d.Get().Query("SELECT * FROM posts")
	// Check for connection error
	if err != nil {
		return nil, err
	}

	// Close connection
	defer rows.Close()

	// Loop through results
	for rows.Next() {
		// Create empty post object
		var post models.Post
		// Retrieve data from columns per row
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Created)
		// Check if there was an error
		if err != nil {
			d.Log("could not scan row: %v\n", err)
			return nil, err
		}
		// Push post to posts-slice
		posts = append(posts, post)
	}

	// Log event
	d.Log("event: GetAllPosts()")

	return posts, nil
}

// GetSinglePost returns a single post, based on id
func (d *Database) GetSinglePost(id int) ([]models.Post, error) {
	// Create an empty slice of posts
	posts := make([]models.Post, 0)

	// Run SQL-query
	rows, err := d.Get().Query("SELECT * FROM posts WHERE id = ?", id)
	// Check for connection error
	if err != nil {
		return nil, err
	}

	// Close connection
	defer rows.Close()

	// Loop through results
	for rows.Next() {
		// Create empty post object
		var post models.Post
		// Retrieve data from columns per row
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Created)
		// Check if there was an error
		if err != nil {
			d.Log("could not scan row: %v\n", err)
			return nil, err
		}
		// Push post to posts-slice
		posts = append(posts, post)
	}

	// Log event
	d.Log("event: GetSinglePost(%d)\n", id)

	return posts, nil
}
